import {GrpcWebFetchTransport} from '@protobuf-ts/grpcweb-transport'
import {GreeterClient} from "./rpc/v1/greeter/greeter.client";
import {HelloRequest} from "./rpc/v1/greeter/greeter";
import {AxiosError} from "axios";

console.log("---")
console.log('start calling sayHello...');
console.log("---")

const demoTransport = new GrpcWebFetchTransport({
    baseUrl: "http://127.0.0.1:28000",
    meta: {
        Authorization: 'TOKEN-888',
    },
});

export const greeterClient = new GreeterClient(demoTransport)

const request = HelloRequest.create()
request.name = "kratos"

greeterClient.sayHello(request, {})
    .then((response) => {
        console.log('message:', response.data.message)
    })
    .catch((cause: AxiosError) => {
        console.log('wrong:', cause)
    })
    .finally(() => {
        console.log('done')
    })

console.log("---")
console.log('return')
console.log("---")
