package structPack

type expStruct struct {
	Mi1 int
	Mf1 float32
}

func NewExpStruct() *expStruct {
	return new(expStruct)
}

func New(a int, b float32) *expStruct {
	expStruct := new(expStruct)
	expStruct.Mi1 = a
	expStruct.Mf1 = b
	return expStruct
}
