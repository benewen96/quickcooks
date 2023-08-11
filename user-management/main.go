package main

import (
	"fmt"
)

func main() {
	context, err := newUserManagementContext()
	if err != nil {
		fmt.Printf("Error creating context: %v", err)
		return
	}
	err = newRouter(context).Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		fmt.Printf("Error starting router: %v", err)
		return
	}
}
