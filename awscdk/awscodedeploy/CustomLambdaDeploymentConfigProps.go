package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties of a reference to a CodeDeploy Lambda Deployment Configuration.
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
type CustomLambdaDeploymentConfigProps struct {
	// The interval, in number of minutes: - For LINEAR, how frequently additional traffic is shifted - For CANARY, how long to shift traffic before the full deployment.
	// Experimental.
	Interval awscdk.Duration `field:"required" json:"interval" yaml:"interval"`
	// The integer percentage of traffic to shift: - For LINEAR, the percentage to shift every interval - For CANARY, the percentage to shift until the interval passes, before the full deployment.
	// Experimental.
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
	// The type of deployment config, either CANARY or LINEAR.
	// Experimental.
	Type CustomLambdaDeploymentConfigType `field:"required" json:"type" yaml:"type"`
	// The verbatim name of the deployment config.
	//
	// Must be unique per account/region.
	// Other parameters cannot be updated if this name is provided.
	// Experimental.
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
}

