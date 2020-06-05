package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo // можно просто написать так, указать только тип
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		//contact: contactInfo{
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 940001,
		},
	}

	// jim.updateName("jimmy") - не работает, потому что не через указатели

	jimPointer := &jim
	jimPointer.updateName("jimmy")

	jim.print()

	/*-------------------- before 040 ----------------------------
		//alex := person{"Alex", "Anderson"} - будет присваивать согласно порядку определенному в объявлении структуры
		//alex := person{firstName: "Alex", lastName: "Aderson"}

		var alex person
		fmt.Printf("%+v", alex) // выведет с полями структуры {firstName: lastName:}

		alex.firstName = "Alex"
		alex.lastName = "Anderson"

		fmt.Printf("%+v", alex)

		//fmt.Println(alex)
	---------------------------------------------------------------*/
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

/*-------------------------------------------------------------------
Не работает потому, что не через указатели
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
---------------------------------------------------------------------*/

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName // в полученный адрес мы пишем новое значение ј
}
