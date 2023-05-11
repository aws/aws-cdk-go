package awsquicksight


// Determines whether or not hidden fields are visible on exported dashbaords.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   exportWithHiddenFieldsOptionProperty := &ExportWithHiddenFieldsOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
type CfnDashboard_ExportWithHiddenFieldsOptionProperty struct {
	// The status of the export with hidden fields options.
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

