package awsiotfleetwise


// The update strategy for the state template.
//
// Vehicles associated with the state template can stream telemetry data with either an `onChange` or `periodic` update strategy.
//
// > Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html) in the *AWS IoT FleetWise Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var onChange interface{}
//
//   stateTemplateUpdateStrategyProperty := &StateTemplateUpdateStrategyProperty{
//   	OnChange: onChange,
//   	Periodic: &PeriodicStateTemplateUpdateStrategyProperty{
//   		StateTemplateUpdateRate: &TimePeriodProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-vehicle-statetemplateupdatestrategy.html
//
type CfnVehicle_StateTemplateUpdateStrategyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-vehicle-statetemplateupdatestrategy.html#cfn-iotfleetwise-vehicle-statetemplateupdatestrategy-onchange
	//
	OnChange interface{} `field:"optional" json:"onChange" yaml:"onChange"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-vehicle-statetemplateupdatestrategy.html#cfn-iotfleetwise-vehicle-statetemplateupdatestrategy-periodic
	//
	Periodic interface{} `field:"optional" json:"periodic" yaml:"periodic"`
}

