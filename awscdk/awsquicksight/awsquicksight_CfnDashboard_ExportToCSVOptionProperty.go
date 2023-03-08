package awsquicksight


// Export to .csv option.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   exportToCSVOptionProperty := &exportToCSVOptionProperty{
//   	availabilityStatus: jsii.String("availabilityStatus"),
//   }
//
type CfnDashboard_ExportToCSVOptionProperty struct {
	// Availability status.
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

