package main

import (
	"log"
	"os"
	_ "searchDemo/matchers"
	"searchDemo/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
