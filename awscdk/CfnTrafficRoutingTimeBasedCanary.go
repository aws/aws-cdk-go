package awscdk


// The traffic routing configuration if `CfnTrafficRoutingConfig.type` is `CfnTrafficRoutingType.TIME_BASED_CANARY`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficRoutingTimeBasedCanary := &CfnTrafficRoutingTimeBasedCanary{
//   	BakeTimeMins: jsii.Number(123),
//   	StepPercentage: jsii.Number(123),
//   }
//
type CfnTrafficRoutingTimeBasedCanary struct {
	// The number of minutes between the first and second traffic shifts of a time-based canary deployment.
	// Default: 5.
	//
	BakeTimeMins *float64 `field:"optional" json:"bakeTimeMins" yaml:"bakeTimeMins"`
	// The percentage of traffic to shift in the first increment of a time-based canary deployment.
	//
	// The step percentage must be 14% or greater.
	// Default: 15.
	//
	StepPercentage *float64 `field:"optional" json:"stepPercentage" yaml:"stepPercentage"`
}

