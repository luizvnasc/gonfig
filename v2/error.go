package gonfig

const (
	// LoadError = "Error loading the configuration"
	LoadError = GonfigError("Error loading the configuration")
	// UnsupportedFileError = "Unsupported file extension"
	UnsupportedFileError = GonfigError("Unsupported file extension")
)

// GonfigError is an error in the module.
type GonfigError string

func (g GonfigError) Error() string {
	return string(g)
}
