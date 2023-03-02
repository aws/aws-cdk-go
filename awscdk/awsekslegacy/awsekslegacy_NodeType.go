package awsekslegacy


// Whether the worker nodes should support GPU or just standard instances.
// Experimental.
type NodeType string

const (
	// Standard instances.
	// Experimental.
	NodeType_STANDARD NodeType = "STANDARD"
	// GPU instances.
	// Experimental.
	NodeType_GPU NodeType = "GPU"
)

