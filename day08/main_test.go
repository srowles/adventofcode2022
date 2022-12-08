package main

import (
	"fmt"
	"testing"
)

func TestTreeFinder(t *testing.T) {
	input := `30373
25512
65332
33549
35390`
	forest := parse(input)
	print(forest)
	fmt.Println(visibleTrees(forest))
	t.FailNow()
}

func TestScenicScore(t *testing.T) {
	input := `30373
25512
65332
33549
35390`
	forest := parse(input)
	print(forest)
	fmt.Println()
	fmt.Println(scenicScore(forest))
	t.FailNow()
}

func TestScenicScore2(t *testing.T) {
	input := `351541053031130431201234
224125240541123033121143
223310430540142302200433
450101154534004044231012
443521502544433404042033
320014551114014133012203
250234251433044340323003`
	forest := parse(input)
	print(forest)
	fmt.Println()
	fmt.Println(scenicScore(forest))
	t.FailNow()
}
