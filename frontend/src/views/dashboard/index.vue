<template>
  <div class="app-container">
    <h1>Vue3-Element-Admin-Thin</h1>
    <h1>浏览器指纹：{{ visitorId }}</h1>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { GetFingerprint, GetHash } from "@/utils/fingerprint";

// 定义组件选项
defineOptions({
  name: "Dashboard",
  inheritAttrs: false,
});

// 定义响应式变量
const visitorId = ref<string>("");

// 在组件挂载时加载 FingerprintJS
onMounted(async () => {
  visitorId.value = await GetFingerprint();
});
</script>

<style lang="scss" scoped></style>
