package functions

func (y *Info) FindTheBestGroup() {
	var res []int

	SmallGroup := 0
	LongGroup := 0

	for i, v := range y.AllGroups {
		b := hi(v)

		if uq(res, b) {
			if SmallGroup == 0 {
				SmallGroup = b
			}

			if SmallGroup > b {
				SmallGroup = b
				y.SG = i
			}

			if LongGroup < b {
				LongGroup = b
				y.GG = i
			}
			y.UniqueGroups = append(y.UniqueGroups, v)
			y.NumberOfGroups++
		}
		res = append(res, b)
	}

	for i, v := range y.UniqueGroups {
		status := y.IsGoodGroup(v)
		if status {
			y.Print(v)
			break
		}
		if i == len(y.UniqueGroups)-1 && !status {
			y.Print(y.AllGroups[y.SG])
			break
		}

	}
}

func (y *Info) IsGoodGroup(matrix [][]string) bool {
	return y.NumberOfAnts <= hi(matrix)
}

func uq(res []int, b int) bool {
	if res == nil {
		return true
	}
	for _, it := range res {
		if it == b {
			return false
		}
	}
	return true
}

func hi(g [][]string) int {
	c := 0
	for _, v := range g {
		for _, h := range v {
			if h != "" {
				c++
			}
		}
	}
	return c
}
