package main

import (
	"fmt"
	"github.com/IvanKr8/zylo/internal/container/engine"
	"github.com/IvanKr8/zylo/internal/zyloFile/finder"
	"github.com/IvanKr8/zylo/internal/zyloFile/parser"
	"log"
)

func main() {
	zyFile, err := finder.ZyloFinder("ZyloFile")
	if err != nil {
		log.Fatalf("Error finding zylo file: %v", err)
	}

	fmt.Printf("Found zylo file: %s\n", zyFile)

	config, err := parser.ZyloParser(zyFile)
	if err != nil {
		log.Fatalf("Error parsing ZyloFile: %v", err)
	}

	if err = engine.CreateContainer(config); err != nil {
		log.Fatalf("Error executing Zylo commands: %v", err)
	}
}
