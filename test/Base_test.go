package Base

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)


/*
单元测试
go test
go test -v -cover
 */
func TestAdd(t *testing.T) {
	var TrueResult  = 3
	result := Add(1,2)
	GetMode(1)
	GetMode(2)
	GetMode(3)
	fmt.Println(result)
	if result != TrueResult{
		t.Fatal("Add() 的计算结果错误")
	}
}

/*
压力测试
go test -test.bench=".*"
 */
func BenchmarkAdd(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	for i:=0;i<b.N;i++{
		randA := rand.Intn(200)
		randB := rand.Intn(500)
		res := Add(randA,randB)
		if res != randA+randB{
			b.Fatal("错误...")
		}
	}
}
