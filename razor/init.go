package razor

var (
	locals      interface{}
	isLocalsSet bool
)

func SetLocals(locs interface{}) {
	locals = locs
	isLocalsSet = true
}

func Locals() interface{} {
	if !isLocalsSet {
		panic("SetLocals was not called")
	}
	return locals
}
