package awsecs


// Instance resource used for bin packing.
type BinPackResource string

const (
	// Fill up hosts' CPU allocations first.
	BinPackResource_CPU BinPackResource = "CPU"
	// Fill up hosts' memory allocations first.
	BinPackResource_MEMORY BinPackResource = "MEMORY"
)

