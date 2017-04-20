package main

import "fmt"

func sum(arrays []int, ch chan int) {
    //fmt.Println(arrays)
    sum := 0
    for _, array := range arrays {
        sum += array
    }
    ch <- sum
}

func main() {
    arrayChan := make(chan int, 20)
    arrayInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
    for t := 0; t < 10; t++ {
        length := len(arrayInt)
        go sum(arrayInt[length-t:], arrayChan)
    }

    arrayResult := [10]int{0}
    for i := 0; i < 10; i++ {
        arrayResult[i] = <-arrayChan
    }
    fmt.Println(arrayResult)
}