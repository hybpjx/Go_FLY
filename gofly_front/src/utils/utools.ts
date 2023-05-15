import cookies from 'js-cookie'

const iTools = {
    // 路由操作命名空间
    Router: {

    },
    // 状态管理命名空间
    Store: {

    },
    // 本地存储命名空间
    LoacalStorage: {
        setItem(key: string, value: any) {
            localStorage.setItem(key, JSON.stringify(value))
        },
        getItem(key: string) {
            const stValue = localStorage.getItem(key)
            try {
                return JSON.parse(stValue as string)
            } catch (e) {
                return stValue
            }
        },
        removeItem(key: string) {
            localStorage.removeItem(key)
        }
    },
    // cookie操作命名空间
    Cookie: {
        setItem(key: string, value: any) {
            cookies.set(key, value, { expires: 30 })
        },
        getItem(key: string, defaultValue: any) {
            const stValue = cookies.get(key) || defaultValue
            try {
                return JSON.parse(stValue as string)
            } catch (e) {
                return stValue
            }
        },
        removeItem(key: string) {
            cookies.remove(key)
        }
    },
    // 日期操作命名空间
    Time: {

    },
    // Dom 对象操作命名空间
    Dom: {

    }
}

export type iTools = typeof iTools

export default iTools