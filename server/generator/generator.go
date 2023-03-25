package main

import (
	"bbs-go/util/simple"

	"bbs-go/model"
)

func main() {
	simple.Generate("./", "bbs-go", simple.GetGenerateStruct(&model.CheckIn{}))
}
