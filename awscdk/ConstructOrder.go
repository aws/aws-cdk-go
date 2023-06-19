package awscdk


// In what order to return constructs.
// Experimental.
type ConstructOrder string

const (
	// Depth-first, pre-order.
	// Experimental.
	ConstructOrder_PREORDER ConstructOrder = "PREORDER"
	// Depth-first, post-order (leaf nodes first).
	// Experimental.
	ConstructOrder_POSTORDER ConstructOrder = "POSTORDER"
)

