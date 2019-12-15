package gonfig

const (
	//LoadError = "Error loading the configuration"
	LoadError = GonfigError("Error loading the configuration")
)

// GonfigError is an error in the module.
type GonfigError string

func (g GonfigError) Error() string {
	return string(g)
}
