package awsscheduler


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderStrategyItemProperty := &capacityProviderStrategyItemProperty{
//   	capacityProvider: jsii.String("capacityProvider"),
//
//   	// the properties below are optional
//   	base: jsii.Number(123),
//   	weight: jsii.Number(123),
//   }
//
type CfnSchedule_CapacityProviderStrategyItemProperty struct {
	// `CfnSchedule.CapacityProviderStrategyItemProperty.CapacityProvider`.
	CapacityProvider *string `field:"required" json:"capacityProvider" yaml:"capacityProvider"`
	// `CfnSchedule.CapacityProviderStrategyItemProperty.Base`.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// `CfnSchedule.CapacityProviderStrategyItemProperty.Weight`.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

