package awscodedeploy


// A configuration that shifts traffic from one version of a Lambda function or ECS task set to another in equal increments, with an equal number of minutes between each increment.
//
// The original and target Lambda function versions or ECS task sets are specified in the deployment's AppSpec file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeBasedLinearProperty := &timeBasedLinearProperty{
//   	linearInterval: jsii.Number(123),
//   	linearPercentage: jsii.Number(123),
//   }
//
type CfnDeploymentConfig_TimeBasedLinearProperty struct {
	// The number of minutes between each incremental traffic shift of a `TimeBasedLinear` deployment.
	LinearInterval *float64 `field:"required" json:"linearInterval" yaml:"linearInterval"`
	// The percentage of traffic that is shifted at the start of each increment of a `TimeBasedLinear` deployment.
	LinearPercentage *float64 `field:"required" json:"linearPercentage" yaml:"linearPercentage"`
}

