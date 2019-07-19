package main

import (
 "GA--final"
 "fmt"
 "gonum.org/v1/plot/plotter"
 _ "gonum.org/v1/plot/plotter"
 "math/rand"
 "time"


 // "time"
)

func main()  {

 var pop1 GA__final.Popluatin
 poptest1 := pop1.Initchrom()
 poptest2 := pop1.Initchrom()


 rand.Seed(time.Now().UnixNano())
 //poptest1  := poptest1
//poptest2 := poptest2
var dataslicefirst plotter.XYs
 var dataslicesecond plotter.XYs
for i := 0;i <=4000; i++ {

 fmt.Print("zhe shi di ",i)
popson,result1 := poptest1.Election()
//var datafirst plotter.XY
datafirst := plotter.XY{
 float64((i)),
 result1}
dataslicefirst = append(dataslicefirst,datafirst)

poptest1 = popson.Soninit()
popson222,result2 := poptest2.Electionsecond()
 datasecond := plotter.XY{
  float64((i)),
  result2}
 dataslicesecond = append(dataslicesecond,datasecond)
 poptest2 = popson222.Soninit()

}
GA__final.Plot(dataslicefirst,dataslicesecond)



}


