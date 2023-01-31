package something

import "fmt"

func sayBye() { //외부 package의 method 이름이 소문자로 시작하면 private하게 동작하므로 외부에서 접근할 수 없다.
	fmt.Println("Bye")
}

func SayHello() { //외부 package로 export를 해서 사용하기 위해서는 method의 이름이나 struct의 이름, 멤버의 이름이 대문자로 시작하여 public으로 동작하도록 해야한다.
	fmt.Println("Hello")
}
