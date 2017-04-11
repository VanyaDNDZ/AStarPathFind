package maze

import (
	"math/rand"
)

var celltypes = map[int]string{0:"o", 1:"v", 2:"v", 3:"v"}

func GenerateMaze(size int) [][]string  {
	rows := make([][]string, size)
	for i:=0; i<size; i++{
		cols := make([]string, size)
		for j:=0; j<size; j++{
			cols[j]= getCell()
		}
		rows[i] = cols
	}
	setEndStart(&rows, size)
	return rows
}

func setEndStart(maze *[][]string, size int){
	start := rand.Intn(size*size)
	end := rand.Intn(size*size)
	(*maze)[start/size][start % size] = "s"
	(*maze)[end/size][end % size] = "e"
}

func getCell() string {
	return celltypes[rand.Intn(3)]
}