package matrix

type matrix struct {
	name string
}

func NewMatrix(name string) *matrix {
	return &matrix{"Tacks"}
}
