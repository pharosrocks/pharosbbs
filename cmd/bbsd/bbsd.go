package main

import (
	"github.com/pharosrocks/pharosbbs/bbs"
)

func main() {
	var bbs = &bbs.Server{}
	bbs.ListenAndServe(":8080")

}
