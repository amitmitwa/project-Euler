package main
import "fmt"
func main(){
	a:=600851475143 
	prime(a)

	

}
func prime(n int) {
	for n % 2 ==0{
		fmt.Println(2)
		n=n/2
	}
	for i:=3; i<=n; i++{
		if n%i==0{
			fmt.Println(i)
			n=n/i
		}
	}
}