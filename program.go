package main

import (
    "flag"
    "fmt"
)

var (
    n = flag.Int("n", 0, "input n:")
    m = flag.Int("m", 0, "input m:")
)

func main() {
    flag.Parse()

    if *n < 0 || *m < 0 {
       fmt.Println("Invalid input, please re-run the program.")
       return
    }

    if *n == 0 && *m == 0 {
       fmt.Println("没有红色方块")
       return
    }

    n1 := *n/3
    m1 := *m/3

    n2 := *n%3
    m2 := *m%3

    result := n1*m1*3

    result += n2*m1 + n1*m2

    if (n2 != 0 || m2 != 0) && n1 > 0 && m1 > 0{
       if n2 + m2 < 3{
          result ++
       }else {
          result += 2
       }
    }

    fmt.Println(fmt.Sprintf("有%d个红色方块", result))
}
