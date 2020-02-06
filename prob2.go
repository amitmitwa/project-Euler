package main
import "fmt"
func main(){
	sum:=0
	max:=4000000
	a,b:=1,1
	c:=a+b
	for c<max{
		if c%2==0{
			sum+=c
		}
		a=c+b
		b=a+c
		c=a+b
	}
fmt.Println(sum)
}