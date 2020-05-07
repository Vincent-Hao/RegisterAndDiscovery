package main

import (
    "fmt"
    "github.com/hashicorp/consul/api"
    "net/http"
)

func test(response http.ResponseWriter, request *http.Request){
    response.Write([]byte("test service01 successful!"))
    fmt.Println(request.URL,request.Host,request.RemoteAddr,request.RequestURI)
}

func main(){
    config := api.DefaultConfig()
    config.Address = "127.0.0.1:8500"
    client,err := api.NewClient(config)
    if err != nil{
        panic(err)
    }
    
    //注册服务信息
    registration := new(api.AgentServiceRegistration)
    registration.Address = "127.0.0.1"
    registration.ID = "service01"
    registration.Name = "TranJson"
    registration.Port = 9999
    registration.Tags = []string{"tranjson"}
    
    //健康检查信息
    check := new(api.AgentServiceCheck)
    check.HTTP = fmt.Sprintf("http://%s:%d",registration.Address,registration.Port)
    check.Timeout = "5s"
    check.Interval = "5s"
    check.DeregisterCriticalServiceAfter = "30s"
    
    registration.Check = check
    
    err = client.Agent().ServiceRegister(registration)
    if err != nil{
        panic(err)
    }
    
    //http service
    http.HandleFunc("/",test)
    err = http.ListenAndServe(":9999",nil)
    if err != nil{
        panic(err)
    }
}