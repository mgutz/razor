package razor

var Locals map[string]interface{}

func SetLocals(locals map[string]interface{}) {
	Locals = locals
}
