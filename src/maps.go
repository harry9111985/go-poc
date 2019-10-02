package main

import "fmt"

func main() {

	//Initializing a map with short hand format
	atheleteHeights := map[string]int{
		"Harish": 170,
		"Suresh": 165,
		"Ramesh": 180,
	}

	fmt.Println(atheleteHeights)

	//Creating an empty map using make function
	atheleteWeights := make(map[string]float64)

	fmt.Println(atheleteWeights)

	//Syntax if the map value is present or not

	name:= "Ramesh"
	_,present := atheleteWeights[name]

	fmt.Printf("Is weight present for %s ? %t \n",name,present)

	atheleteWeights[name] = 64

	weight,weightPresent := atheleteWeights[name]

	if weightPresent {
		fmt.Printf("weight  for %s is %f \n", name, weight)
	}

	//Deleting elements in  maps
	delete(atheleteWeights,name)

	fmt.Println("Length of atheletes weight map is ",len(atheleteWeights))

	_,isPresentAfterDelete := atheleteWeights[name]

	if !isPresentAfterDelete {
		fmt.Printf("%s record deleted successfully \n",name)
	}

	//Maps are pass by reference (Reference type) . Both variables point to the same map. So change in one will reflect in other.
	newAtheleteWeights := atheleteWeights

	newAtheleteWeights[name] = 55
	fmt.Println("Old Athelete weights map after change in the new weights map is now  ",atheleteWeights)
}
