package dict

import "fmt"

func ExampleFromWordPairsCsv() {
	words := []string{"Apfel,apple", "Birne,pear", "Banane,banana"}
	d := FromWordPairsCsv(words)

	fmt.Println(d.GetDe("Apfel"))
	fmt.Println(d.GetDe("Birne"))
	fmt.Println(d.GetDe("Banane"))
	fmt.Println(d.GetDe("Kirsche"))

	// Output:
	// {Apfel apple}
	// {Birne pear}
	// {Banane banana}
	// { }
}

func ExampleReadFileLines() {
	words := ReadFileLines("test_dictionary.csv")
	for _, w := range words {
		fmt.Println(w)
	}
	//fmt.Println(len(words))

	// Output:
	//Apfel,apple
	//Birne,pear
	//Erdbeere,strawberry
	//Wald,wood
	//Wald,forest
	//Haus,house
	//Haus,mansion
	//KatzeCat
	//Monster,monster,creature,beast
}

func ExampleFromFile() {
	d := FromFile("test_dictionary.csv")

	fmt.Println(d.Size())
	fmt.Println(d)

	fmt.Println(d.LookupAll("Apfel"))
	fmt.Println(d.LookupAll("Birne"))
	fmt.Println(d.LookupAll("Erdbeere"))
	fmt.Println(d.LookupAll("Wald"))
	fmt.Println(d.LookupAll("Haus"))
	fmt.Println(d.LookupAll("Katze"))
	fmt.Println(d.LookupAll("Monster"))

	// Output:
	// 7
	// [apple]
	// [pear]
	// [strawberry]
	// [wood forest]
	// [house mansion]
	// []
	// []
}
