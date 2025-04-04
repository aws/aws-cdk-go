package awscdk


// Traffic routing configuration settings.
//
// The type of the `CfnCodeDeployBlueGreenHookProps.trafficRoutingConfig` property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficRoutingConfig := &CfnTrafficRoutingConfig{
//   	Type: cdk.CfnTrafficRoutingType_ALL_AT_ONCE,
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
type CfnTrafficRoutingConfig struct {
	// The type of traffic shifting used by the blue-green deployment configuration.
	Type CfnTrafficRoutingType `field:"required" json:"type" yaml:"type"`
	// The configuration for traffic routing when `type` is `CfnTrafficRoutingType.TIME_BASED_CANARY`.
	// Default: - none.
	//
	TimeBasedCanary *CfnTrafficRoutingTimeBasedCanary `field:"optional" json:"timeBasedCanary" yaml:"timeBasedCanary"`
	// The configuration for traffic routing when `type` is `CfnTrafficRoutingType.TIME_BASED_LINEAR`.
	// Default: - none.
	//
	TimeBasedLinear *CfnTrafficRoutingTimeBasedLinear `field:"optional" json:"timeBasedLinear" yaml:"timeBasedLinear"`
}

