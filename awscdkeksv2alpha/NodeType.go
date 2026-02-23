package awscdkeksv2alpha


// Whether the worker nodes should support GPU or just standard instances.
// Deprecated.
type NodeType string

const (
	// Standard instances.
	// Deprecated.
	NodeType_STANDARD NodeType = "STANDARD"
	// GPU instances.
	// Deprecated.
	NodeType_GPU NodeType = "GPU"
	// Inferentia instances.
	// Deprecated.
	NodeType_INFERENTIA NodeType = "INFERENTIA"
	// Trainium instances.
	// Deprecated.
	NodeType_TRAINIUM NodeType = "TRAINIUM"
)

