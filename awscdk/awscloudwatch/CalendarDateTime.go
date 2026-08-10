package awscloudwatch


// Represents a calendar date and time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   calendarDateTime := &CalendarDateTime{
//   	Day: jsii.Number(123),
//   	Hour: jsii.Number(123),
//   	Minute: jsii.Number(123),
//   	Month: jsii.Number(123),
//   	Year: jsii.Number(123),
//   }
//
type CalendarDateTime struct {
	// The day of the date.
	//
	// Valid range: 1-31.
	Day *float64 `field:"required" json:"day" yaml:"day"`
	// The hour of the time.
	//
	// Valid range: 0-23.
	Hour *float64 `field:"required" json:"hour" yaml:"hour"`
	// The minute of the time.
	//
	// Valid range: 0-59.
	Minute *float64 `field:"required" json:"minute" yaml:"minute"`
	// The month of the date.
	//
	// Valid range: 1-12.
	Month *float64 `field:"required" json:"month" yaml:"month"`
	// The year of the date.
	Year *float64 `field:"required" json:"year" yaml:"year"`
}

