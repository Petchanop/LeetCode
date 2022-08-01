package leetcode

func maximumTime(time string) string {
	MaxTime := []byte(time)
	for i, digit := range MaxTime {
		if digit == 63 {
			if i == 0 && MaxTime[i] == 63 {
				if MaxTime[1] < 52 && MaxTime[1] >= 48 {
					MaxTime[i] = 50
				}
				if 52 <= MaxTime[1] && MaxTime[1] <= 57 {
					MaxTime[i] = 49
				} else {
					MaxTime[i] = 50
				}
			}
			if i == 1 && MaxTime[0] == 48 {
				MaxTime[1] = 57
			}
			if i == 1 && MaxTime[0] == 49 {
				MaxTime[1] = 57
			}
			if i == 1 && MaxTime[0] == 50 {
				MaxTime[1] = 51
			}
			if i == 3 && MaxTime[i] == 63 {
				MaxTime[3] = 53
			}
			if i == 4 && MaxTime[i] == 63 {
				MaxTime[4] = 57
			}
		}
	}
	NewTime := string(MaxTime)
	return NewTime
}
