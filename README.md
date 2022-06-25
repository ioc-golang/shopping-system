# 基于 IOC-golang 的简易电商系统示例

本项目为使用 ioc-golang 框架开发的分布式电商系统，展示了框架的依赖注入能力和运维能力。对于更详细的能力：配置文件加载、配置注入、RPC 等的能力，可以参阅本项目源代码。

在本教程中，我们会运行这个系统，验证系统提供的服务 API，最后使用 IOC 框架提供的调试能力，展示完整调用链路，性能瓶颈排查，性能优化。

## 1. 系统架构

### 1.1 架构详解

本系统主要包括六个模块，前三个模块为使用 ioc-golang 编写的业务应用。

- 接入层：shopping-ui 对接前端http请求，负责鉴权、请求下游具体服务。

- 活动业务模块：festival ，负责活动页信息的生成。在本项目中，一个活动页由一组卡片构成，卡片包含商品卡片和广告卡片。活动模块会首先调用下游的 商品/广告 推荐模块。获取 id 列表，再从数据库中拉取详情信息。

- 推荐模块：product、advertisement，负责用户画像分析，将适合被推荐的 商品/广告 id 列表回传给上游服务 festival。

  

- 可视化模块：visualization， 负责整个系统的可视化监控。在本例子中，主要展示全链路追踪能力。

- 缓存层：Redis Server

- 数据层：Mysql server

![tracing3](https://raw.githubusercontent.com/ioc-golang/ioc-golang-website/main/resources/img/shopping-system/structure.png)

### 1.2 项目代码结构

```
tree .                                                                     
.
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── README_CN.md
├── cmd
│   ├── advertisement 
│   ├── festival
│   ├── product
│   └── shopping-ui
├── configs 
├── deploy
├── go.mod
├── go.sum
├── internal
└── pkg
    ├── model
    └── service
        ├── advertisement
        ├── festival
        └── product
```

- cmd/ 包含四个业务应用的入口
- congis/ 包含整个系统所有应用的全平台配置信息。
- deploy/ 包含部署所需的信息：k8s资源定义、docker-compose 文件
- internal/ 包含鉴权模块
- pkg/model/ 包含对象定义信息，包括传输对象、数据存储对象等
- pkg/service/ 包含各模块 RPC 服务的实现、接口和业务逻辑。

## 2. 如何运行

我们提供三种方式运行本系统：

- 部署至 k8s 集群【推荐】

- 本地 docker 环境部署

  **请确保运行在 linux/amd64 机器上**

- 基于源代码运行程序。

### 2.1 部署系统至 k8s 集群

#### 2.1.1 准备工作

一个可访问公网数据的 K8S 集群，本地可通过 kubectl 命令访问集群：

```bash
% kubectl get ns
NAME                  STATUS   AGE
default               Active   7d7h
kube-node-lease       Active   7d7h
kube-public           Active   7d7h
kube-system           Active   7d7h
```

#### 2.1.2 部署系统

如您已经将本项目 clone 到本地，根目录执行 `make deploy-to-k8s `

如果您未 clone 本工程，执行如下命令安装全部资源。

```bash
kubectl apply -f=https://raw.githubusercontent.com/ioc-golang/shopping-system/main/deploy/k8s/namespace/namespace.yaml -f=https://raw.github
usercontent.com/ioc-golang/shopping-system/main/deploy/k8s/shopping-system/mysql.yaml -f=https://raw.githubusercontent.com/ioc-golang/shopping-system/main/deploy/k8s/shopping-system/r
edis.yaml  -f=https://raw.githubusercontent.com/ioc-golang/shopping-system/main/deploy/k8s/shopping-system/trace.yaml  -f=https://raw.githubusercontent.com/ioc-golang/shopping-system/main/deploy/k8s/shopping-system/shopping-system.yaml 
```

可看到日志打印：

```bash
namespace/ioc-shopping-system created
kubectl apply -f ./deploy/k8s/shopping-system
deployment.apps/mysql created
service/mysql-svc created
deployment.apps/redis created
service/redis-svc created
deployment.apps/ioc-shopping-system-festival created
service/ioc-shopping-system-festival-svc created
deployment.apps/ioc-shopping-system-product created
service/ioc-shopping-system-product-svc created
deployment.apps/ioc-shopping-system-advertisement created
service/ioc-shopping-system-advertisement-svc created
deployment.apps/ioc-shopping-system-shopping-ui created
service/ioc-shopping-system-shopping-ui-svc created
deployment.apps/jaeger-collector created
service/jaeger-collector-svc created
deployment.apps/jaeger-query created
service/jaeger-query-svc created
deployment.apps/elasticsearch created
service/elasticsearch-svc created
```

这会创建名为 ioc-shopping-system 的 namespace ，包含整个项目的所有资源。

```bash
 kubectl get pods -n ioc-shopping-system --watch
```

等待所有应用 pod 都呈现 Running 状态，期间可能由于资源启动顺序的问题出现 Error，等待一段时间即可。

```bash
NAME                                                 READY   STATUS    RESTARTS      AGE
elasticsearch-86dc78f867-dpcdp                       1/1     Running   0             53s
ioc-shopping-system-advertisement-7f4dd68b7d-bv4gg   1/1     Running   0             52s
ioc-shopping-system-festival-749ff7d44f-9ml99        1/1     Running   0             52s
ioc-shopping-system-product-b68875864-6x6nc          1/1     Running   0             52s
ioc-shopping-system-shopping-ui-5d79684c9d-fjkrm     1/1     Running   0             52s
jaeger-collector-7958d5cc64-8ck9z                    1/1     Running   2 (36s ago)   53s
jaeger-query-65967677cd-sfbqp                        1/1     Running   2 (33s ago)   53s
mysql-748bdb86fc-mqjfc                               1/1     Running   0             93s
redis-697477d557-2mxhv                               1/1     Running   0             53s
```

#### 2.1.3 验证服务

1. 开启两个终端，将必要的应用服务端口 port-forward 到本地。

   - 将 shopping-ui-svc 服务的 8080（前端）、1999（debug）端口映射到本地

   ```bash
   kubectl port-forward -n ioc-shopping-system svc/ioc-shopping-system-shopping-ui-svc 8080 1999
   Forwarding from 127.0.0.1:8080 -> 8080
   Forwarding from [::1]:8080 -> 8080
   Forwarding from 127.0.0.1:1999 -> 1999
   Forwarding from [::1]:1999 -> 1999
   ```

   - 将 jaeger-query-svc 服务的 16686（可视化ui 端口）映射到本地

   ```bash
   kubectl port-forward svc/jaeger-query-svc -n ioc-shopping-system 16686
   Forwarding from 127.0.0.1:16686 -> 16686
   Forwarding from [::1]:16686 -> 16686
   ```

2. 调用服务；

   shopping-ui 暴露的前端 API 包含一个 /festival/listCards 路由的 Get 方法，user_id 大于0 为合法，num 参数表示期望拉取的商品卡片数目。

   返回值为卡片列表，其中包含了商品卡片和广告卡片详情。

   ```bash
   % curl -i -X GET 'localhost:8080/festival/listCards?user_id=1&num=10'
   HTTP/1.1 200 OK
   Content-Type: application/json; charset=utf-8
   Date: Sat, 25 Jun 2022 09:51:32 GMT
   Content-Length: 1806
   
   {"Cards":[{"CardType":1,"ADsContent":"","Product":{"ID":3,"Name":"t-shirt","ProductType":"clothes","Price":9.15,"PictureURI":""}},{"CardType":1,"ADsContent":"","Product":{"ID":1,"Name":"shoes","ProductType":"clothes","Price":120,"PictureURI":""}},{"CardType":1,"ADsContent":"","Product":{"ID":3,"Name":"t-shirt","ProductType":"clothes","Price":9.15,"PictureURI":""}},{"CardType":1,"ADsContent":"","Product":{"ID":3,"Name":"t-shirt","ProductType":"clothes","Price":9.15,"PictureURI":""}},{"CardType":1,"ADsContent":"","Product":{"ID":2,"Name":"pen","ProductType":"usage","Price":80,"PictureURI":""}},{"CardType":1,"ADsContent":"","Product":{"ID":1,"Name":"shoes","ProductType":"clothes","Price":120,"PictureURI":""}},{"CardType":1,"ADsContent":"","Product":{"ID":2,"Name":"pen","ProductType":"usage","Price":80,"PictureURI":""}},{"CardType":1,"ADsContent":"","Product":{"ID":3,"Name":"t-shirt","ProductType":"clothes","Price":9.15,"PictureURI":""}},{"CardType":1,"ADsContent":"","Product":{"ID":2,"Name":"pen","ProductType":"usage","Price":80,"PictureURI":""}},{"CardType":1,"ADsContent":"","Product":{"ID":1,"Name":"shoes","ProductType":"clothes","Price":120,"PictureURI":""}},{"CardType":2,"ADsContent":"Can you imagine the ...","Product":{"ID":0,"Name":"","ProductType":"","Price":0,"PictureURI":""}},{"CardType":2,"ADsContent":"Look, this shirt is cheap.","Product":{"ID":0,"Name":"","ProductType":"","Price":0,"PictureURI":""}},{"CardType":2,"ADsContent":"Can you imagine the ...","Product":{"ID":0,"Name":"","ProductType":"","Price":0,"PictureURI":""}},{"CardType":2,"ADsContent":"Can you imagine the ...","Product":{"ID":0,"Name":"","ProductType":"","Price":0,"PictureURI":""}},{"CardType":2,"ADsContent":"Good sale, xxx...","Product":{"ID":0,"Name":"","ProductType":"","Price":0,"PictureURI":""}}]}
   ```

   可以看到请求正常相应，商品系统正常工作。这些数据是 product 服务在启动时初始化至数据库的。

### 2.2 使用 docker 部署

#### 2.2.1 准备工作

一台运行了 docker server、安装了 docker 、docker-compose 的 linux/amd64 机器。

```bash
% docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
% docker-compose version
docker-compose version 1.18.0, build 8dd22a9
...
```

#### 2.2.2 部署系统

```bash
% wget https://raw.githubusercontent.com/ioc-golang/shopping-system/main/deploy/docker-compose/docker-compose.yaml && docker-compose -f ./docker-compose.yaml up -d
...
2022-06-25 18:09:31 (32.3 KB/s) - ‘docker-compose.yaml’ saved [2585/2585]
Creating ioc-shopping-system-advertisement ... doneone
Creating network "root_shopping-system" with the default driver
Creating jaeger-collector ... done
Creating ioc-shopping-system-festival ... 
Creating redis ... 
Creating ioc-shopping-system-shopping-ui ... 
Creating root_elasticsearch_1                 ... 
Creating ioc-shopping-system-product ... 
Creating ioc-shopping-system-advertisement ... 
Creating jaeger-collector ... 
Creating query ... 
```

使用 `docker ps` 命令查看正在运行的进程。如果部分容器 STATUS 不为 up，则尝试通过 `docker restart` 手动重启容器，确保容器最终都正常运行。

#### 2.2.3 验证服务

```bash
curl -i -X GET 'localhost:8080/festival/listCards?user_id=1&num=10'
```

即可拉取到商品卡片详情列表。

### 2.3 使用 goland 基于代码运行单个应用

Clone 代码到本地，使用 Goland 打开。

点击 Goland 右上角运行配置，可看到可以运行的四个进程。点击 运行/调试 ，即可运行代码。

## 3. 使用 IOC-golang 的调试能力

上一节我们成功将系统部署在 k8s 或者 docker ，并验证了业务能力，我们也已经将必要的端口映射到了 localhost。

在本节，我们使用框架提供的调试能力，可视化服务接口和调用链路，排查性能瓶颈，优化系统性能。

在开始之前，您需要安装最新版本 iocli 工具。

```bash
% go install github.com/alibaba/ioc-golang/iocli@master
% iocli 
hello
```

### 3.1 查看应用接口和方法

前面的介绍中，我们已经将系统接入层模块：shopping-ui 应用的8080、1999端口映射到了本地。

```bash
% iocli list
github.com/alibaba/ioc-golang/extension/autowire/rpc/protocol/protocol_impl.IOCProtocol
[Invoke Export]

github.com/ioc-golang/shopping-system/internal/auth.Authenticator
[Check]

github.com/ioc-golang/shopping-system/pkg/service/festival/api.serviceIOCRPCClient
[ListCards ListCachedCards]

```

可查看到这个应用的全部接口和对应方法列表。

### 3.2 监听调用参数

我们监听鉴权接口的 Check 方法的调用：

```bash
iocli watch github.com/ioc-golang/shopping-system/internal/auth.Authenticator Check
```

新开一个终端，发起调用

```bash
curl -i -X GET 'localhost:8080/festival/listCards?user_id=1&num=10'
```

可查看到被监听方法的调用参数和返回值，user id 为1，鉴权成功。

```bash
 % iocli watch github.com/ioc-golang/shopping-system/internal/auth.Authenticator Check
========== On Call ==========
github.com/ioc-golang/shopping-system/internal/auth.Authenticator.Check()
Param 1: (int64) 1

========== On Response ==========
github.com/ioc-golang/shopping-system/internal/auth.Authenticator.Check()
Response 1: (bool) true

```

### 3.3 排查性能瓶颈

#### 性能问题可视化

在本系统中，存在大量数据库访问的操作，会造成极大的性能损耗。我们可以可视化这个问题。

使用 `iocli trace` 方法可以发起针对一个接口请求的全链路追踪，我们以 shopping-ui 服务的接入层，通过 RPC 调用下游客户端接口：github.com/ioc-golang/shopping-system/pkg/service/festival/api.serviceIOCRPCClient 的 ListCards 方法为例。

```bash
% iocli trace github.com/ioc-golang/shopping-system/pkg/service/festival/api.serviceIOCRPCClient ListCards
Tracing data is sending to jaeger-collector-svc:14268
```

以k8s部署为例，从日志中可看到，链路追踪到数据都发送至了 jaeger-collector-svc:14268 

我们尝试发起几次调用

```bash
curl -i -X GET 'localhost:8080/festival/listCards?user_id=1&num=10'
```

浏览器打开 localhost:16686，该端口已经在前面的介绍中映射至了 jaeger-query 应用。点击 FindTraces，可看到近期的调用 trace 信息：

![tracing3](https://raw.githubusercontent.com/ioc-golang/ioc-golang-website/main/resources/img/shopping-system/tracing1.png)

几次调用，耗时大致为22ms，可以看到性能是很差的。进入详情，我们可以观测到横跨四个应用的，方法粒度的完整调用链路和耗时，可以很容易地找到耗时瓶颈，为 festival 应用获取完 id 后，从数据库获取卡片详情的 First 方法。

![tracing3](https://raw.githubusercontent.com/ioc-golang/ioc-golang-website/main/resources/img/shopping-system/tracing2.png)

#### 性能优化可视化

我们尝试为卡片详情增加缓存，在本系统中，已经实现了这一策略。我们只需要调用 shopping-ui 的另一个方法 ListCachedCards 就可以感受到差别了。

发起针对 ListCachedCards 方法的追踪

```bash
% iocli trace github.com/ioc-golang/shopping-system/pkg/service/festival/api.serviceIOCRPCClient ListCachedCards
```

多次调用 ListCachedCards 方法

```bash
curl -i -X GET 'localhost:8080/festival/listCachedCards?user_id=1&num=10'
```

刷新页面，进入最近的调用详情

![tracing3](https://raw.githubusercontent.com/ioc-golang/ioc-golang-website/main/resources/img/shopping-system/tracing3.png)

可看到总耗时已经缩短到了8ms，详情信息尝试从 Redis 获取 ，整个系统的性能瓶颈已经不再是详情获取，而是商品/广告推荐模块。为追求高性能，可以进一步优化。

## 4. 总结

本项目展示了使用 ioc-golang 开发的应用系统，主要展示了依赖注入能力和接口运维能力。

针对运维能力，框架可以在业务无侵入的情况下，为所有的接口封装代理层，提供可扩展的运维能力。本例子中的分布式场景下，跨进程，跨组件的方法粒度全链路追踪，正是这一能力的体现。