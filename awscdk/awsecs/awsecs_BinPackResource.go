package awsecs


// Instance resource used for bin packing.
// Experimental.
type BinPackResource string

const (
	// Fill up hosts' CPU allocations first.
	// Experimental.
	BinPackResource_CPU BinPackResource = "CPU"
	// Fill up hosts' memory allocations first.
	// Experimental.
	BinPackResource_MEMORY BinPackResource = "MEMORY"
)

