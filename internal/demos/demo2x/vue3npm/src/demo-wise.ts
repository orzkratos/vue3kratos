import {GrpcWebFetchTransport} from '@protobuf-ts/grpcweb-transport'

// 导入所有客户端
import {GreeterClient} from "./rpc/v1/greeter/greeter.client";
import {RpcdemoClient} from "./rpc/rpcdemo/rpcdemo.client";
import {RpcpingClient} from "./rpc/rpcping/rpcping.client";

// 导入所有消息类型
import {HelloRequest} from "./rpc/v1/greeter/greeter";
import {
    CreateRpcDemoRequest,
    DeleteRpcDemoRequest,
    SelectRpcDemoRequest,
    UpdateRpcDemoRequest
} from "./rpc/rpcdemo/rpcdemo";
import {StringValue} from "./rpc/google/protobuf/wrappers";

console.log("=".repeat(50))
console.log('🚀 Vue3 Kratos 全接口集成演示开始')
console.log("=".repeat(50))

// 创建共享的传输实例
const demoTransport = new GrpcWebFetchTransport({
    baseUrl: "http://127.0.0.1:28000",
    meta: {
        Authorization: 'TOKEN-888',
    },
});

// 创建所有客户端实例
const greeterClient = new GreeterClient(demoTransport)
const rpcdemoClient = new RpcdemoClient(demoTransport)
const rpcpingClient = new RpcpingClient(demoTransport)

// 1. Greeter 服务演示
async function demoGreeterService() {
    console.log("\n📢 【Greeter 服务演示】")
    console.log("-".repeat(30))

    const request = HelloRequest.create({
        name: "Vue3 Kratos Integration"
    })

    try {
        const response = await greeterClient.sayHello(request, {})
        console.log('✅ Greeter 调用成功:')
        console.log(`   请求: name = "${request.name}"`)
        console.log(`   响应: message = "${response.data.message}"`)
        return true
    } catch (cause: any) {
        console.error('❌ Greeter 调用失败:', cause)
        return false
    }
}

// 2. RpcPing 服务演示
async function demoRpcPingService() {
    console.log("\n🏓 【RpcPing 服务演示】")
    console.log("-".repeat(30))

    const request = StringValue.create({
        value: "全接口演示 Ping 测试"
    })

    try {
        const response = await rpcpingClient.ping(request, {})
        console.log('✅ Ping 调用成功:')
        console.log(`   请求: value = "${request.value}"`)
        console.log(`   响应: value = "${response.data.value}"`)
        return true
    } catch (cause: any) {
        console.error('❌ Ping 调用失败:', cause)
        return false
    }
}

// 3. RpcDemo CRUD 服务演示
async function demoRpcDemoService() {
    console.log("\n📋 【RpcDemo CRUD 服务演示】")
    console.log("-".repeat(30))

    let isSuccess = true

    try {
        // 创建
        console.log('🔨 执行 Create 操作...')
        const createRequest = CreateRpcDemoRequest.create({
            code: "ALL_DEMO_001",
            name: "全接口演示项目",
            type: "integration"
        })
        const createResponse = await rpcdemoClient.createRpcDemo(createRequest, {})
        console.log(`✅ Create 成功，code: ${createResponse.data.code}`)

        // 更新
        console.log('📝 执行 Update 操作...')
        const updateRequest = UpdateRpcDemoRequest.create({
            code: createResponse.data.code,
            name: "全接口演示项目(已更新)"
        })
        const updateResponse = await rpcdemoClient.updateRpcDemo(updateRequest, {})
        console.log(`✅ Update 成功，更新后数据:`, updateResponse.data)

        // 查询
        console.log('🔍 执行 Select 操作...')
        const selectRequest = SelectRpcDemoRequest.create({
            type: "integration"
        })
        const selectResponse = await rpcdemoClient.selectRpcDemo(selectRequest, {})
        console.log(`✅ Select 成功，找到 ${selectResponse.data.demos.length} 条记录:`)
        selectResponse.data.demos.forEach((item, index) => {
            console.log(`   ${index + 1}. code:${item.code}, name:${item.name}, type:${item.type}`)
        })

        // 删除
        console.log('🗑️  执行 Delete 操作...')
        const deleteRequest = DeleteRpcDemoRequest.create({
            code: createResponse.data.code
        })
        const deleteResponse = await rpcdemoClient.deleteRpcDemo(deleteRequest, {})
        console.log(`✅ Delete 成功，删除的 code: ${deleteResponse.data.code}`)

    } catch (cause: any) {
        console.error('❌ RpcDemo CRUD 操作失败:', cause)
        isSuccess = false
    }

    return isSuccess
}

// 统计和报告函数
function printSummary(results: { service: string, success: boolean }[]) {
    console.log("\n" + "=".repeat(50))
    console.log('📊 演示结果汇总')
    console.log("=".repeat(50))

    const successCount = results.filter(r => r.success).length
    const totalCount = results.length

    results.forEach(result => {
        const status = result.success ? '✅ 成功' : '❌ 失败'
        console.log(`${status} ${result.service}`)
    })

    console.log("-".repeat(30))
    console.log(`📈 总体成功率: ${successCount}/${totalCount} (${((successCount / totalCount) * 100).toFixed(1)}%)`)

    if (successCount === totalCount) {
        console.log('🎉 恭喜！所有接口演示均成功完成！')
        console.log('🚀 Vue3 Kratos 集成工作完美！')
    } else {
        console.log('⚠️  部分接口演示失败，请检查服务状态')
    }
}

// 主演示流程
async function runCompleteDemo() {
    const results: { service: string, success: boolean }[] = []

    console.log('🔄 开始执行全接口演示流程...\n')

    // 按顺序执行所有演示
    results.push({
        service: 'Greeter 服务',
        success: await demoGreeterService()
    })

    results.push({
        service: 'RpcPing 服务',
        success: await demoRpcPingService()
    })

    results.push({
        service: 'RpcDemo CRUD 服务',
        success: await demoRpcDemoService()
    })

    // 打印汇总报告
    printSummary(results)

    console.log("\n" + "=".repeat(50))
    console.log('🏁 Vue3 Kratos 全接口集成演示结束')
    console.log("=".repeat(50))
}

// 启动完整演示
runCompleteDemo().catch(err => {
    console.error('演示过程中发生未处理的错误:', err)
})