export interface ISysCfgBmodItem {
    name: string; // 模块名称
    enable: boolean; // 是否启用
}

export interface ISysCfg {
    baseUrl: string; // 主机地址与监听端口
    bmodNames: ISysCfgBmodItem[];// 业务模块列表
}

const iSysCfg: ISysCfg = {
    baseUrl: "https://192.168.2.51:8080",
    bmodNames: [{
        name: "blog",
        enable: true
    }],
}

export default iSysCfg;