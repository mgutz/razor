package razor

// M is a map.
type M map[string]interface{}

// App is map containing app scoped variables.
var App map[string]interface{}

// SetAppState sets App map.
func SetAppState(state map[string]interface{}) {
	App = state
}

// Sections holds rendered content of sections for a template.
type Sections map[string]*SafeBuffer
