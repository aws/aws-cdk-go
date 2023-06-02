package awscdksyntheticsalpha


// All known Lambda runtime families.
// Experimental.
type RuntimeFamily string

const (
	// All Lambda runtimes that depend on Node.js.
	// Experimental.
	RuntimeFamily_NODEJS RuntimeFamily = "NODEJS"
	// All lambda runtimes that depend on Python.
	// Experimental.
	RuntimeFamily_PYTHON RuntimeFamily = "PYTHON"
	// Any future runtime family.
	// Experimental.
	RuntimeFamily_OTHER RuntimeFamily = "OTHER"
)

