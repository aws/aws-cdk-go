// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Traffic routing configuration settings.
//
// The type of the {@link CfnCodeDeployBlueGreenHookProps.trafficRoutingConfig} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficRoutingConfig := &cfnTrafficRoutingConfig{
//   	type: cdk.cfnTrafficRoutingType_ALL_AT_ONCE,
//
//   	// the properties below are optional
//   	timeBasedCanary: &cfnTrafficRoutingTimeBasedCanary{
//   		bakeTimeMins: jsii.Number(123),
//   		stepPercentage: jsii.Number(123),
//   	},
//   	timeBasedLinear: &cfnTrafficRoutingTimeBasedLinear{
//   		bakeTimeMins: jsii.Number(123),
//   		stepPercentage: jsii.Number(123),
//   	},
//   }
//
type CfnTrafficRoutingConfig struct {
	// The type of traffic shifting used by the blue-green deployment configuration.
	Type CfnTrafficRoutingType `field:"required" json:"type" yaml:"type"`
	// The configuration for traffic routing when {@link type} is {@link CfnTrafficRoutingType.TIME_BASED_CANARY}.
	TimeBasedCanary *CfnTrafficRoutingTimeBasedCanary `field:"optional" json:"timeBasedCanary" yaml:"timeBasedCanary"`
	// The configuration for traffic routing when {@link type} is {@link CfnTrafficRoutingType.TIME_BASED_LINEAR}.
	TimeBasedLinear *CfnTrafficRoutingTimeBasedLinear `field:"optional" json:"timeBasedLinear" yaml:"timeBasedLinear"`
}

