package awscodedeploy


// The configuration that specifies how traffic is shifted from one version of a Lambda function to another version during an AWS Lambda deployment, or from one Amazon ECS task set to another during an Amazon ECS deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficRoutingConfigProperty := &trafficRoutingConfigProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	timeBasedCanary: &timeBasedCanaryProperty{
//   		canaryInterval: jsii.Number(123),
//   		canaryPercentage: jsii.Number(123),
//   	},
//   	timeBasedLinear: &timeBasedLinearProperty{
//   		linearInterval: jsii.Number(123),
//   		linearPercentage: jsii.Number(123),
//   	},
//   }
//
type CfnDeploymentConfig_TrafficRoutingConfigProperty struct {
	// The type of traffic shifting ( `TimeBasedCanary` or `TimeBasedLinear` ) used by a deployment configuration.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A configuration that shifts traffic from one version of a Lambda function or ECS task set to another in two increments.
	//
	// The original and target Lambda function versions or ECS task sets are specified in the deployment's AppSpec file.
	TimeBasedCanary interface{} `field:"optional" json:"timeBasedCanary" yaml:"timeBasedCanary"`
	// A configuration that shifts traffic from one version of a Lambda function or Amazon ECS task set to another in equal increments, with an equal number of minutes between each increment.
	//
	// The original and target Lambda function versions or Amazon ECS task sets are specified in the deployment's AppSpec file.
	TimeBasedLinear interface{} `field:"optional" json:"timeBasedLinear" yaml:"timeBasedLinear"`
}

