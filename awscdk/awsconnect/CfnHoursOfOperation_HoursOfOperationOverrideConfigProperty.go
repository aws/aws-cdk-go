package awsconnect


// Contains information about the hours of operation override.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hoursOfOperationOverrideConfigProperty := &HoursOfOperationOverrideConfigProperty{
//   	Day: jsii.String("day"),
//   	EndTime: &OverrideTimeSliceProperty{
//   		Hours: jsii.Number(123),
//   		Minutes: jsii.Number(123),
//   	},
//   	StartTime: &OverrideTimeSliceProperty{
//   		Hours: jsii.Number(123),
//   		Minutes: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverrideconfig.html
//
type CfnHoursOfOperation_HoursOfOperationOverrideConfigProperty struct {
	// The day that the hours of operation override applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverrideconfig.html#cfn-connect-hoursofoperation-hoursofoperationoverrideconfig-day
	//
	Day *string `field:"required" json:"day" yaml:"day"`
	// The start time or end time for an an hours of operation override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverrideconfig.html#cfn-connect-hoursofoperation-hoursofoperationoverrideconfig-endtime
	//
	EndTime interface{} `field:"required" json:"endTime" yaml:"endTime"`
	// The start time or end time for an an hours of operation override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverrideconfig.html#cfn-connect-hoursofoperation-hoursofoperationoverrideconfig-starttime
	//
	StartTime interface{} `field:"required" json:"startTime" yaml:"startTime"`
}

