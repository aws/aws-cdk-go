package awssynthetics


// All known Lambda runtime families.
type RuntimeFamily string

const (
	// All Lambda runtimes that depend on Node.js.
	RuntimeFamily_NODEJS RuntimeFamily = "NODEJS"
	// All lambda runtimes that depend on Python.
	RuntimeFamily_PYTHON RuntimeFamily = "PYTHON"
	// Any future runtime family.
	RuntimeFamily_OTHER RuntimeFamily = "OTHER"
)

