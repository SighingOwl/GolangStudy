package main

import (
	"fmt"
	"time"
)

func main() {
	//Bank account
	/*
		account := accounts.NewAccount("JongEun") //account 초기화
		account.Deposit(10)
		fmt.Println(account.Balance())

		err := account.Withdraw(20)
		if err != nil { //error가 발생하면 이에 대한 처리도 개발자가 직접 해야한다. 하지 않으면 Go에서는 관여하지 않고 넘어간다.
			//log.Fatalln(err) //log.Fatalln은 Println()을 호출시켜 err 메세지를 출력하고 프로그램을 종료시킨다.
			fmt.Println(err)
		}

		fmt.Println(account.Balance(), account.Owner())
		fmt.Println(account)
	*/

	//Dictionary methods
	//init dict, search
	/*
		dictionary := mydict.Dictionary{"first": "First word"}
		definition, err := dictionary.Search("first")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(definition)
		}
	*/

	//Add word
	/*
		dictionary := mydict.Dictionary{}
		word := "hello"
		definition := "Greeting"
		err := dictionary.Add(word, definition)
		if err != nil {
			fmt.Println(err)
		}
		hello, _ := dictionary.Search(word)
		fmt.Println("found", word, "definition:", hello)
		err2 := dictionary.Add(word, definition)
		if err2 != nil {
			fmt.Println(err2)
		}
	*/

	//Update word
	/*
		dictionary := mydict.Dictionary{}
		baseWord := "hello"
		dictionary.Add(baseWord, "First")
		err := dictionary.Update("abc", "Second")
		if err != nil {
			fmt.Println(err)
		}
		word, _ := dictionary.Search(baseWord)
		fmt.Println(word)
	*/

	//Delete word
	/*
		dictionary := mydict.Dictionary{}
		baseWord := "hello"
		dictionary.Add(baseWord, "First")
		dictionary.Search(baseWord)
		dictionary.Delete(baseWord)
		word, err := dictionary.Search(baseWord)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(word)
	*/

	//Goroutine
	/*
		go sexyCount("JE") // 병렬프로세싱을 사용하기 위해서는 호출할 함수 앞에 "go"를 사용해야한다.
		//go sexyCount("Porsche") //Goroutine은 main 함수가 실행되는 동안에만 실행된다. main 함수에서 호출하는 모든 함수에 "go"를 사용하면 모든 Goroutine을 동시에 실행하고 끝이 나버린다. main 함수는 Goroutine이 종료될 때까지 기다리지 않아서 그대로 main 함수가 종료된다.
		sexyCount("Porsche") // 앞에 "go"가 없는 이 함수가 위의 함수와 동시에 실행될 수 있는 이유는 main함수가 이 함수를 카운팅하고 있기 때문이다.
	*/
	/*
		go sexyCount("JE")
		go sexyCount("Porsche")
		time.Sleep(time.Second * 5) // main 함수가 종료되지 않도록 하는 장치가 있다면 위의 예시처럼 해도 된다.
	*/

	//channel
	/*
		c := make(chan bool) // channel을 만들때 make(chan "type")을 사용한다.
		people := [2]string{"JE", "William"}
		for _, person := range people {
			go isSexy(person, c)
		}
		result := <-c // channel의 메세지를 받는 방법, channel을 통해 메세지를 받을 때 main 함수는 어떤 답이 올때까지 기다리므로 main 함수는 종료되지 않는다.
		fmt.Println(result)
		fmt.Println(<-c) // 이렇게 해도 된다.
		fmt.Println(<-c) // 코드에서 실제로 동작하고 있는 Goroutine은 2개여서 channel에서 메세지를 받을 수 있지만 마지막 하나는 메세지를 계속 기다리게 되서 deadlock이 발생한다.
	*/
	/*
		c := make(chan string)
		people := [2]string{"JE", "William"}
		for _, person := range people {
			go isSexy(person, c)
		}
		fmt.Println("Waiting for messages")
		resultOne := <-c
		resultTwo := <-c
		resultThree := <-c
		fmt.Println("Received this message:", resultOne)
		fmt.Println("Received this message:", resultTwo)   // receiving messages is blocking operation
		fmt.Println("Received thsi message:", resultThree) // deadlock
	*/
	c := make(chan string)
	people := [5]string{"JE", "William", "A", "B", "C"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for ", i)
		fmt.Println(<-c) // 5개의 message receiver를 만든다고 생각하면 된다.
	}

}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, c chan string) { // channel을 인자로 사용할 때는 chan과 channel을 사용해 보낼 메세지의 형식을 지정해야한다.
	time.Sleep(time.Second * 5)
	c <- person + " is sexy" // channel로 값을 보내고 싶을 때 chan 인자에 <-를 사용해서 값을 지정한다.
}
