package awscodedeploy


// Lambda Deployment config type.
//
// Example:
//   var application lambdaApplication
//   var alias alias
//   config := codedeploy.NewCustomLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &CustomLambdaDeploymentConfigProps{
//   	Type: codedeploy.CustomLambdaDeploymentConfigType_CANARY,
//   	Interval: awscdk.Duration_Minutes(jsii.Number(1)),
//   	Percentage: jsii.Number(5),
//   })
//   deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &LambdaDeploymentGroupProps{
//   	Application: Application,
//   	Alias: Alias,
//   	DeploymentConfig: config,
//   })
//
// Experimental.
type CustomLambdaDeploymentConfigType string

const (
	// Canary deployment type.
	// Experimental.
	CustomLambdaDeploymentConfigType_CANARY CustomLambdaDeploymentConfigType = "CANARY"
	// Linear deployment type.
	// Experimental.
	CustomLambdaDeploymentConfigType_LINEAR CustomLambdaDeploymentConfigType = "LINEAR"
)

