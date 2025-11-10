package awsamazonmq


// The parameters that determine the WeeklyStartTime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceWindowProperty := &MaintenanceWindowProperty{
//   	DayOfWeek: jsii.String("dayOfWeek"),
//   	TimeOfDay: jsii.String("timeOfDay"),
//   	TimeZone: jsii.String("timeZone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-maintenancewindow.html
//
type CfnBroker_MaintenanceWindowProperty struct {
	// Required.
	//
	// The day of the week.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-maintenancewindow.html#cfn-amazonmq-broker-maintenancewindow-dayofweek
	//
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Required.
	//
	// The time, in 24-hour format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-maintenancewindow.html#cfn-amazonmq-broker-maintenancewindow-timeofday
	//
	TimeOfDay *string `field:"required" json:"timeOfDay" yaml:"timeOfDay"`
	// The time zone, UTC by default, in either the Country/City format, or the UTC offset format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-maintenancewindow.html#cfn-amazonmq-broker-maintenancewindow-timezone
	//
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
}

