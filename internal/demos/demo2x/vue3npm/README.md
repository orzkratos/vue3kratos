# 前端相关代码

这里有和前端配合相关的代码。

注意这不是一个完整的项目，而是只有几个有用的文件，里面有基本的逻辑。

但是因为在IDE里面总是提示找不到包名，因而把它做成子项目（带包管理）。

## 使用举例
通过外面的 `copyrpc` 的 `main` 执行即可得到你想到的代码。

## 备选方案
通过前面的 `copyrpc` 已经可以解决问题，其他备选方案是，通过js命令改造生成的客户端。

例如，在当前目录里执行改造逻辑：
```bash
npm run grpcrewrite -- src/rpc/v1/greeter/greeter.client.ts
```
得到的代码和 `copyrpc` 是完全等效的。

当然具体使用时推荐使用绝对路径:
```bash
npm run grpcrewrite -- /xxx/yyy/zzz/src/rpc/v1/greeter/greeter.client.ts
```
得到的代码和 `copyrpc` 是完全等效的。

## 其它信息
由于需要用到 js 因此需要在 package.json 里面配置 `"type": "module",` 这项内容。

由于代码中会用到 `BigInt` 因此需要设置 `"target": "ES2023",` 使用最新的稳定 `tsconfig.json` 规则。
