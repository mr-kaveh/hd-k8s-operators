package main

import (
	"errors"
	"fmt"
)

type Workload struct {
	Name            string
	DesiredReplicas int
	ActualReplicas  int
}

func printNamespace(namespace string) (string, error) {
	if namespace == "" {
		return "", errors.New("ns must not be empty")
	}

	return "Watching namespace : " + namespace, nil

}
func reconcile(workload Workload) Workload {

	workload.ActualReplicas = workload.DesiredReplicas
	return workload

}

func main() {

	workloads := []Workload{
		{Name: "web", DesiredReplicas: 3, ActualReplicas: 1},
		{Name: "api", DesiredReplicas: 2, ActualReplicas: 2},
	}

	for _, workload := range workloads {
		fmt.Println(reconcile(workload).Name, "already matches desired state")
	}
	// for _, workload := range workloads {
	// 	if workload.ActualReplicas != workload.DesiredReplicas {
	// 		fmt.Println(workload.Name, " need reconciliation")
	// 	} else {
	// 		fmt.Println(workload.Name, " already matches desired state")
	// 	}

	// }

}
