package awsmedialivealpha


// Maintenance window settings for the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var maintenanceDay MaintenanceDay
//
//   maintenanceSettings := &MaintenanceSettings{
//   	Day: maintenanceDay,
//
//   	// the properties below are optional
//   	Time: jsii.String("time"),
//   }
//
// Experimental.
type MaintenanceSettings struct {
	// The day of the week for maintenance.
	// Experimental.
	Day MaintenanceDay `field:"required" json:"day" yaml:"day"`
	// The start time for maintenance in UTC (HH:MM format, e.g. '02:00').
	// Default: '02:00'.
	//
	// Experimental.
	Time *string `field:"optional" json:"time" yaml:"time"`
}

