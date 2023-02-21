package awsquicksight


// An ad hoc (one-time) filtering option.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   adHocFilteringOptionProperty := &AdHocFilteringOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
type CfnDashboard_AdHocFilteringOptionProperty struct {
	// Availability status.
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

