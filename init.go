package razor

var App map[string]interface{}

func SetAppState(state map[string]interface{}) {
	App = state
}
