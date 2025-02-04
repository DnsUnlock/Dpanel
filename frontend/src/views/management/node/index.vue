<script setup lang="ts">
import SlideCapt from "@/views/components/slide-capt.vue";
import { NMessageProvider } from "naive-ui";
const dialogTableVisible = ref(false);
import { getCurrentInstance } from "vue";

const { proxy } = getCurrentInstance()!;

function sendMessage() {
  proxy.$nexus
    .post("/test", { name: "button" })
    .then((resp: any) => {
      // 此处处理服务器响应
      console.log("服务器响应:", resp);
    })
    .catch((err: Error) => {
      // 此处处理错误
      console.error("请求错误:", err);
    });
}
</script>
<template>
  <div>
    <el-dialog
      v-model="dialogTableVisible"
      title="人机验证"
      style="width: 326px; padding: 16px 12px; text-align: center"
    >
      <n-message-provider>
        <div class="item"><slide-capt /></div>
      </n-message-provider>
    </el-dialog>
    <div class="app-container">
      <h1>Vue3-Element-Admin-Thin</h1>
      <el-button plain @click="dialogTableVisible = true">人机验证</el-button>
    </div>
    <div>
      <el-button plain @click="sendMessage">发送请求</el-button>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.item {
  margin-top: -10px;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
