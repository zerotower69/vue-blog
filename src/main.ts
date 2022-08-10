import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import { message } from "ant-design-vue";
import Antd from "ant-design-vue";
import "ant-design-vue/dist/antd.css";

const app = createApp(App);
app.config.globalProperties.$message = message;
app.use(Antd);
app.use(store).use(router).mount("#app");
