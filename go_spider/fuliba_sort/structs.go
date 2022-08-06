package fuliba_sort

type FulibaSub struct {
	Title string
	Like  int
	Reply int
	Read  int
	Url   string
	Time  string
}

type FulibaSubs []*FulibaSub

func (f FulibaSubs) Len() int {
	return len(f)
}

func (f FulibaSubs) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f FulibaSubs) Less(i, j int) bool {
	return f[i].Read > f[j].Read
	//return f[i].Like > f[j].Like
}
