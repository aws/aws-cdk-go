package awscodedeploy


// Represents the structure to pass into the underlying CfnDeploymentConfig class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficRoutingConfig := &trafficRoutingConfig{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	timeBasedCanary: &canaryTrafficRoutingConfig{
//   		canaryInterval: jsii.Number(123),
//   		canaryPercentage: jsii.Number(123),
//   	},
//   	timeBasedLinear: &linearTrafficRoutingConfig{
//   		linearInterval: jsii.Number(123),
//   		linearPercentage: jsii.Number(123),
//   	},
//   }
//
type TrafficRoutingConfig struct {
	// The type of traffic shifting ( `TimeBasedCanary` or `TimeBasedLinear` ) used by a deployment configuration.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A configuration that shifts traffic from one version of a Lambda function or ECS task set to another in two increments.
	TimeBasedCanary *CanaryTrafficRoutingConfig `field:"optional" json:"timeBasedCanary" yaml:"timeBasedCanary"`
	// A configuration that shifts traffic from one version of a Lambda function or Amazon ECS task set to another in equal increments, with an equal number of minutes between each increment.
	TimeBasedLinear *LinearTrafficRoutingConfig `field:"optional" json:"timeBasedLinear" yaml:"timeBasedLinear"`
}

