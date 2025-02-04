import { v4 as uuidv4 } from "uuid";
import type { App, Plugin } from "vue";

// 记录请求信息
interface RequestMessage {
  id: string;
  method: string;
  path: string;
  body?: unknown;
  header?: Record<string, string>;
}

// 记录响应信息
interface ResponseMessage {
  id: string;
  status: number;
  body?: unknown;
  header?: Record<string, string>;
}

// 记录请求配置
interface RequestConfig {
  headers?: Record<string, string>;
  timeout?: number;
  [key: string]: any;
}

class NexusClient {
  // 记录连接地址
  private path: string;
  // WebSocket 连接对象
  private conn: WebSocket | null;
  // 存储等待响应的请求回调
  private pending: Record<string, (resp: ResponseMessage) => void>;
  // 是否已手动关闭
  private closed: boolean;
  // 默认缓存请求头
  private headers: Record<string, string>;
  // 当前的重连尝试次数
  private reconnectAttempts: number;
  // 最大重连次数
  private readonly maxReconnectAttempts: number;
  // 每次重连延迟，单位毫秒
  private readonly reconnectDelay: number;
  // 记录上一次发送消息的时间
  private lastSendTime: number;
  // 心跳定时器
  private heartbeatTimer: number | null;

  constructor(scheme: string, host: string, path: string) {
    // 检查连接参数是否完整，不完整直接抛错
    if (!scheme || !host || !path) {
      throw new Error("连接参数不足，无法初始化");
    }

    // 组合 URL
    this.path = `${scheme}://${host}${path}`;
    // 初始化 WebSocket 连接对象
    this.conn = null;
    // 存储请求的回调对象
    this.pending = {};
    // 默认情况下未关闭
    this.closed = false;
    // 初始化请求头
    this.headers = {};
    // 初始化重连配置
    this.reconnectAttempts = 0;
    this.maxReconnectAttempts = 5;
    this.reconnectDelay = 2000;
    // 初始化上一次发送消息的时间
    this.lastSendTime = Date.now();
    // 心跳定时器初始化
    this.heartbeatTimer = null;
    this.initializeDefaultHeaders(this.path);
    // 创建实例时直接连接
    this.connect();
  }

  // 初始化默认请求头
  private async initializeDefaultHeaders(path: string): Promise<void> {
    // 使用fetch去获取服务器返回的响应

    if (path.startsWith("wss://")) {
      path = path.replace("wss://", "https://");
    } else if (path.startsWith("ws://")) {
      path = path.replace("ws://", "http://");
    }

    const response = await fetch(path, {
      method: "GET",
    });

    if (!response.ok) {
      throw new Error(`请求失败，状态码: ${response.status}`);
    }

    // 将Headers对象转换为普通键值对
    const headersObj: Record<string, string> = {};
    response.headers.forEach((value, key) => {
      headersObj[key] = value;
    });

    // 将解析后的响应头赋值到this.headers
    this.headers = headersObj;
  }

  // 解析请求头字符串
  private parseHeaders(headers: string): Record<string, string> {
    return headers
      .split("\r\n")
      .filter((header) => header.includes(":"))
      .reduce(
        (acc, line) => {
          const [key, value] = line.split(": ");
          acc[key.trim()] = value.trim();
          return acc;
        },
        {} as Record<string, string>
      );
  }

  // 开始连接
  connect(): void {
    // 如果已手动关闭，不再继续连接
    if (this.closed) {
      return;
    }
    // 重置重连次数
    this.reconnectAttempts = 0;
    // 创建 WebSocket 连接
    this.conn = new WebSocket(this.path);

    // 接收到服务器消息时调用
    this.conn.onmessage = (event: MessageEvent) => {
      this.readMessages(event.data);
    };

    // 监听错误事件
    this.conn.onerror = () => {
      // 可以在这里打印或记录日志
    };

    // 监听连接关闭事件
    this.conn.onclose = () => {
      // 如果不是手动关闭，启动重连
      if (!this.closed) {
        this.handleReconnect();
      }
    };

    // 启动心跳检测
    this.startHeartbeat();
  }

  // 处理重连逻辑
  private handleReconnect(): void {
    // 如果超出最大重连次数，则直接返回
    if (this.reconnectAttempts >= this.maxReconnectAttempts) {
      return;
    }
    this.reconnectAttempts++;
    setTimeout(() => {
      // 如果依然没有手动关闭，继续尝试重连
      if (!this.closed) {
        this.connect();
      }
    }, this.reconnectDelay);
  }

  // 心跳函数，每秒检查一次，如果10s内未发送任何消息则发送ping
  private startHeartbeat(): void {
    // 如果心跳已经在运行则不重复开启
    if (this.heartbeatTimer) {
      return;
    }
    // 每秒检查一次
    this.heartbeatTimer = window.setInterval(() => {
      // 如果连接不可用，直接返回
      if (!this.conn || this.conn.readyState !== WebSocket.OPEN) {
        return;
      }
      // 判断距离上一次发送消息是否超过10s
      const now = Date.now();
      if (now - this.lastSendTime >= 10000) {
        try {
          // 发送ping消息
          this.conn.send(JSON.stringify({ Type: "ping" }));
          // 更新最后发送消息时间
          this.lastSendTime = now;
        } catch (_) {
          // 可以在这里打印或记录日志
        }
      }
    }, 1000);
  }

  // GET 方法
  async get(
    path: string,
    config: RequestConfig = {}
  ): Promise<ResponseMessage> {
    return this.request({
      method: "GET",
      path,
      data: null,
      config,
    });
  }

  // POST 方法
  async post(
    path: string,
    data: unknown,
    config: RequestConfig = {}
  ): Promise<ResponseMessage> {
    return this.request({
      method: "POST",
      path,
      data,
      config,
    });
  }

  // PUT 方法
  async put(
    path: string,
    data: unknown,
    config: RequestConfig = {}
  ): Promise<ResponseMessage> {
    return this.request({
      method: "PUT",
      path,
      data,
      config,
    });
  }

  // DELETE 方法
  async delete(
    path: string,
    config: RequestConfig = {}
  ): Promise<ResponseMessage> {
    return this.request({
      method: "DELETE",
      path,
      data: null,
      config,
    });
  }

  // 统一处理请求逻辑
  private async request(options: {
    method: string;
    path: string;
    data: unknown;
    config: RequestConfig;
  }): Promise<ResponseMessage> {
    const { method, path, data, config } = options;
    const reqMessage: RequestMessage = {
      id: uuidv4(),
      method: method,
      path: path,
      body: data || {},
      header: { ...this.headers, ...config.headers },
    };
    return this.sendRequest(reqMessage, config.timeout);
  }

  // 发送请求并等待响应
  private sendRequest(
    req: RequestMessage,
    timeout = 10000
  ): Promise<ResponseMessage> {
    return new Promise((resolve, reject) => {
      // 如果连接已关闭或尚未建立，直接返回错误
      if (
        !this.conn ||
        this.closed ||
        this.conn.readyState !== WebSocket.OPEN
      ) {
        reject(new Error("连接不可用"));
        return;
      }

      // 更新最后发送消息时间
      this.lastSendTime = Date.now();

      // 存储该请求对应的回调
      this.pending[req.id] = (resp: ResponseMessage) => {
        resolve(resp);
      };

      // 发送请求
      try {
        this.conn.send(JSON.stringify(req));
      } catch (err) {
        delete this.pending[req.id];
        reject(err);
        return;
      }

      // 启动超时计时
      setTimeout(() => {
        if (this.pending[req.id]) {
          delete this.pending[req.id];
          reject(new Error("请求超时"));
        }
      }, timeout);
    });
  }

  // 解析并调用对应请求的回调
  private readMessages(message: string): void {
    let resp: ResponseMessage;
    try {
      resp = JSON.parse(message);
    } catch (_) {
      // 如果消息无法解析为JSON，这里可以记录日志后return
      return;
    }
    if (resp && resp.id && this.pending[resp.id]) {
      const callback = this.pending[resp.id];
      callback(resp);
      delete this.pending[resp.id];
    }
  }

  // 手动关闭连接
  Close(): void {
    this.closed = true;
    // 清除心跳定时器
    if (this.heartbeatTimer) {
      clearInterval(this.heartbeatTimer);
      this.heartbeatTimer = null;
    }
    if (this.conn) {
      this.conn.close();
    }
  }

  // 返回默认请求头
  private defaultHeader(): Record<string, string> {
    return this.headers;
  }
}

// 封装为 Vue 插件
const NexusClientPlugin: Plugin = {
  install(
    app: App,
    options?: { scheme?: string; host?: string; path?: string }
  ): void {
    const { scheme, host, path } = options || {};
    // 如果参数不完整，则直接抛错
    if (!scheme || !host || !path) {
      throw new Error("初始化参数不足，无法创建 NexusClient");
    }
    app.config.globalProperties.$nexus = new NexusClient(scheme, host, path);
  },
};

export { NexusClient, NexusClientPlugin };
export default NexusClientPlugin;
