package main
import "fmt"


func main(){
    fmt.Println(add[int32](2,3))
}
   
func add[T int16|int32|float32](a T,b T) T{
  return a+b
}

