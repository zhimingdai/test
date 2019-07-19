package GA__final

type micoservice struct {
	msreq int
	cpures float64
	disk float64
}

type proconsmap map[micoservice]proconslist

type proconslist []int


//var p1 proconslist = proconslist{}

type allmirco []micoservice
type node struct {
	cpund float64
	disknd float64
}

type Nodelist []node
type Gene  []int
type Genelist []Gene

type Chromosome struct{
	gl  Genelist  //染色体里的基因表现
	nl Nodelist       //在这种基因表现下的节点的资源情况
}
type Popluatin []Chromosome
type Chrombalance map[float64]Chromosome
type Poplutationgene []Genelist


