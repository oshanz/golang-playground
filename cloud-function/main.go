package main

import (
	_ "local/functions/hello"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	funcframework.StartHostPort("127.0.0.1", "8080")
}
