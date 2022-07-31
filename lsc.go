// copyright (c) gamma410.win 

package main

import (
  "fmt"
  "flag"
  "strconv"
)

func main(){
  flag.Parse()
  arg := flag.Arg(0)
  i, _ := strconv.Atoi(arg)

  fmt.Println(">>>",calc(i),"m")
}

func calc(i int) int {
  // 落雷地点までの距離を返す
  data := i * 340
  return data
}
