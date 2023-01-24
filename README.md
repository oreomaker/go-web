# go-web

this is a microservice demo in golang

reference: https://github.com/cloudwego/biz-demo

## Structure

### api gateway

```
api/
├── biz                   // 业务逻辑
│   ├── handler           // 控制层逻辑
│   ├── router            // herz生成的路由
│   └── model             // Request、Response定义
├── pkg                   // 第三方依赖封装和配置
│   ├── errno             // 统一错误状态码
│   ├── mw                // middleware
│   └── consts            // 常量
└── kitex_gen             // kitex生成的rpc客户端代码
```

### user service
```
user/
├── kitex_gen             // kitex生成
├── pkg                   // 第三方依赖封装和配置，由api复制
│   ├── errno             // 统一错误状态码
│   ├── mw                // middleware
│   └── consts            // 常量
└── handler               // rpc执行逻辑实现
```

## Problems To Be Solved
api网关和user服务位于两个mod中，常量定义同时存在于两个mod中；如果将两个模块放到一个mod中，使用herz和kitex生成的代码会有问题，需要手动修改；