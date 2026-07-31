package awsmediaconnectalpha


// Maintenance Window configuration for MediaConnect Flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   maintenanceWindow := &MaintenanceWindow{
//   	MaintenanceDay: mediaconnect_alpha.MaintenanceDay_MONDAY,
//   	MaintenanceStartHour: jsii.String("maintenanceStartHour"),
//   }
//
// Experimental.
type MaintenanceWindow struct {
	// A day of a week when the maintenance will happen.
	// Experimental.
	MaintenanceDay MaintenanceDay `field:"required" json:"maintenanceDay" yaml:"maintenanceDay"`
	// UTC time when the maintenance will happen.
	//
	// Use 24-hour HH:MM format. Minutes must be 00. Example: 13:00.
	// Experimental.
	MaintenanceStartHour *string `field:"required" json:"maintenanceStartHour" yaml:"maintenanceStartHour"`
}

