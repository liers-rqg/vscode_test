package demo6

type Hero struct {
	Name  string
	Score float64
}

type HeroSlice []Hero

//实现了sort包下的一个接口以实现结构体排序
func (hs *HeroSlice) Len() int {
	return len(*hs)
}
func (hs *HeroSlice) Less(i, j int) bool {
	return (*hs)[i].Score > (*hs)[j].Score
}
func (hs *HeroSlice) Swap(i, j int) {
	(*hs)[i], (*hs)[j] = (*hs)[j], (*hs)[i]
}
