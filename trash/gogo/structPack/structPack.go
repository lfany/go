package structPack

type expStruct struct {
	Mi1 int
	Mf1 float32
}

func NewExpStruct() *expStruct {
	return new(expStruct)
}
