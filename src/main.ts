import {createApp} from 'vue'
import App from "./App.vue"
import {
    Button,
    message
} from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css';
const app=createApp(App);

app
.use(Button)
.mount('#app')
//@ts-ignore
window.app=app;

app.config.globalProperties.$message=message
