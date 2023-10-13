package awscdksyntheticsalpha


// All known Lambda runtime families.
// Deprecated.
type RuntimeFamily string

const (
	// All Lambda runtimes that depend on Node.js.
	// Deprecated.
	RuntimeFamily_NODEJS RuntimeFamily = "NODEJS"
	// All lambda runtimes that depend on Python.
	// Deprecated.
	RuntimeFamily_PYTHON RuntimeFamily = "PYTHON"
	// Any future runtime family.
	// Deprecated.
	RuntimeFamily_OTHER RuntimeFamily = "OTHER"
)

