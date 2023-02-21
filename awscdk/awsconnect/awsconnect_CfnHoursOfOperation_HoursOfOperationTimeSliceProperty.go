package awsconnect


// The start time or end time for an hours of operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hoursOfOperationTimeSliceProperty := &hoursOfOperationTimeSliceProperty{
//   	hours: jsii.Number(123),
//   	minutes: jsii.Number(123),
//   }
//
type CfnHoursOfOperation_HoursOfOperationTimeSliceProperty struct {
	// The hours.
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
	// The minutes.
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

