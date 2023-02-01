package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw") //err 메시지를 생성할 때는 "errOOO"로 작성해야한다.

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
/*
func (a Account) Deposit(amount int) { //method 작성법 : 1. "func"과 func name 사이에 (receiver type(struct))을 작성, 2. receiver는 struct의 첫 글자를 따서 소문자로 지어야 한다.
	a.balance += amount //이 코드가 작동하지 않는 이유 : Go에서 object와 struct에 관여하는 방법 때문이다. a.account는 호출된 account 객체의 복사본이므로 값을 변경해도 실제 객체에는 반영되지 않는다.
}
*/
func (a *Account) Deposit(amount int) { //위 코드에서 복사본이 아닌 원본의 값을 변경하기 위해서는 (receiver *type)->pointer receiver를 사용해야 한다.
	a.balance += amount
}

// Balance of yout account
func (a Account) Balance() int { //객체의 값만을 읽는 method는 보안을 위해 복사본의 값을 읽도록 한다. pointer receiver는 사용하지 않는다.
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error { //인출을 할 때 인출요청보다 잔액이 적으면 인출이 되어서는 안 된다. Go에는 exception이 없어서 개발자가 직접 error handling을 해야한다.
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount

	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string { //String() method는 struct 자체를 출력할 때 출력되는 것을 설정할 수 있다. Python의 "__str__"과 동일
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
