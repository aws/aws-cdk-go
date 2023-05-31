package awsquicksight


// Determines if hidden fields are included in an exported dashboard.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   exportHiddenFieldsOptionProperty := &ExportHiddenFieldsOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
type CfnDashboard_ExportHiddenFieldsOptionProperty struct {
	// The status of the export hidden fields options of a dashbaord.
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

