
func FindCorner(x int, y int) []int {
	Corner := []int{x, y}
    
	return Corner
}

func CheckOverlap(start []int, end []int, rec2 []int) bool {
    for x,y := start[0], start[1] ; x < end[0] ; x++ {
		if x > rec2[0] && x < rec2[2] {
			if y > rec2[1] && y < rec2[3] {
				return true
			} else {
				return false
			}	
		}
	}
    return false
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {

	var ButtomLeft1 []int = FindCorner(rec1[0], rec1[1])
	var TopLeft1 []int = FindCorner(rec1[0], rec1[3])
	var ButtomRight1 []int = FindCorner(rec1[2], rec1[1])
	var TopRight1 []int = FindCorner(rec1[2], rec1[3])
    var ButtomLeft2 []int = FindCorner(rec2[0], rec2[1])
	var TopLeft2 []int = FindCorner(rec2[0], rec2[3])
	var ButtomRight2 []int = FindCorner(rec2[2], rec2[1])
	var TopRight2 []int = FindCorner(rec2[2], rec2[3])
	if CheckOverlap(ButtomLeft1, TopLeft1, rec2) {
		return true
	}
	if CheckOverlap(ButtomLeft1, ButtomRight1, rec2) {
		return true
	}
    if CheckOverlap(ButtomRight1, TopRight1, rec2) {
		return true
	}
	if CheckOverlap(TopLeft1, TopRight1,rec2) {
		return true
	}
   if CheckOverlap(ButtomLeft2, TopLeft2, rec1) {
		return true
	}
	if CheckOverlap(TopLeft2, TopRight2, rec1) {
		return true
	}
    if CheckOverlap(ButtomRight2, TopRight2, rec1) {
		return true
	}
	if CheckOverlap(ButtomLeft2, ButtomRight2,rec1) {
		return true
	}
    return false
}
testcase 
[-930,154,-278,985]
[-912, 362,-630, 910]

[4,0,6,6]
[-5,-3,4,2]

[206110757,199595273,851054072,983256382]
[-898677230,-47320317,-858614307,736177810]

[-191960168,717271550,-151449330,847014606]
[-911502337,-435873766,-627023978,179876431]

[-836959955,-724514313,368220652,-449465289]
[-703416829,-64916327,713807924,838108468]