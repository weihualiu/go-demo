package main

import "fmt"
//import "time"


type PersonHandler interface {
	Batch(origs <-chan Person) <-chan Person
	Handle(orig *Person)
}

type PersonHandlerImpl struct {}

type Person string

func main() {
	handler := getPersonHandler()
	origs := make(chan Person, 100)
	dests := handler.Batch(origs)
	fetchPerson(origs)
	sign := savePerson(dests)
	<-sign
}

func (handler PersonHandlerImpl) Batch(origs <-chan Person) <-chan Person {
	dests := make(chan Person, 100)
	go func(){
		for{
			p, ok := <-origs
			if !ok {
				close(dests)
				break
			}
			handler.Handle(&p)
			dests <- p
		}
	}()
	return dests
}

func (handler PersonHandlerImpl) Handle(orig *Person) {
	fmt.Printf("%s\n", string(*orig))	
}

func getPersonHandler() PersonHandler {
	return new(PersonHandlerImpl)	
}

func fetchPerson(origs chan<- Person) {
	go func(){
		for i := 0; i< 100; i++ {
			p := new(Person)
			*p = Person("person." + string(i))
			origs <- *p
		}
		close(origs)
	}()
}

func savePerson(dest <-chan Person) <-chan byte {
	sign := make(chan byte)
	go func(){
		for{
			p, ok := <-dest
			if !ok {
				sign <- 0
				close(sign)
				break
			}
			fmt.Printf("save %s\n", string(p))
		}
	}()
	return sign
}


