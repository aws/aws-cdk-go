package awsecs


// Propagate tags for Managed Instances.
type PropagateManagedInstancesTags string

const (
	// Propagate tags from the capacity provider.
	PropagateManagedInstancesTags_CAPACITY_PROVIDER PropagateManagedInstancesTags = "CAPACITY_PROVIDER"
	// Do not propagate tags.
	PropagateManagedInstancesTags_NONE PropagateManagedInstancesTags = "NONE"
)

