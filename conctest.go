package main

import (
  "fmt"
  "math/rand"
)

func sort(started chan bool, finished chan bool, list []int) {
  pivot := list[0]
  length := len(list)
  storeIndex := 0
  list[0], list[length-1] = list[length-1], list[0]
  for index, value := range list[:length-1] {
    if value < pivot {
      list[index], list[storeIndex] = list[storeIndex], value
      storeIndex += 1
    }
  }

  list[storeIndex], list[length-1] = pivot, list[storeIndex]

  if storeIndex > 1 {
    sublist := list[:storeIndex]
    started<-true
    go sort(started, finished, sublist)

  }

  if storeIndex + 2 < length {
    sublist := list[storeIndex+1:]
    started<-true
    go sort(started, finished, sublist)
  }

  finished<-true
}

func main() {
  // Generate huge list
  x := make([]int, 10000000)

  for i := 0; i<len(x); i++ {
    x[i] = rand.Intn(len(x)*2)
  }

  // Make some channels n stuff
  started := make(chan bool, len(x))
  finished := make(chan bool, len(x))

  started<-true

  // Quicksort!
  fmt.Println("Starting 'Quick' sort...")
  sort(started, finished, x[:])

  // Wait for goroutines to finish
  for {
    l := len(started)
    if l==0 {
      break
    }

    <-started
    <-finished
  }


  // Validate the list
  prev := -1
  for idx, val := range x {
    if val < prev {
      fmt.Printf("Error at index %d (%d > %d)!\n", idx, val, prev)
      return
    }

    prev = val
  }

  // We're done!
  fmt.Println("Valid sorted list!")
}
