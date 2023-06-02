package awsquicksight


// The axis sort options for a visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visualAxisSortOptionProperty := &VisualAxisSortOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
type CfnDashboard_VisualAxisSortOptionProperty struct {
	// The availaiblity status of a visual's axis sort options.
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

