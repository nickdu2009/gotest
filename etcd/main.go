package main

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.33.10:2379"},
		DialTimeout: 5 * time.Second,
	})

	checkErr(err)

	kv := clientv3.NewKV(cli)

	putResp, err := kv.Put(context.TODO(), "/test/a", "something")
	checkErr(err)
	log.Printf("put [/test/a] response %v\n", putResp)
	_, err = kv.Put(context.TODO(), "/test/b", "another")
	checkErr(err)
	_, err = kv.Put(context.TODO(), "/testxxx", "干扰")

	getResp, err := kv.Get(context.TODO(), "/test/a")
	checkErr(err)
	log.Printf("get [/test/a] response %v\n", getResp)

	rangeResp, err := kv.Get(context.TODO(), "/test/", clientv3.WithPrefix())
	checkErr(err)
	log.Printf("range get [/test/] response %v\n", rangeResp)

	lease := clientv3.NewLease(cli)

	grantResp, err := lease.Grant(context.TODO(), 10)
	checkErr(err)
	log.Printf("lease grant response %v", grantResp)
	_, err = kv.Put(context.TODO(), "/test/expireme", "gone...", clientv3.WithLease(grantResp.ID))
	checkErr(err)
}
