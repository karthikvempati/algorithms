package routine

import (
  "testing"
  "time"
  "math/rand"
)


func TestPrintInt(t *testing.T)  {
  for i:=0;i<2;i++ {
    go PrintInt(i)
  }
  amt := time.Duration(rand.Intn(250))
  time.Sleep(time.Millisecond * amt)
}
