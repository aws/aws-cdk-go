package awsimagebuilderalpha


// The service in which a container should be registered.
// Experimental.
type RepositoryService string

const (
	// Indicates the container should be registered in ECR.
	// Experimental.
	RepositoryService_ECR RepositoryService = "ECR"
)

