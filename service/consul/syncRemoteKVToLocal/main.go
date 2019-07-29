package main

import (
	"log"

	"github.com/hashicorp/consul/api"
)

func connectRemote() *api.Client {
	// Get a new client
	config := api.DefaultConfig()
	config.Address = "192.168.199.32:8500"
	remoteClient, err := api.NewClient(config)
	if err != nil {
		log.Panicln("remote api.NewClient failed")
		panic(err)
	}

	log.Println("remote client connected")

	return remoteClient
}

func connectLocal() *api.Client {
	// Get a new client
	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	localClient, err := api.NewClient(config)
	if err != nil {
		log.Panicln("local api.NewClient failed")
		panic(err)
	}

	log.Println("local client connected")

	return localClient
}

func main() {
	// Get a new client
	remoteClient := connectRemote()
	localClient := connectLocal()

	// Get a handle to the KV API
	remoteKV := remoteClient.KV()
	localKV := localClient.KV()

	kvPairs, _, err := remoteKV.List("myapp/database", nil)
	if err != nil {
		log.Println("List failed")
		panic(err)
	}

	log.Println("client listed")

	for _, v := range kvPairs {
		// log.Printf("%s = %s\n", v.Key, v.Value)

		p := &api.KVPair{Key: v.Key, Value: v.Value}
		_, err = localKV.Put(p, nil)
		if err != nil {
			panic(err)
		}
	}

	log.Println("completed!")
}
