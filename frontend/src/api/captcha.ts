import Lodash from "lodash";
import request from "@/utils/request";
import { useMessage } from "naive-ui";
import Qs from "qs";
import { onMounted, reactive, watch } from "vue";

interface Config {
  getApi: string; // 验证码获取接口地址
  checkApi: string; // 验证码验证接口地址
}

interface Point {
  x: number; // x 轴坐标
  y: number; // y 轴坐标
}

interface CaptchaData {
  image: string; // 验证码背景图片的 base64 数据
  thumb: string; // 验证码滑块图片的 base64 数据
  captKey: string; // 验证码的唯一标识符
  thumbX: number; // 滑块的 x 轴坐标
  thumbY: number; // 滑块的 y 轴坐标
  thumbWidth: number; // 滑块的宽度
  thumbHeight: number; // 滑块的高度
}

interface State {
  popoverVisible: boolean; // 验证码弹出框是否可见
  type: string; // 当前验证码的状态类型（"default", "success", "error"）
}

export const useHandler = (config: Config) => {
  const message = useMessage();
  // 定义组件状态，包含验证码弹出框是否显示以及验证码类型
  const state = reactive<State>({ popoverVisible: false, type: "default" });
  // 定义验证码相关数据，包括图片信息、验证码唯一标识符等
  const cData = reactive<CaptchaData>({
    image: "",
    thumb: "",
    captKey: "",
    thumbX: 0,
    thumbY: 0,
    thumbWidth: 0,
    thumbHeight: 0,
  });

  // 点击事件，显示验证码弹出框
  const clickEvent = (): void => {
    state.popoverVisible = true;
  };

  // 验证码弹出框可见性变化事件
  const visibleChangeEvent = (visible: boolean): void => {
    state.popoverVisible = visible;
  };

  // 关闭验证码弹出框事件
  const closeEvent = (): void => {
    state.popoverVisible = false;
  };

  // 请求验证码数据
  const requestCaptchaData = (): void => {
    request<any, any>({
      url: config.getApi,
      method: "get",
    })
      .then((data) => {
        // 设置验证码数据
        cData.image = data["image_base64"] || "";
        cData.thumb = data["tile_base64"] || "";
        cData.captKey = data["captcha_key"] || "";
        cData.thumbX = data["tile_x"] || 0;
        cData.thumbY = data["tile_y"] || 0;
        cData.thumbWidth = data["tile_width"] || 0;
        cData.thumbHeight = data["tile_height"] || 0;
      })
      .catch((e) => {
        console.warn(e);
      });
  };

  // 刷新验证码事件
  const refreshEvent = (): void => {
    requestCaptchaData();
  };

  // 确认验证码事件
  const confirmEvent = (point: Point, clear: () => void): void => {
    request<any, any>({
      url: config.checkApi,
      method: "post",
      data: {
        point: [point.x, point.y].join(","),
        key: cData.captKey || "",
      },
    })
      .then((data) => {
        console.log(data);
        // 验证完成后，清除滑块位置并重新请求验证码
        setTimeout(() => {
          clear();
          requestCaptchaData();
        }, 1000);
      })
      .catch((e) => {
        console.warn(e);
      });
  };

  // 监听验证码弹出框可见性，当弹出框显示时请求验证码数据
  watch(
    () => state.popoverVisible,
    () => {
      if (state.popoverVisible) {
        requestCaptchaData();
      }
    }
  );

  // 组件挂载时请求验证码数据
  onMounted(() => {
    requestCaptchaData();
  });

  return {
    state,
    data: cData,
    visibleChangeEvent,
    clickEvent,
    closeEvent,
    refreshEvent,
    confirmEvent,
  };
};
