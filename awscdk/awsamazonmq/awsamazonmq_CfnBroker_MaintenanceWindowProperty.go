package awsamazonmq


// The parameters that determine the `WeeklyStartTime` to apply pending updates or patches to the broker.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceWindowProperty := &maintenanceWindowProperty{
//   	dayOfWeek: jsii.String("dayOfWeek"),
//   	timeOfDay: jsii.String("timeOfDay"),
//   	timeZone: jsii.String("timeZone"),
//   }
//
type CfnBroker_MaintenanceWindowProperty struct {
	// The day of the week.
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// The time, in 24-hour format.
	TimeOfDay *string `field:"required" json:"timeOfDay" yaml:"timeOfDay"`
	// The time zone, UTC by default, in either the Country/City format, or the UTC offset format.
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
}

