package main

import "fmt"

func printNamespace(namespace string) (result string) {
	// fmt.Println("Watching namespace :", namespace)
	result = "Watching namespace : " + namespace
	return
}

func main() {

	fmt.Println(printNamespace("default"))
	fmt.Println(printNamespace("kube-system"))
}
