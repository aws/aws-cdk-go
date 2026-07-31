package awsmediaconnectalpha


// Configuration for scheduled maintenance windows.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   maintenanceConfiguration := &MaintenanceConfiguration{
//   	Day: mediaconnect_alpha.MaintenanceDay_MONDAY,
//   	Time: jsii.String("time"),
//   }
//
// Experimental.
type MaintenanceConfiguration struct {
	// Day of the week for maintenance (e.g., 'Monday', 'Tuesday').
	// Experimental.
	Day MaintenanceDay `field:"required" json:"day" yaml:"day"`
	// Time of day for maintenance in HH:MM format (e.g., '02:00').
	// Experimental.
	Time *string `field:"required" json:"time" yaml:"time"`
}

