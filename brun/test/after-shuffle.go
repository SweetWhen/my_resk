package main

import (
	"fmt"
	"myresk/infra/algo"
)

func main() {
	fmt.Printf("%v\n",
		algo.AfterShuffle(int64(10), int64(100)*100))
}
