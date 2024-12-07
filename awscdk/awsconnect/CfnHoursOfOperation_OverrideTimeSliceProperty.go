package awsconnect


// The start time or end time for an an hours of operation override.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   overrideTimeSliceProperty := &OverrideTimeSliceProperty{
//   	Hours: jsii.Number(123),
//   	Minutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-overridetimeslice.html
//
type CfnHoursOfOperation_OverrideTimeSliceProperty struct {
	// The hours.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-overridetimeslice.html#cfn-connect-hoursofoperation-overridetimeslice-hours
	//
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
	// The minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-overridetimeslice.html#cfn-connect-hoursofoperation-overridetimeslice-minutes
	//
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

