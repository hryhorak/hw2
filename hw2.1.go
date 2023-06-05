package main

import "fmt"

type Zookeeper struct {
	Name string
}

type Animal struct {
	Species string
	Escaped bool
	Noise   string
}

type Cage struct {
	Animals []Animal
}

func main() {

	zookeeper := Zookeeper{
		Name: "Alex",
	}

	animals := []Animal{
		{Species: "Lion", Escaped: true, Noise: "Roar"},
		{Species: "Crocodile", Escaped: false, Noise: "*    *"},
		{Species: "Monkey", Escaped: true, Noise: "Scream"},
		{Species: "Elephant", Escaped: true, Noise: "Trumpet"},
		{Species: "Wolf", Escaped: true, Noise: "Howling"},
	}

	cage := Cage{
		Animals: animals,
	}

	fmt.Println("Start catching:")

	for i, animal := range animals {
		if animal.Escaped {
			animals[i].Escaped = false
			cage.Animals = append(cage.Animals, animal)
			fmt.Println(animal.Species + " say " + "'" + animal.Noise + "'")
		}
	}

	fmt.Println("----------------------------")

	fmt.Println("Animals after catching")

	for _, animal := range animals {
		fmt.Println(animal.Species, "escaped status: ", animal.Escaped)
	}

	fmt.Println("Zookeper has name ", zookeeper.Name)
}
