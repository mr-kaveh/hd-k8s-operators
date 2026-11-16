package main

import (
	"errors"
	"fmt"
)

func printNamespace(namespace string) (string, error) {
	if namespace == "" {
		return "", errors.New("ns must not be empty")
	}

	return "Watching namespace : " + namespace, nil

}

func main() {

	// message, err := printNamespace("")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println(message)
	namespaces := []string{"default", "", "kube-system"}

	for _, namespace := range namespaces {
		message, err := printNamespace(namespace)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println(message)
	}

}
