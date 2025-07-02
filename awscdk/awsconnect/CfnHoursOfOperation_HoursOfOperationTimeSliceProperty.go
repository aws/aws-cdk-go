package awsconnect


// The start time or end time for an hours of operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hoursOfOperationTimeSliceProperty := &HoursOfOperationTimeSliceProperty{
//   	Hours: jsii.Number(123),
//   	Minutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationtimeslice.html
//
type CfnHoursOfOperation_HoursOfOperationTimeSliceProperty struct {
	// The hours.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationtimeslice.html#cfn-connect-hoursofoperation-hoursofoperationtimeslice-hours
	//
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
	// The minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationtimeslice.html#cfn-connect-hoursofoperation-hoursofoperationtimeslice-minutes
	//
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

