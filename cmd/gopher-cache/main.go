package main

import (
	"fmt"

	// "github.com/MaksMakarskyi/gopher-cache/internal/gopherobject"
	"github.com/MaksMakarskyi/gopher-cache/internal/ops/stringops"
	"github.com/MaksMakarskyi/gopher-cache/internal/storage"
)

func main() {
	s := storage.NewStorage()

	fmt.Println(stringops.Get(s, "foo"))

}
