import { createApp } from "vue";
import App from "./App.vue";
import setupPlugins from "@/plugins";
// 本地SVG图标
import "virtual:svg-icons-register";

// 样式
import "element-plus/theme-chalk/dark/css-vars.css";
import "@/styles/index.scss";
import "uno.css";
import "animate.css";
import "go-captcha-vue/dist/style.css";
import GoCaptcha from "go-captcha-vue";
const app = createApp(App);
app.use(GoCaptcha);
app.use(setupPlugins);
app.mount("#app");
