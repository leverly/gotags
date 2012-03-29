package Test

type Struct struct {
	Field1, Field2 int
	field3         string
	field4         *bool
}

func NewStruct() *Struct {
	return &Struct{}
}

func (s Struct) F1() ([]bool, [2]*string) {
}

func (Struct) F2() (result bool) {
}

type TestEmbed struct {
	Struct
	*io.Writer
}

func NewTestEmbed() TestEmbed {
}