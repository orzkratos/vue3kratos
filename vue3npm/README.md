# 前端相关代码

这里有和前端配合相关的代码。

注意这不是一个完整的项目，而是只有几个有用的文件，里面有基本的逻辑。

但是因为在IDE里面总是提示找不到包名，因而把它做成子项目（带包管理）。

## 使用举例
需要在当前目录里执行：
```
npm run rpcrewrite -- /xxx/src/rpc/rpc_admin_login/admin_login.client.ts
```
当然您也可以把它拷贝到您的 vue3 项目里，把命令也拷贝粘贴修改，依然可以使用。

## 其它信息
由于需要用到 js 因此需要在 package.json 里面配置 `"type": "module",` 这项内容。
