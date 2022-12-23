package awseks


// Whether the worker nodes should support GPU or just standard instances.
type NodeType string

const (
	// Standard instances.
	NodeType_STANDARD NodeType = "STANDARD"
	// GPU instances.
	NodeType_GPU NodeType = "GPU"
	// Inferentia instances.
	NodeType_INFERENTIA NodeType = "INFERENTIA"
)

