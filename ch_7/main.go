package main

import (
	"fmt"
)

type Score int 
type Person struct{
	FirstName string
    LastName  string
    Age       int
}
type Converter func(string) int
type TeamScores map[string] int

func (p *Person) String() string{
		p.FirstName = "sam"
	    return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

type Cou struct{
	total int
}

func (c *Cou) Inc(){
	c.total++
}

func (c Cou) Incw(){
	c.total++
}

func main(){
	person := Person{FirstName: "bob" , LastName: "Smith" ,Age: 30}
	fmt.Println(person.String())
	fmt.Println(person.FirstName)

	var c Cou
	c = Cou{total: 1}
	
	c.Incw()
	fmt.Println(c)

	c.Inc()
	fmt.Println(c)

}
