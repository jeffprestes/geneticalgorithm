package main

import "fmt"

func main() {
	newIndividual := "kelbrKfdNrD"
	indexOfItemToBeChanged := 11
	changeACromossomT(indexOfItemToBeChanged, "#", newIndividual)
	generateNewIndividualT(float64(0.5), newIndividual, "qwertyuiopa", float64(0.1))
}

func changeACromossomT(indexOfItemToBeChanged int, newCromossom string, oldIndividual string) (newIndividual string) {
	newIndividual = oldIndividual[:indexOfItemToBeChanged]
	newIndividual += newCromossom
	if indexOfItemToBeChanged < len(oldIndividual) {
		newIndividual = string(append([]byte(newIndividual), oldIndividual[indexOfItemToBeChanged+1:]...))
	}
	fmt.Printf("indexOfItemToBeChanged: [%d] - newCromossom: [%s] \nLenOldIndividual: [%d] - oldIndividual: [%s] - lenNewIndividual: [%d] - newIndividual: [%s]\n", indexOfItemToBeChanged, newCromossom, len(oldIndividual), oldIndividual, len(newIndividual), newIndividual)
	return
}

func generateNewIndividualT(crossover float64, bestOldIndividual string, oldSecondBestIndividual string, mutationIndex float64) (newCreatedIndividual string) {
	max := len(bestOldIndividual)
	posA := int(round(float64(max)*crossover, 0))
	posB := max - posA
	fmt.Printf("[generateNewIndividual] Crossover: %+v - BestOldIndividual: %+v - mutationIndex: %+v - posA: %+v - posB: %+v\n", crossover, bestOldIndividual, mutationIndex, posA, posB)
	// i := 0
	// for i < posA {
	// 	pos := rand.Intn(posA)
	// 	fmt.Println("ch: ")
	// 	newCreatedIndividual += string(bestOldIndividual[pos])
	// 	i++
	// }
	// i = 0
	// for i < posB {
	// 	pos := rand.Intn(posB)
	// 	fmt.Println("ch: ")
	// 	newCreatedIndividual += string(oldSecondBestIndividual[pos])
	// 	i++
	// }
	//fmt.Println("[generateNewIndividual] bestOld: ", bestOldIndividual[:posA-1], " - old: ", oldIndividual[posB:])
	newCreatedIndividual = bestOldIndividual[:posA]
	newCreatedIndividual += oldSecondBestIndividual[:posB]
	fmt.Printf("[generateNewIndividual] New Individual before mutation: [%s] len: [%d]\n", newCreatedIndividual, len(newCreatedIndividual))
	//newCreatedIndividual = mutateAnIndividual(mutationIndex, characteristicsSet, newCreatedIndividual)
	//fmt.Printf("[generateNewIndividual] New Individual after mutation: [%s] len: [%d]\n", newCreatedIndividual, len(newCreatedIndividual))
	return
}

func round(v float64, decimals int) float64 {
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	return float64(int((v*pow)+0.5)) / pow
}
