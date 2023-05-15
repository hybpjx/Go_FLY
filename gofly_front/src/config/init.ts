import app from "./app";
import Tools from "@/utils/utools";

//  声明全局变量相关的类型
type IGlobalVarsKey = 'app' | 'lpk' | 'Tools' | 'Ajax';
type iGlobalVars = {
    [key in IGlobalVarsKey]?: any;
}


const iGobalVars: iGlobalVars = {
    app, // 全局应用对象，包含一些全局数据与操作的方法
    Tools //
}

Object.keys(iGobalVars).forEach(stkey => {
    (window as any)[stkey as IGlobalVarsKey] = iGobalVars[stkey as IGlobalVarsKey];
})

export const initApp = async () => {

}