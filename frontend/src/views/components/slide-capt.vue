<template>
  <gocaptcha-slide
    :data="handler.data"
    :events="{
      close: handler.closeEvent,
      refresh: handler.refreshEvent,
      confirm: handler.confirmEvent,
    }"
  />
</template>

<script setup>
import { useHandler } from "@/api/captcha";
const handler = useHandler({
  getApi: "http://127.0.0.1:6606/api/v1/captcha",
  checkApi: "http://127.0.0.1:6606/api/v1/captcha"
});
</script>

<style></style>
