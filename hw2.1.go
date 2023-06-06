package main

import "fmt"

type Zookeeper struct {
	Name string
	Cage Cage
}

type Animal struct {
	Species string
	Escaped bool
	Noise   string
}

type Cage struct {
	Animals []Animal
}

func (z *Zookeeper) CatchAnimal(animal Animal) {
	z.Cage.Animals = append(z.Cage.Animals, animal)
}

func main() {

	zookeeper := Zookeeper{
		Name: "Alex",
		Cage: Cage{},
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
			zookeeper.CatchAnimal(animal)
			fmt.Printf(`%s says "%s"%s`, animal.Species, animal.Noise, "\n")
		}
	}

	fmt.Println("\n", "----------------------------")

	fmt.Println("Animals after catching")

	for _, animal := range animals {
		fmt.Println(animal.Species, "escaped status: ", animal.Escaped)
	}

	fmt.Println("Zookeper has name ", zookeeper.Name)

	fmt.Printf("There are %d animals in cage", len(cage.Animals))
}
