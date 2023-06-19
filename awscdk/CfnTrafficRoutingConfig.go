package awscdk


// Traffic routing configuration settings.
//
// The type of the {@link CfnCodeDeployBlueGreenHookProps.trafficRoutingConfig} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficRoutingConfig := &CfnTrafficRoutingConfig{
//   	Type: monocdk.CfnTrafficRoutingType_ALL_AT_ONCE,
//
//   	// the properties below are optional
//   	TimeBasedCanary: &CfnTrafficRoutingTimeBasedCanary{
//   		BakeTimeMins: jsii.Number(123),
//   		StepPercentage: jsii.Number(123),
//   	},
//   	TimeBasedLinear: &CfnTrafficRoutingTimeBasedLinear{
//   		BakeTimeMins: jsii.Number(123),
//   		StepPercentage: jsii.Number(123),
//   	},
//   }
//
// Experimental.
type CfnTrafficRoutingConfig struct {
	// The type of traffic shifting used by the blue-green deployment configuration.
	// Experimental.
	Type CfnTrafficRoutingType `field:"required" json:"type" yaml:"type"`
	// The configuration for traffic routing when {@link type} is {@link CfnTrafficRoutingType.TIME_BASED_CANARY}.
	// Experimental.
	TimeBasedCanary *CfnTrafficRoutingTimeBasedCanary `field:"optional" json:"timeBasedCanary" yaml:"timeBasedCanary"`
	// The configuration for traffic routing when {@link type} is {@link CfnTrafficRoutingType.TIME_BASED_LINEAR}.
	// Experimental.
	TimeBasedLinear *CfnTrafficRoutingTimeBasedLinear `field:"optional" json:"timeBasedLinear" yaml:"timeBasedLinear"`
}

