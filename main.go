package main

import (
	"fmt"
	"mydict"
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
}
