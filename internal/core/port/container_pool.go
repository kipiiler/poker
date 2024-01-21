package main

type ContainerPool struct {
	poolSize int
	availableContainers chan string
	inUseContainers map[string]bool
	mu 


}