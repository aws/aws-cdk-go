package awsfsx


// Properties required for setting up a weekly maintenance time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lustreMaintenanceTimeProps := &LustreMaintenanceTimeProps{
//   	Day: awscdk.Aws_fsx.Weekday_MONDAY,
//   	Hour: jsii.Number(123),
//   	Minute: jsii.Number(123),
//   }
//
// Experimental.
type LustreMaintenanceTimeProps struct {
	// The day of the week for maintenance to be performed.
	// Experimental.
	Day Weekday `field:"required" json:"day" yaml:"day"`
	// The hour of the day (from 0-24) for maintenance to be performed.
	// Experimental.
	Hour *float64 `field:"required" json:"hour" yaml:"hour"`
	// The minute of the hour (from 0-59) for maintenance to be performed.
	// Experimental.
	Minute *float64 `field:"required" json:"minute" yaml:"minute"`
}

