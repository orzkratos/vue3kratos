import {GrpcWebFetchTransport} from '@protobuf-ts/grpcweb-transport'
import {RpcpingClient} from "./rpc/rpcping/rpcping.client";
import {StringValue} from "./rpc/google/protobuf/wrappers";

console.log("---")
console.log('开始调用 RpcPing 接口...');
console.log("---")

// 创建传输实例
const demoTransport = new GrpcWebFetchTransport({
    baseUrl: "http://127.0.0.1:28000",
    meta: {
        Authorization: 'TOKEN-888',
    },
});

export const rpcpingClient = new RpcpingClient(demoTransport)

// 演示 Ping 接口
async function demoPing() {
    console.log("=== 测试 Ping 接口 ===")

    // 创建 StringValue 请求
    const request = StringValue.create({
        value: "Hello from Vue3 Kratos!"
    })

    try {
        const response = await rpcpingClient.ping(request, {})
        console.log('Ping 成功，服务器响应:', response.data.value)
        return response.data.value
    } catch (cause: any) {
        console.error('Ping 失败:', cause)
        throw cause
    }
}

// 演示多次 Ping 调用
async function demoMultiplePings() {
    const messages = [
        "Ping 1: Vue3 integration test",
        "Ping 2: HTTP via gRPC works!",
        "Ping 3: Kratos backend rocks",
        "Ping 4: TypeScript client success",
        "Ping 5: Final test complete"
    ]

    console.log("=== 测试多次 Ping 调用 ===")

    for (let i = 0; i < messages.length; i++) {
        try {
            const request = StringValue.create({
                value: messages[i]
            })

            const response = await rpcpingClient.ping(request, {})
            console.log(`第 ${i + 1} 次 Ping:`)
            console.log(`  发送: ${messages[i]}`)
            console.log(`  响应: ${response.data.value}`)

            // 添加小延迟以便观察
            await new Promise(resolve => setTimeout(resolve, 500))

        } catch (cause: any) {
            console.error(`第 ${i + 1} 次 Ping 失败:`, cause)
        }
    }
}

// 执行完整的 Ping 演示
async function runPingDemo() {
    try {
        // 1. 单次 Ping
        await demoPing()

        console.log("\n")

        // 2. 多次 Ping
        await demoMultiplePings()

        console.log("---")
        console.log('RpcPing 演示完成！')
        console.log("---")

    } catch (err) {
        console.error('Ping 演示过程中发生错误:', err)
    }
}

// 启动演示
runPingDemo()

console.log("---")
console.log('RpcPing 演示已启动...')
console.log("---")