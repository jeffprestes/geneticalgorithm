package main

import "fmt"

func main() {
	newIndividual := "kelbrKfdNrD"
	indexOfItemToBeChanged := 11
	changeACromossom(indexOfItemToBeChanged, "#", newIndividual)
}

func changeACromossom(indexOfItemToBeChanged int, newCromossom string, oldIndividual string) (newIndividual string) {
	newIndividual = oldIndividual[:indexOfItemToBeChanged]
	newIndividual += newCromossom
	if indexOfItemToBeChanged < len(oldIndividual) {
		newIndividual = string(append([]byte(newIndividual), oldIndividual[indexOfItemToBeChanged+1:]...))
	}
	fmt.Printf("indexOfItemToBeChanged: [%d] - newCromossom: [%s] \nLenOldIndividual: [%d] - oldIndividual: [%s] - lenNewIndividual: [%d] - newIndividual: [%s]\n", indexOfItemToBeChanged, newCromossom, len(oldIndividual), oldIndividual, len(newIndividual), newIndividual)
	return
}
