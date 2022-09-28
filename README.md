```shell
# protobuf生成命令
protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative book.proto

```


# consul服务注册发现
## docker-compose搭建
```yaml
version: '3.1'

services:
  consul1:
    image: consul
    container_name: node1
    command: agent -server -bootstrap-expect=3 -node=node1 -bind=0.0.0.0 -client=0.0.0.0 -datacenter=dc1

  consul2:
    image: consul
    container_name: node2
    command: agent -server -retry-join=node1 -node=node2 -bind=0.0.0.0 -client=0.0.0.0 -datacenter=dc1
    depends_on:
      - consul1

  consul3:
    image: consul
    container_name: node3
    command: agent -server -retry-join=node1 -node=node3 -bind=0.0.0.0 -client=0.0.0.0 -datacenter=dc1
    depends_on:
      - consul1

  consul4:
    image: consul
    container_name: node4
    command: agent -retry-join=node1 -node=node4 -bind=0.0.0.0 -client=0.0.0.0 -datacenter=dc1 -ui
    ports:
      - "8500:8500"
    depends_on:
      - consul2
      - consul3
```
访问ui http://ip:8500

##  服务注册、发现
### HTTP(需要手动调用)
```go
// 服务注册、心跳
func Reg(host,name,id string,port int,tags []string)error{
	config := api.DefaultConfig()
	h := consulHost
	p := consulPort
	config.Address = fmt.Sprintf("%s:%d",h,p)
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	agentServiceRegistration := new(api.AgentServiceRegistration)
	agentServiceRegistration.Address = config.Address
	agentServiceRegistration.Port = port
	agentServiceRegistration.ID = id
	agentServiceRegistration.Name = name
	agentServiceRegistration.Tags = tags

	// 需要自己写一个健康检查api,返回{"msg":"OK"}
	severAddr := fmt.Sprintf("http://%s:%d/health", host, port)
	check := api.AgentServiceCheck{
		HTTP:     severAddr,
		Timeout:  "3s",
		// 每秒测一次
		Interval: "1s",
		// 5秒不通自动注销
		DeregisterCriticalServiceAfter: "5s",
	}
	agentServiceRegistration.Check = &check

	return client.Agent().ServiceRegister(agentServiceRegistration)
}

// 获取服务列表
func GetServerList() error{
	config := api.DefaultConfig()
	h := consulHost
	p := consulPort
	config.Address = fmt.Sprintf("%s:%d",h,p)
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	services, err := client.Agent().Services()
	if err != nil {
		return err
	}

	for k, v := range services {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("--------------------")
	}

	return nil
}

// 过滤服务
func FilterService() error{
	config := api.DefaultConfig()
	h := consulHost
	p := consulPort
	config.Address = fmt.Sprintf("%s:%d",h,p)
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	services, err := client.Agent().ServicesWithFilter("Service==accountWeb")
	if err != nil {
		return err
	}

	for k, v := range services {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("--------------------")
	}

	return nil
}
```

### GRPC
```go
func main(){
	ip := flag.String("ip","192.168.0.112","输入ip")
	port := flag.Int("port",9095,"输入端口")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d",*ip,*port)

	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server,&biz.AccountServer{})
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	grpc_health_v1.RegisterHealthServer(server,health.NewServer())
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d",
		internal.ViperConf.Consul.Host,
		internal.ViperConf.Consul.Port)
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	checkAddr := fmt.Sprintf("%s:%d",internal.ViperConf.AccountSrv.Host,internal.ViperConf.AccountSrv.Port)
	check := &api.AgentServiceCheck{
		GRPC: checkAddr,
		Timeout: "3s",
		Interval: "1s",
		DeregisterCriticalServiceAfter: "5s",
	}
	reg := &api.AgentServiceRegistration{
		Name: internal.ViperConf.AccountSrv.SrvName,
		ID: internal.ViperConf.AccountSrv.SrvName,
		Port: internal.ViperConf.AccountSrv.Port,
		Tags: internal.ViperConf.AccountSrv.Tags,
		Address: internal.ViperConf.AccountSrv.Host,
		Check: check,
	}

	err = client.Agent().ServiceRegister(reg)
	if err != nil {
		panic(err)
	}

	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
```

## Gin中调用Consul访问Grpc
```go
func AccountListHandler(c *gin.Context) {

	// 1、获取consul client
	config := api.DefaultConfig()
	consulAddr := fmt.Sprintf("%s:%d", internal.ViperConf.Consul.Host, internal.ViperConf.Consul.Port)
	config.Address = consulAddr
	consulClient, err := api.NewClient(config)
	if err != nil {
		zap.S().Error("AccountHandler,创建consul client失败:",err.Error())
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":"服务器内部错误",
		})
		return
	}

	accountSrvHost := ""
	accountSrvPort := 0

	serverList, err := consulClient.Agent().ServicesWithFilter("Service==accountSrv")
	if err != nil {
		zap.S().Error("AccountHandler,consul获取服务列表失败:",err.Error())
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":"服务器内部错误",
		})
		return
	}
	for _, v := range serverList {
		accountSrvHost = v.Address
		accountSrvPort = v.Port
	}

	// 2. gin获取传入参数
	pageNo, _ := strconv.ParseInt(c.Query("pageNo"), 10, 32)
	pageSize, _ := strconv.ParseInt(c.Query("pageSize"), 10, 32)

	// 3.使用consul的地址
	grpcAddr := fmt.Sprintf("%s:%d",accountSrvHost,accountSrvPort)
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		s := fmt.Sprintf("AccountListHandler-GRPC拨号失败:%s", err.Error())
		log.Logger.Info(s)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	client := pb.NewAccountServiceClient(conn)
	list, err := client.GetAccountList(context.Background(), &pb.PagingRequest{
		PageNo:   uint32(pageNo),
		PageSize: uint32(pageSize),
	})
	if err != nil {
		s := fmt.Sprintf("AccountListHandler-GRPC拨号失败:%s", err.Error())
		log.Logger.Info(s)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	var accountResList []*res.Account4Res
	for _, item := range list.AccountList {
		tmp := item
		accountResList = append(accountResList, pb2Res(tmp))
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":   "ok",
		"total": list.Total,
		"data":  accountResList,
	})
}
```