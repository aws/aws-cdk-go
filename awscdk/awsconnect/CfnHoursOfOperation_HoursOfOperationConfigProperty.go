package awsconnect


// Contains information about the hours of operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hoursOfOperationConfigProperty := &HoursOfOperationConfigProperty{
//   	Day: jsii.String("day"),
//   	EndTime: &HoursOfOperationTimeSliceProperty{
//   		Hours: jsii.Number(123),
//   		Minutes: jsii.Number(123),
//   	},
//   	StartTime: &HoursOfOperationTimeSliceProperty{
//   		Hours: jsii.Number(123),
//   		Minutes: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationconfig.html
//
type CfnHoursOfOperation_HoursOfOperationConfigProperty struct {
	// The day that the hours of operation applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationconfig.html#cfn-connect-hoursofoperation-hoursofoperationconfig-day
	//
	Day *string `field:"required" json:"day" yaml:"day"`
	// The end time that your contact center closes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationconfig.html#cfn-connect-hoursofoperation-hoursofoperationconfig-endtime
	//
	EndTime interface{} `field:"required" json:"endTime" yaml:"endTime"`
	// The start time that your contact center opens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationconfig.html#cfn-connect-hoursofoperation-hoursofoperationconfig-starttime
	//
	StartTime interface{} `field:"required" json:"startTime" yaml:"startTime"`
}

