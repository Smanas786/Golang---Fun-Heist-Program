package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100)
  if (eludedGuards <= 50) {
    fmt.Println("Looks like you've managed to make it part the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Caught by Guards, Plan a better disguise next time?")
  }
  fmt.Println("Is Heist on?", isHeistOn)
  if isHeistOn {
    openedVault := rand.Intn(100)
    // Let Say 30% chance of opening the vault
    if openedVault <= 70 {
      fmt.Println("Grab and GO!")
    } else {
        isHeistOn = false
        fmt.Println("Vault is too hard to break, Back out, and run!")
    }
    leftSafely := rand.Intn(5)
    fmt.Println("Is Heist on?", isHeistOn)
    if (isHeistOn) {
      switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("Caught by Security while leaving")
      case 1:
        isHeistOn = false
        fmt.Println("Caught by Cops while leaving")
      case 2:
        isHeistOn = false
        fmt.Println("Caught by worker while leaving")
      case 4:
        isHeistOn = false
        fmt.Println("Left the money bags inside vault")
      case 5:
        isHeistOn = true
        fmt.Println("Left Safely")
      }
      fmt.Println("Is Heist on?", isHeistOn)
      if isHeistOn {        
        fmt.Println("Start the gateway Car!")
        amtStolen := 10000 + rand.Intn(1000000)
        fmt.Printf("Total $%d grabbed from this successful heist!",amtStolen)
      }
    }
  }  
}
