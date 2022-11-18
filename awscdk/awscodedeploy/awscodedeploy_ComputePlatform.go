package awscodedeploy


// The compute platform of a deployment configuration.
type ComputePlatform string

const (
	// The deployment will target EC2 instances or on-premise servers.
	ComputePlatform_SERVER ComputePlatform = "SERVER"
	// The deployment will target a Lambda function.
	ComputePlatform_LAMBDA ComputePlatform = "LAMBDA"
	// The deployment will target an ECS server.
	ComputePlatform_ECS ComputePlatform = "ECS"
)

