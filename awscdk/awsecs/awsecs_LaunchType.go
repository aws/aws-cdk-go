package awsecs


// The launch type of an ECS service.
// Experimental.
type LaunchType string

const (
	// The service will be launched using the EC2 launch type.
	// Experimental.
	LaunchType_EC2 LaunchType = "EC2"
	// The service will be launched using the FARGATE launch type.
	// Experimental.
	LaunchType_FARGATE LaunchType = "FARGATE"
	// The service will be launched using the EXTERNAL launch type.
	// Experimental.
	LaunchType_EXTERNAL LaunchType = "EXTERNAL"
)

