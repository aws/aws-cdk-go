package awscodedeploy


// Represents the configuration specific to linear traffic shifting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linearTrafficRoutingConfig := &linearTrafficRoutingConfig{
//   	linearInterval: jsii.Number(123),
//   	linearPercentage: jsii.Number(123),
//   }
//
type LinearTrafficRoutingConfig struct {
	// The number of minutes between each incremental traffic shift of a `TimeBasedLinear` deployment.
	LinearInterval *float64 `field:"required" json:"linearInterval" yaml:"linearInterval"`
	// The percentage of traffic that is shifted at the start of each increment of a `TimeBasedLinear` deployment.
	LinearPercentage *float64 `field:"required" json:"linearPercentage" yaml:"linearPercentage"`
}

