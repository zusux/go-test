package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main()  {
	//SetRedis("set","a",map[string]interface{}{"id":1,"name":"lisi"})
	res,err := GetRedis("get","a")

	fmt.Println(res,err)
}


func GetRedis(command,key string)(res interface{},err error){
	conn ,err := redis.Dial("tcp",":6379")
	if err != nil{
		return
	}
	rel,err := redis.Bytes(conn.Do(command,key))
	if err != nil{
		return
	}

	res,err = ByteDecode(rel)
	return
}

func SetRedis(command,key string,args interface{})(res interface{},err error){
	conn ,err := redis.Dial("tcp",":6379")
	b,err := ByteEncode(args)
	if err != nil{
		return
	}
	rel,err := redis.Bytes(conn.Do(command,key,b.Bytes()))
	if err != nil{
		return
	}
	res,err = ByteDecode(rel)
	return
}


func ByteEncode(v interface{})(bytes.Buffer,error){
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(v)
	return buffer,err
}
func ByteDecode(byt []byte)(src map[string]interface{},err error){
	dec := gob.NewDecoder(bytes.NewReader(byt))
	err = dec.Decode(&src)
	return
}
