package awsiotfleetwise


// Vehicles associated with the state template will stream telemetry data during a specified time period.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   periodicStateTemplateUpdateStrategyProperty := &PeriodicStateTemplateUpdateStrategyProperty{
//   	StateTemplateUpdateRate: &TimePeriodProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-vehicle-periodicstatetemplateupdatestrategy.html
//
type CfnVehicle_PeriodicStateTemplateUpdateStrategyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-vehicle-periodicstatetemplateupdatestrategy.html#cfn-iotfleetwise-vehicle-periodicstatetemplateupdatestrategy-statetemplateupdaterate
	//
	StateTemplateUpdateRate interface{} `field:"required" json:"stateTemplateUpdateRate" yaml:"stateTemplateUpdateRate"`
}

