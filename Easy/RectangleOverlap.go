func CalculateBase(rec []int) []int {
	var Base [2]int
	Base[0] = rec[2]
	Base[1] = rec[1]
	return Base
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	var base1 []int = CalculateBase(rec1)
	var base2 []int = CalculateBase(rec2)
	for i := 0; i < len(rec2); i++ {

	}

}