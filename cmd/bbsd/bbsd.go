package main

import (
	"fmt"

	"github.com/pharosrocks/pharosbbs/bbs"
)

func main() {
	fmt.Println("test")

	var bbs = &bbs.Server{}
	bbs.ListenAndServe(":8080")

}
