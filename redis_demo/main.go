package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

var pool *redis.Pool

func init() {
	log.Println("before main ???")
	pool = &redis.Pool{
		MaxIdle:     20,  //最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
		MaxActive:   120, //最大的激活连接数，表示同时最多有N个连接
		IdleTimeout: 350, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	log.Println("after init...")
}

func connPool() {
	client := pool.Get()
	defer client.Close()

	_, err := client.Do("SET", "names", "redis-pool")
	if err != nil {
		fmt.Println("set error in connPool:", err)
		return
	}

	r, err := redis.String(client.Do("GET", "names"))
	if err != nil {
		fmt.Println("get error in connPool:", err)
		return
	}

	fmt.Println("connPool get", r)
}

func singConnectRedis() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379") //, redis.DialPassword("Itran_2430！@#3.1415926")) 未设置密码
	if err != nil {
		fmt.Printf("conn redis failed, err:", err)
		return
	}
	defer c.Close()

	// set
	_, err = c.Do("SET", "name", "redis-go")
	if err != nil {
		println("err")
		return
	}

	// get
	r, err := redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println("got err here...", err)
		return
	}
	fmt.Println("get name:", r)

	//hset
	_, err = c.Do("HSET", "names123", "redis123", "hset-value-gooooooooooooood")
	if err != nil {
		fmt.Println(err)
		return
	}
	//hget
	r, err = redis.String(c.Do("HGET", "names123", "redis123"))
	if err != nil {
		fmt.Println("hget err: ", err)
		return
	}
	fmt.Println("hget names redis: ", r)

	//exipre
	_, err = c.Do("expire", "names", 5)
	if err != nil {
		fmt.Println("expire err: ", err)
		return
	}

	// 管道
	//使用Send()，Flush()，Receive()方法支持管道化操作
	//Send向连接的输出缓冲中写入命令。
	//Flush将连接的输出缓冲清空并写入服务器端。
	//Recevie按照FIFO顺序依次读取服务器的响应
	c.Send("SET", "name1", "redis001")
	c.Send("SET", "name2", "redis002")
	c.Flush()

	v, err := c.Receive()
	fmt.Printf("v: %v, err: %v \n", v, err)

	v, err = c.Receive()
	fmt.Printf("v: %v, err: %v \n", v, err)

	//v, err = c.Receive() // 夯住，一直等待
	//fmt.Printf("v:%v,err:%v\n", v, err)
}

func main() {
	connPool()
	log.Println("not run here...???")
	singConnectRedis()
}
