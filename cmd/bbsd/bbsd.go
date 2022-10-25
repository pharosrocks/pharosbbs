package main

import (
	"github.com/joho/godotenv"
	"github.com/pharosrocks/pharosbbs/bbs"
)

func main() {
	var bbs = bbs.NewServer()
	godotenv.Load()
	bbs.ListenAndServe(":80")

}
