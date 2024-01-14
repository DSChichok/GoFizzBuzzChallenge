package main
import ("fmt")

func main(){
    var fizzy = 3
    var buzzy = 5
    
    var i = 0
    for i = 0; i <= 100; i++{
        if i % fizzy == 0 && i % buzzy == 0 {
            fmt.Println("FizzBuzz")
        } else if i % fizzy == 0 {
            fmt.Println("Fizz")
        } else if i % buzzy == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}
