package gibson

var (
	gibson *GibsonStruct
)

type GibsonStruct struct {
}

func GetInstance() *GibsonStruct {
	if gibson == nil {
		gibson = &GibsonStruct{}
	}
	return gibson
}

func (*GibsonStruct) Run() {
	r := NewRouter()
	r.Run()
}
