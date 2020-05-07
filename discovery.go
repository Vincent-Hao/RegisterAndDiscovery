package main

import (
    "fmt"
    "github.com/hashicorp/consul/api"
)

func main() {
    config := api.DefaultConfig()
    config.Address = "127.0.0.1:8500"
    client,err := api.NewClient(config)
    if err != nil{
        panic(err)
    }
    
    service,_,err := client.Health().Service("TranJson","tranjson",true,nil)
    if err != nil{
        panic(err)
    }
    fmt.Println(service[0].Service.Service)
}
