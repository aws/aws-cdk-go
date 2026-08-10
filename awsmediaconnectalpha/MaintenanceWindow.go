package awsmediaconnectalpha


// Configuration for scheduled maintenance windows.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   maintenanceWindow := &MaintenanceWindow{
//   	Day: mediaconnect_alpha.MaintenanceDay_MONDAY,
//   	Time: jsii.String("time"),
//   }
//
// Experimental.
type MaintenanceWindow struct {
	// A day of a week when the maintenance will happen.
	// Experimental.
	Day MaintenanceDay `field:"required" json:"day" yaml:"day"`
	// The maintenance start time in UTC, 24-hour HH:MM format.
	//
	// Minutes must be 00 (e.g., '02:00', '13:00').
	// Experimental.
	Time *string `field:"required" json:"time" yaml:"time"`
}

