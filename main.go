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
	go sexyCount("JE")
	go sexyCount("Porsche")
	time.Sleep(time.Second * 5) // main 함수가 종료되지 않도록 하는 장치가 있다면 위의 예시처럼 해도 된다.

}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
