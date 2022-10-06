package awscodedeploy


// Represents the configuration specific to canary traffic shifting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   canaryTrafficRoutingConfig := &canaryTrafficRoutingConfig{
//   	canaryInterval: jsii.Number(123),
//   	canaryPercentage: jsii.Number(123),
//   }
//
type CanaryTrafficRoutingConfig struct {
	// The number of minutes between the first and second traffic shifts of a `TimeBasedCanary` deployment.
	CanaryInterval *float64 `field:"required" json:"canaryInterval" yaml:"canaryInterval"`
	// The percentage of traffic to shift in the first increment of a `TimeBasedCanary` deployment.
	CanaryPercentage *float64 `field:"required" json:"canaryPercentage" yaml:"canaryPercentage"`
}

