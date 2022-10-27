package awscodedeploy


// Lambda Deployment config type.
// Deprecated: Use `LambdaDeploymentConfig`.
type CustomLambdaDeploymentConfigType string

const (
	// Canary deployment type.
	// Deprecated: Use `LambdaDeploymentConfig`.
	CustomLambdaDeploymentConfigType_CANARY CustomLambdaDeploymentConfigType = "CANARY"
	// Linear deployment type.
	// Deprecated: Use `LambdaDeploymentConfig`.
	CustomLambdaDeploymentConfigType_LINEAR CustomLambdaDeploymentConfigType = "LINEAR"
)

