
package main
import "fmt"

func main(){
	var numero int = 1
	fmt.Println(numero)
	n := 2
	fmt.Println(n)
	s := 2
	fmt.Println(s)
	b := 22
	fmt.Println(b)
	a:= 2
	fmt.Println(a)

	//changing values
	//numero = "numero"
	//causes error because cant change datatypet to string
	fmt.Println("numero default", numero)
	numero = 333 
	fmt.Println("new numero", numero)

	var  name, roll = "Aish", 1
	fmt.Println(name)
	fmt.Println(roll)
}
