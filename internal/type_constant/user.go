package type_constant

type GenderType string

func (e GenderType) String() string {
	return string(e)
}

const (
	GenderType_Gender_Male   GenderType = "Male"
	GenderType_Gender_Female GenderType = "Fenale"
)
