package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Pseudo codigo do algoritmo genetico
 Define as categorias ou o conjunto de dados que limitam
 Define o fitness ou com o que se vai comparar
 Define a carga genetica que será usada de cada item para gerar novos individuos (crossover)
 Define o percentual aceitável para que o sistema dê como concluído (distancia Hamming)
 Define o tamanho da população
 Define o máximo de gerações
 Define a taxa de mutacao
 Cria uma população original
 Avalia população com base no fitnes, ou seja, o quanto o resultado esta próximo do que se espera.
	 Se for maior ou igual a distancia hamming finaliza o programa
	 Se for menor, gera-se novos individuos
		 Se for um algoritmo elitista, armazena os dois que mais se aproximaram e os usam para gerar os novos individuos
		 Se não for um algoritmo elitista, gera-se individuos com dois elementos aleatórios
		 Daí gera-se novos individuos com base na taxa de crossover e na mutacao e em quantidade definida do tamanho da população
		 Adiciona mais um ao contador de gerações passadas
*/

func main() {

	var bestScoreMemory int
	var bestIndividualMemory string
	var hammingInNumberOfDigits int
	var posHourGlass int

	/*
		PARAMETERS
	*/
	characteristicsSet := "abcdefghijklmnopqrstuvxzABCDEFGHIJLKMNOPQRSTUVXZ "
	fitness := "Hello World"
	crossover := 0.5
	mutationIndex := 0.2
	populationSize := 2000
	numGeneration := 0
	maxGenerations := 100000
	strongestSurvive := false
	isolatedPopulation := false
	reinforce := true
	hamming := 100

	/*
		TIMING
	*/
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	t1 := time.Now().In(loc)
	fmt.Println("Starting processing at: ", t1.Format("2006-01-02 15:04:05"))

	hammingInNumberOfDigits = int(round(float64(len(fitness)*hamming/100), 0))
	//fmt.Println("Hamming in digits is: ", hammingInNumberOfDigits)

	pop := generateNewPopulation(len(fitness), characteristicsSet, populationSize)
	numGeneration++

	for numGeneration <= maxGenerations {
		bestIndividual, secondBestIndividual, bestScore := calculatePopulationScore(fitness, pop)
		if bestScore > bestScoreMemory {
			fmt.Printf("\n\n ===== Evaluation results ===== \n\n Best Individual: [%s]\n Best score: [%d]\n Generations: [%d]\n\n", bestIndividual, bestScore, numGeneration)
			bestScoreMemory = bestScore
			bestIndividualMemory = bestIndividual
		}

		if reinforce && bestScore < bestScoreMemory {
			bestIndividual = bestIndividualMemory
		}

		if bestScore == hammingInNumberOfDigits {
			fmt.Println("A maquina achou!!!")
			fmt.Println("Esse é o texto: ", bestIndividual)
			fmt.Printf("Numero de gerações necessárias: [%d]\n\n", numGeneration)
			return
		}
		pop = generateMutatedPopulation(crossover, strongestSurvive, isolatedPopulation, bestIndividual, secondBestIndividual, pop, mutationIndex, characteristicsSet)
		numGeneration++
		posHourGlass++
		if posHourGlass == 50 {
			posHourGlass = 0
			fmt.Print(".")
		}
	}
	t2 := time.Now().In(loc)
	dif := t2.Sub(t1).Minutes()
	fmt.Printf("\n\n ===== End results ===== \n\n Finishing processing at: [%s]\n Time elapsed: [%v]\n Best Individual: [%s]\n Best score: [%d]\n\n", t2.Format("2006-01-02 15:04:05"), dif, bestIndividualMemory, bestScoreMemory)
}

func generateNewPopulation(fitnessSize int, characteristicsSet string, populationSize int) (population []string) {
	for len(population) <= populationSize {
		var strTemp string
		for len(strTemp) < fitnessSize {
			maxPos := len(characteristicsSet)
			pos := rand.Intn(maxPos)
			strTemp += characteristicsSet[pos : pos+1]
		}
		population = append(population, strTemp)
	}
	return
}

func generateMutatedPopulation(crossover float64, isElitistAlgorithm bool, isIsolatedPopulation bool, oldBestIndividual string, oldSecondBestIndividual string, parentPopulation []string, mutationIndex float64, characteristicsSet string) (population []string) {
	if len(oldBestIndividual) < 2 || len(oldBestIndividual) != len(oldSecondBestIndividual) {
		panic("[generateMutatedPopulation] I couldn't create a new individual. best: " + oldBestIndividual + " - secondBest: " + oldSecondBestIndividual)
	}
	var i int
	if isElitistAlgorithm {
		population = append(population, oldBestIndividual)
		i++
	}
	if isIsolatedPopulation {
		for i < len(parentPopulation) {
			population = append(population, generateNewIndividualElitist(crossover, oldBestIndividual, oldSecondBestIndividual, mutationIndex, characteristicsSet))
			i++
		}
	} else {
		for i < len(parentPopulation) {
			population = append(population, generateNewIndividual(crossover, oldBestIndividual, parentPopulation[i], mutationIndex, characteristicsSet))
			i++
		}
	}
	return
}

func generateNewIndividual(crossover float64, bestOldIndividual string, oldIndividual string, mutationIndex float64, characteristicsSet string) (newCreatedIndividual string) {
	if len(bestOldIndividual) < 2 || len(bestOldIndividual) != len(oldIndividual) {
		panic("[generateNewIndividual] I couldn't create a new individual. best: " + bestOldIndividual + " - secondBest: " + oldIndividual)
	}
	//fmt.Printf("[generateNewIndividual] Crossover: %+v - BestOldIndividual: %+v - oldIndividual: %+v - mutationIndex: %+v - characSet: %+v\n", crossover, bestOldIndividual, oldIndividual, mutationIndex, characteristicsSet)
	max := len(bestOldIndividual)
	posA := int(round(float64(max)*crossover, 0))
	posB := max - posA
	temp := ""
	//fmt.Printf("[generateNewIndividual] Crossover: %+v - BestOldIndividual: %+v - mutationIndex: %+v - posA: %+v - posB: %+v\n", crossover, bestOldIndividual, mutationIndex, posA, posB)
	i := 0
	for i < posA {
		pos := rand.Intn(posA)
		//fmt.Println("pos: ", pos)
		//fmt.Println("ch: ", string(bestOldIndividual[pos]))
		temp += string(bestOldIndividual[pos])
		i++
	}
	i = 0
	for i < posB {
		pos := rand.Intn(posB)
		//fmt.Println("pos: ", pos)
		//fmt.Println("ch: ", string(oldIndividual[pos]))
		temp += string(oldIndividual[pos])
		i++
	}
	i = 0
	for i < max {
		pos := rand.Intn(max)
		//fmt.Println("pos: ", pos)
		//fmt.Println("ch: ", string(oldIndividual[pos]))
		newCreatedIndividual += string(temp[pos])
		i++
	}
	//fmt.Printf("[generateNewIndividual] New Individual before mutation: [%s] len: [%d]\n", newCreatedIndividual, len(newCreatedIndividual))
	newCreatedIndividual = mutateAnIndividual(mutationIndex, characteristicsSet, newCreatedIndividual)
	//fmt.Printf("[generateNewIndividual] New Individual after mutation: [%s] len: [%d]\n", newCreatedIndividual, len(newCreatedIndividual))
	return
}

func generateNewIndividualElitist(crossover float64, bestOldIndividual string, oldSecondBestIndividual string, mutationIndex float64, characteristicsSet string) (newCreatedIndividual string) {
	if len(bestOldIndividual) < 2 || len(bestOldIndividual) != len(oldSecondBestIndividual) {
		panic("[generateNewIndividualElitist] I couldn't create a new individual. best: " + bestOldIndividual + " - secondBest: " + oldSecondBestIndividual)
	}
	max := len(bestOldIndividual)
	posA := int(round(float64(max)*crossover, 0))
	posB := max - posA
	//fmt.Printf("[generateNewIndividualElitist] Crossover: %+v - BestOldIndividual: %+v - SecondBestOldIndividual: %+v - mutationIndex: %+v - posA: %+v - posB: %+v\n", crossover, bestOldIndividual, oldSecondBestIndividual, mutationIndex, posA, posB)
	i := 0
	for i < posA {
		pos := rand.Intn(posA)
		//fmt.Println("pos: ", pos)
		//fmt.Println("ch: ", string(bestOldIndividual[pos]))
		newCreatedIndividual += string(bestOldIndividual[pos])
		i++
	}
	i = 0
	for i < posB {
		pos := rand.Intn(posB)
		//fmt.Println("pos: ", pos)
		//fmt.Println("ch: ", string(oldSecondBestIndividual[pos]))
		newCreatedIndividual += string(oldSecondBestIndividual[pos])
		i++
	}
	//fmt.Printf("[generateNewIndividual] New Individual before mutation: [%s] len: [%d]\n", newCreatedIndividual, len(newCreatedIndividual))
	newCreatedIndividual = mutateAnIndividual(mutationIndex, characteristicsSet, newCreatedIndividual)
	//fmt.Printf("[generateNewIndividual] New Individual after mutation: [%s] len: [%d]\n", newCreatedIndividual, len(newCreatedIndividual))
	return
}

func mutateAnIndividual(mutationIndex float64, characteristicsSet string, oldIndividual string) (newMutatedIndividual string) {
	numItemsToChange := int(round(float64(len(oldIndividual))*mutationIndex, 0))
	numItemsChanged := 0
	newMutatedIndividual = oldIndividual
	for numItemsChanged <= numItemsToChange {
		maxPos := len(characteristicsSet)
		pos := rand.Intn(maxPos)
		item := characteristicsSet[pos : pos+1]
		pos = rand.Intn(len(oldIndividual) - 1)
		//fmt.Println("[mutateAnIndividual] Individual before: ", newMutatedIndividual)
		newMutatedIndividual = changeACromossom(pos, item, newMutatedIndividual)
		//fmt.Println("[mutateAnIndividual] Individual after: ", newMutatedIndividual)
		numItemsChanged++
	}
	return
}

func changeACromossom(indexOfItemToBeChanged int, newCromossom string, oldIndividual string) (newIndividual string) {
	newIndividual = oldIndividual[:indexOfItemToBeChanged]
	newIndividual += newCromossom
	if indexOfItemToBeChanged < len(oldIndividual) {
		newIndividual = string(append([]byte(newIndividual), oldIndividual[indexOfItemToBeChanged+1:]...))
	}
	//fmt.Printf("[changeACromossom] indexOfItemToBeChanged: [%d] - newCromossom: [%s] - LenOldIndividual: [%d] - oldIndividual: [%s] - newIndividual: [%s]\n", indexOfItemToBeChanged, newCromossom, len(oldIndividual), oldIndividual, newIndividual)
	return
}

/*
================================
=  EVALUATION FUNCTIONS
================================
*/
func calculateIndividualScore(fitness string, individual string) (score int) {
	if len(individual) != len(fitness) {
		fmt.Printf("[calculateIndividualScore] lenIndividual: [%d] - Ind: [%s] - lenFitness: [%d] - Fit: [%s]\n", len(individual), individual, len(fitness), fitness)
		return
	}
	for pos := 0; pos < len(fitness); pos++ {
		//fmt.Printf("[calculateIndividualScore] Pos: [%d] - Individual: [%s] - Fitness: [%s]\n", pos, individual[pos:pos+1], fitness[pos:pos+1])
		if individual[pos:pos+1] == fitness[pos:pos+1] {
			score++
		}
	}
	return
}

func calculatePopulationScore(fitness string, population []string) (bestIndividual string, secondBestIndividual string, bestScore int) {
	for _, individual := range population {
		tempScore := calculateIndividualScore(fitness, individual)
		if tempScore > bestScore {
			secondBestIndividual = bestIndividual
			bestScore = tempScore
			bestIndividual = individual
		}
	}
	if secondBestIndividual == "" {
		secondBestIndividual = bestIndividual
	}
	return
}

func round(v float64, decimals int) float64 {
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	return float64(int((v*pow)+0.5)) / pow
}
