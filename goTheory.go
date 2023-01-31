/*
Nomadcoder go 강의 따라해보기
*/

package main

import ( //vscode에서만 함수에 대한 패키지를 자동으로 추가해준다.
	"fmt"
	"strings"
)

// 1.3 Functions part 1
func multiply(a, b int) int { //JS나 Python과 달리 인자와 반환값에 type을 지정해야한다. (a int, b int)처럼 하나하나 type을 지정하거나 (a, b int)처럼 한번에 지정해도 된다.
	return a * b
}

// 1.4 Functions part 2
func lenAndUpper(name string) (int, string) { //반환값이 여러 type이 될 수 있는 점은 Go언어의 독특한 특징이다.
	return len(name), strings.ToUpper(name)
}

func lenAndUpper_naked(name string) (length int, uppercase string) { //반환값을 미리 선언해서 반환하는 것을 naked return이라고 한다.
	defer fmt.Println(("I'm done")) //함수의 실행이 끝난 후에 실행되는 코드
	length = len(name)
	uppercase = strings.ToUpper(name)
	return //이곳에 반환할 변수를 작성하지 않아도 된다. 작성해도 정상적으로 동작한다.
}

func repeatMe(words ...string) { //인자 type앞에 "..."을 붙이면 해당 type의 인자를 여러 개 받을 수 있다.
	fmt.Println(words)
}

// 1.5 for, range, ...args
func superAdd(numbers ...int) int { //Go에서는 반복문을 사용할 때 for문만 사용할 수 있다.
	//fmt.Println((numbers))

	/*
		for index, number := range numbers { //for 반복문을 사용할 때 range 함수를 사용해서 반복횟수를 지정할 수 있다. range함수의 반환값 중 첫번째는 index이고 두번째는 value이다. 이것을 응용하면 index를 사용하지 않고 value에 접근할 수도 있다.
			fmt.Println(index, number)
		}

		for i := 0; i < len(numbers); i++ { //C언어에 사용하는 방법으로 for loop 횟수를 지정할 수도 있다.
			fmt.Println(numbers[i])
		}
	*/
	total := 0
	for _, num := range numbers {
		total += num
	}

	return total
}

// 1.6 If with a Twist
func canIDrink(age int) bool {
	if koreanAge := age - 2; koreanAge < 18 { //variable expression : if-else의 조건에서만 사용하기 위한 변수를 생성하여 사용할 수 있다.
		return false
	}
	return true
}

// 1.7 Switch
func canIDrink_sw(age int) bool { //switch는 C나 Java처럼 사용하면 된다.
	/*
		switch {
		case age < 18:
			return false
		case age == 18:
			return true
		case age > 50:
			return false
		}
		return false
	*/
	switch koreanAge := age - 2; koreanAge { //if-else처럼 variable expression을 사용할 수 있다.
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

// 1.11 Structs
type person struct { //struct는 객체에 대한 구조를 정의한다.
	name    string
	age     int
	favFood []string
}

func main() {
	//1.2 Variables and Constants
	name := "JongEun" // == var name string = "JongEun" -> :=를 사용하면 go가 적절한 type을 찾아준다.
	fmt.Println(name)

	//1.3 Functions part 1
	fmt.Println(multiply(2, 2))

	//1.4 Functions part 2
	totalLength, upperName := lenAndUpper("JongEun")
	fmt.Println(totalLength, upperName)
	totalLength2, _ := lenAndUpper("William") //함수의 반환값이 여러개 일 때 무시하고 싶은 값이 있으면 "_"를 사용한다.
	fmt.Println(totalLength2)
	totalLength3, upperName3 := lenAndUpper_naked("Audi A6 e-tron")
	fmt.Println(totalLength3, upperName3)

	repeatMe("B&O", "Sony", "Desker", "Lularlab", "Audi")

	//1.5 for, range, ...agrs
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	//1.6 If with a Twist
	fmt.Println(canIDrink(16))
	//1.7 Switch
	fmt.Println(canIDrink_sw(18))

	//1.8 Pointers
	a := 2
	b := a //주소 복사가 아닌 값만 가져간다.
	a = 10
	fmt.Println(a, b)

	c := 2
	d := &c
	c = 5
	*d = 20           //*를 사용해서 주소값에 있는 값을 확인하여 변경할 수 있다.
	fmt.Println(c, d) //주소를 출력하고 싶으면 C나 Java처럼 "&"를 붙이면 된다.
	fmt.Println(&c, d)
	fmt.Println(*d) //"*"는 see through하는 느낌이라고 생각하면 된다. 메모리주소에 저장된 값을 확인할 수 있다.

	//1.9 Arrays and Slices
	names := [5]string{"JongEun", "William"} //Go에서 array를 생성하려면 먼저 array의 크기와 type을 명시해야한다.
	names[2] = "Borum"
	names[3] = "Dongwol"
	names[4] = "Ilu"
	fmt.Println(names)

	cars := []string{"Taycan", "911", "718", "A6"} //Go에서 slice는 기본적으로 array이지만 크기가 정해져 있지 않다.
	cars = append(cars, "ev6")                     //Append를 사용해서 slice에서 요소를 추가할 수 있다. slice = append(slice, value)
	fmt.Println(cars)

	//1.10 maps
	car := map[string]string{"name": "A6", "Productor": "Audi"} //map을 생성할 때 key의 type은 []안에 지정하고 value의 type은 []옆에 지정한다.
	fmt.Println(car)

	for key, value := range car { //map의 요소들을 for loop와 range를 사용해서 출력할 수 있다.
		fmt.Println(key, value)
	}

	//1.11 Structs
	favFood := []string{"chicken", "pasta"}
	William := person{name: "William", age: 26, favFood: favFood} //Struct를 사용해서 객체를 만들 때 어떤 요소에 값을 넣을껀지 지정하는 것이 더 깔끔하다.
	fmt.Println(William)
}
