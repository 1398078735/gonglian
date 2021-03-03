package main

import (
	"fmt"
	"tengyuanhome/chain"
)

func main(){
	block0:=chain.CreateGenesis([]byte("hello"))

	block1:=chain.NewBlock(block0.Height,block0.Hash,[]byte("hello"))
	fmt.Println(block0)
	fmt.Println(block1)
}
