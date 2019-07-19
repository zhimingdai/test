package GA__final

import (
      "math"
)
//产生了新的个体后 要重新 remainsource一下 更新对应的值
func (chr *Chromosome) Remainresour( ){
      //var usecpu float64
      //var usedisk float64
      nl1 := chr.nl
      var j int
      for i :=0;i<14;i++{  // i代表对应的ms的编号
            for _,j = range chr.gl[i]{  //j代表对应的具体的主机
                  // fmt.Println(j)
                  nl1[j].cpund = nl1[j].cpund - testmicro[i].cpures

                  nl1[j].disknd = nl1[j].disknd - testmicro[i].disk
            }
            // fmt.Println("这是第",i,"个任务:",nl1[j].cpund)

      }
     // for j:=0;j<=49 ;j++  {
            //fmt.Printf("主机编号为%v，剩余的资源为cpu资源为：%v，剩余的disk资源为：%v\n",j,nl1[j].cpund,nl1[j].disknd)

     // }
}

func (chr *Chromosome) Judegevalidityfirst() bool  {
      nl := chr.nl
      //var nodeuseout []int
     // var nodeuseremain []int
      for i:=0;i<80;i++{
            if nl[i].disknd <=0  || nl[i].cpund <= 0 {
                  return  false

            }
      }
      return    true

}

func (chr *Chromosome) Judegevaliditysecond() ([]int,[]int ){
      nl := chr.nl
      var nodeuseout []int
      var nodeuseremain []int
      for i:=0;i<80;i++{
            if nl[i].disknd <=0  || nl[i].cpund <= 0 {
                   nodeuseout = append(nodeuseout,i)

            }else {
                  nodeuseremain = append(nodeuseremain,i)
            }
      }
      return   nodeuseremain,nodeuseout

}
func (Chrom *Chromosome) Singlebalance( ) float64{
      nl := Chrom.nl
      var singlebalance float64
      for i:=0;i<80;i++  {
            singlebalance = singlebalance + 10 -math.Abs( (4- nl[i].cpund )/ 4 - (100 - nl[i].disknd) / 100)*10
      }
      singlebalance = singlebalance / float64(len(nl))
      return singlebalance
}

func (Chrom *Chromosome) Useageblance ( ) []float64{
      nl := Chrom.nl
      var  usagebalance []float64
      var sum float64
      //var usednode int
      for i:=0;i<80;i++{
            /*if nl[i].cpund !=4 || nl[i].disknd !=100{
                  usednode ++
            }*/

            usageres := ((4 - nl[i].cpund) / 4 + (100 - nl[i].disknd) / 100) /2
            if usageres >0.3 &&usageres <0.8{
                  usageres = usageres *0.8
            }else {
                  usageres = usageres *0.1
            }

            usagebalance = append(usagebalance,usageres)
            sum = sum + usageres
      }
      //fmt.Println("已经使用的node的个数为：",usednode)
     // fmt.Println("使用度之和",sum)
      return usagebalance
}

/*func Stddeviation(data []float64)  float64 {
      //数学期望
      var sum float64 = 0
      for _, v := range data {
            sum += v
      }
      u := float64(sum) / float64(len(data))

      //标准差
      var variance float64
      for _, v := range data {
            variance += math.Pow((v - u), 2)
      }
      stddeviation := math.Sqrt(variance / float64(len(data)))
      //fmt.Println("σ:", σ)
     // fmt.Println("μ:", μ)
      return stddeviation

}*/

func (Chrom *Chromosome ) Clusterbalance ( ) float64{
      //var a,b float64
     // nl := Chrom.nl
      a := 0.5
      b := 0.5
      singlebalance := Chrom.Singlebalance()
      usagebalance := Chrom.Useageblance()
      var sumusage float64
      for i := range usagebalance{
           sumusage = sumusage + usagebalance[i]
      }
      var clusterbalance float64
      clusterbalance = b* sumusage + a*singlebalance
      return clusterbalance

}
