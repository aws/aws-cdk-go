package awsconnect


// Contains information about the hours of operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hoursOfOperationConfigProperty := &hoursOfOperationConfigProperty{
//   	day: jsii.String("day"),
//   	endTime: &hoursOfOperationTimeSliceProperty{
//   		hours: jsii.Number(123),
//   		minutes: jsii.Number(123),
//   	},
//   	startTime: &hoursOfOperationTimeSliceProperty{
//   		hours: jsii.Number(123),
//   		minutes: jsii.Number(123),
//   	},
//   }
//
type CfnHoursOfOperation_HoursOfOperationConfigProperty struct {
	// The day that the hours of operation applies to.
	Day *string `field:"required" json:"day" yaml:"day"`
	// The end time that your contact center closes.
	EndTime interface{} `field:"required" json:"endTime" yaml:"endTime"`
	// The start time that your contact center opens.
	StartTime interface{} `field:"required" json:"startTime" yaml:"startTime"`
}

