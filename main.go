package main

import (
	"accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("JongEun") //account 초기화
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil { //error가 발생하면 이에 대한 처리도 개발자가 직접 해야한다. 하지 않으면 Go에서는 관여하지 않고 넘어간다.
		//log.Fatalln(err) //log.Fatalln은 Println()을 호출시켜 err 메세지를 출력하고 프로그램을 종료시킨다.
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
