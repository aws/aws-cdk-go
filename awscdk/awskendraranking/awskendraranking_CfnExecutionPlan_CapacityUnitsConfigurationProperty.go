package awskendraranking


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityUnitsConfigurationProperty := &capacityUnitsConfigurationProperty{
//   	rescoreCapacityUnits: jsii.Number(123),
//   }
//
type CfnExecutionPlan_CapacityUnitsConfigurationProperty struct {
	// `CfnExecutionPlan.CapacityUnitsConfigurationProperty.RescoreCapacityUnits`.
	RescoreCapacityUnits *float64 `field:"required" json:"rescoreCapacityUnits" yaml:"rescoreCapacityUnits"`
}

