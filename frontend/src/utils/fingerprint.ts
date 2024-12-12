import FingerprintJS from "@fingerprintjs/fingerprintjs";
import { FINGER_KEY } from "@/enums/CacheEnum";
import md5 from "md5"; // 使用 md5 库

export function SetFingerprintToken() {
  // 加载 FingerprintJS 库
  FingerprintJS.load()
    .then(fp => fp.get())
    .then(result => {
      const visitorId = result.visitorId;
      const existingToken = localStorage.getItem(FINGER_KEY);

      // 如果 FINGER_KEY 不存在或值不同，则更新 Token
      if (!existingToken || existingToken.split('=')[1] !== visitorId) {
        // 更新值
        localStorage.setItem(FINGER_KEY, visitorId);
      } else {
        // fingerprint Token 已存在且相同
      }
    })
    .catch(error => {
      console.error("FingerprintJS 加载失败:", error);
    });
}

export async function GetFingerprint(): Promise<string> {
  try {
    const fp = await FingerprintJS.load();
    const result = await fp.get();
    return result.visitorId;
  } catch (error) {
    console.error("FingerprintJS 加载失败:", error);
    return "";
  }
}

export async function GetHash(): Promise<string> {
  try {
    const fp = await FingerprintJS.load();
    const result = await fp.get();
    const visitorId = result.visitorId;
    const timestamp = Math.floor(Date.now() / 1000); // 获取当前时间戳（秒）

    // 拼接指纹和时间戳并生成 MD5 哈希值
    const hash = md5(`${visitorId}${timestamp}`);
    let hashStr = "";

    if (Array.isArray(hash)) {
      hashStr = hash.join('');
    } else {
      hashStr = hash;
    }

    return hashStr;
  } catch (error) {
    console.error("FingerprintJS 加载失败:", error);
    return "";
  }
}

