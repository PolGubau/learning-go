package main

import "slices"

func main() {

	// animals := [3]string{"dog", "cat", "bird"}

	// for i := 0; i < len(animals); i++ {
	// 	println(animals[i])
	// }

	animalSlice := []string{
		"dog",
		"cat",
		"bird",
	}

	animalSlice = append(animalSlice, "fish")

	animalSlice = slices.Delete(animalSlice, 0,1)

	// for i := 0; i < len(animalSlice); i++ {
	// 	println(animalSlice[i])
	// }

	for i,value := range animalSlice {
		println(i,value)
	}





}