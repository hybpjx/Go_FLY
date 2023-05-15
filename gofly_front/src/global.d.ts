import { IApp } from "./config/app";
import { iTools } from "./utils/utools";


declare global {
    declare namespace GlobalType {
        type iKey = string | number;
        type iRecord = Record<iKey, any>;
    };
    const app = IApp;
    interface Window {
        app: IApp; // 全局APP方法 挂载一些全局数据与操作方法。
        Tools: ITools; // 全局公用方法; 全局工具库对象
    };
}


declare module "vue" {
    interface ComponentCustomProperties {
        app: IApp;
    }
}
export { }
