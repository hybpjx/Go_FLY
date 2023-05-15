import { createApp } from 'vue'
import App from './App.vue'

import { initApp } from './config/init'

import "normalize.css/normalize.css"
import "./assets/fonts/iconfont.css"
import "./assets/styles/global.scss"

(async () => {

    // =================================================================
    // =初始化系统基础配置信息（保证所有的模块的基础数据加载完成后，才创建UI）
    // 1. 全局变量（app:应用 挂载全局的方法）,语言包（lpk:获取文本内容），Ajax，Tools 的定义
    // 2. 异步加载基础模块的配置信息
    // 3. 异步加载业务模块的配置信息并且完成一些基础的初始化

    initApp()

    // =================================================================
    // = 初始化UI
    const uiAPP = createApp(App)

    // =================================================================
    // = 注册全局组件 

    // =================================================================
    // = 向根组件绑定全局对象
    uiAPP.config.globalProperties.app = window.app
    uiAPP.config.globalProperties.app = window.Tools;
    // =================================================================
    // = 初始化状态管理与路由，并且渲染根组件

    uiAPP.mount("#app")
})()
