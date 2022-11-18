package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties of a reference to a CodeDeploy Lambda Deployment Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customLambdaDeploymentConfigProps := &customLambdaDeploymentConfigProps{
//   	interval: cdk.duration.minutes(jsii.Number(30)),
//   	percentage: jsii.Number(123),
//   	type: awscdk.Aws_codedeploy.customLambdaDeploymentConfigType_CANARY,
//
//   	// the properties below are optional
//   	deploymentConfigName: jsii.String("deploymentConfigName"),
//   }
//
// Deprecated: Use `LambdaDeploymentConfig`.
type CustomLambdaDeploymentConfigProps struct {
	// The interval, in number of minutes: - For LINEAR, how frequently additional traffic is shifted - For CANARY, how long to shift traffic before the full deployment.
	// Deprecated: Use `LambdaDeploymentConfig`.
	Interval awscdk.Duration `field:"required" json:"interval" yaml:"interval"`
	// The integer percentage of traffic to shift: - For LINEAR, the percentage to shift every interval - For CANARY, the percentage to shift until the interval passes, before the full deployment.
	// Deprecated: Use `LambdaDeploymentConfig`.
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
	// The type of deployment config, either CANARY or LINEAR.
	// Deprecated: Use `LambdaDeploymentConfig`.
	Type CustomLambdaDeploymentConfigType `field:"required" json:"type" yaml:"type"`
	// The verbatim name of the deployment config.
	//
	// Must be unique per account/region.
	// Other parameters cannot be updated if this name is provided.
	// Deprecated: Use `LambdaDeploymentConfig`.
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
}

