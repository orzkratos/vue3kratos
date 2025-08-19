import {GrpcWebFetchTransport} from '@protobuf-ts/grpcweb-transport'
import {RpcdemoClient} from "./rpc/rpcdemo/rpcdemo.client";
import {
    CreateRpcDemoRequest,
    DeleteRpcDemoRequest,
    SelectRpcDemoRequest,
    UpdateRpcDemoRequest
} from "./rpc/rpcdemo/rpcdemo";

console.log("---")
console.log('开始调用 RpcDemo CRUD 接口...');
console.log("---")

// 创建共享的传输实例
const demoTransport = new GrpcWebFetchTransport({
    baseUrl: "http://127.0.0.1:28000",
    meta: {
        Authorization: 'TOKEN-888',
    },
});

export const rpcdemoClient = new RpcdemoClient(demoTransport)

// 演示 CreateRpcDemo 接口
async function demoCreateRpcDemo() {
    console.log("=== 测试 CreateRpcDemo 接口 ===")

    const request = CreateRpcDemoRequest.create({
        code: "DEMO001",
        name: "Vue3 Kratos Demo",
        type: "frontend"
    })

    try {
        const response = await rpcdemoClient.createRpcDemo(request, {})
        console.log('创建成功，返回 code:', response.data.code)
        return response.data.code
    } catch (cause: any) {
        console.error('创建失败:', cause)
        throw cause
    }
}

// 演示 UpdateRpcDemo 接口
async function demoUpdateRpcDemo(code: string) {
    console.log("=== 测试 UpdateRpcDemo 接口 ===")

    const request = UpdateRpcDemoRequest.create({
        code: code,
        name: "Updated Vue3 Kratos Demo"
    })

    try {
        const response = await rpcdemoClient.updateRpcDemo(request, {})
        console.log('更新成功，返回:', response.data)
        return response.data
    } catch (cause: any) {
        console.error('更新失败:', cause)
        throw cause
    }
}

// 演示 SelectRpcDemo 接口
async function demoSelectRpcDemo() {
    console.log("=== 测试 SelectRpcDemo 接口 ===")

    const request = SelectRpcDemoRequest.create({
        type: "frontend"
    })

    try {
        const response = await rpcdemoClient.selectRpcDemo(request, {})
        console.log('查询成功，找到 demos 数量:', response.data.demos.length)
        response.data.demos.forEach((item, index) => {
            console.log(`Demo ${index + 1}:`, item)
        })
        return response.data.demos
    } catch (cause: any) {
        console.error('查询失败:', cause)
        throw cause
    }
}

// 演示 DeleteRpcDemo 接口
async function demoDeleteRpcDemo(code: string) {
    console.log("=== 测试 DeleteRpcDemo 接口 ===")

    const request = DeleteRpcDemoRequest.create({
        code: code
    })

    try {
        const response = await rpcdemoClient.deleteRpcDemo(request, {})
        console.log('删除成功，返回 code:', response.data.code)
        return response.data.code
    } catch (cause: any) {
        console.error('删除失败:', cause)
        throw cause
    }
}

// 执行完整的 CRUD 演示流程
async function runCrudDemo() {
    try {
        // 1. 创建
        const createdCode = await demoCreateRpcDemo()

        // 2. 更新
        await demoUpdateRpcDemo(createdCode)

        // 3. 查询
        await demoSelectRpcDemo()

        // 4. 删除
        await demoDeleteRpcDemo(createdCode)

        console.log("---")
        console.log('RpcDemo CRUD 演示完成！')
        console.log("---")

    } catch (err) {
        console.error('CRUD 演示过程中发生错误:', err)
    }
}

// 启动演示
runCrudDemo()

console.log("---")
console.log('RpcDemo CRUD 演示已启动...')
console.log("---")