package awsiotfleetwise


// The length of time between state template updates.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timePeriodProperty := &TimePeriodProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-vehicle-timeperiod.html
//
type CfnVehicle_TimePeriodProperty struct {
	// A unit of time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-vehicle-timeperiod.html#cfn-iotfleetwise-vehicle-timeperiod-unit
	//
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// A number of time units.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-vehicle-timeperiod.html#cfn-iotfleetwise-vehicle-timeperiod-value
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

