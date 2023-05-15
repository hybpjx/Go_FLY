import iSysCfg, { ISysCfg, ISysCfgBmodItem } from "./syscfg"

const app = {
    // ! 获取系统配置信息
    getConfig<T>(key: keyof ISysCfg): T {
        return iSysCfg[key] as T;
    },
    // ! 判断是否启用了指定的业务模块
    checkBmodIsEnable(stModuleName: string): Boolean {
        const bmodNames: ISysCfgBmodItem[] = app.getConfig<ISysCfgBmodItem[]>("bmodNames");
        if (bmodNames.find(item => item.name == stModuleName && item.enable == true)) {
            return true;
        }
        return false
    }
}

export type IApp = typeof app

export default app

