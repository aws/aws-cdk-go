package awsquicksight


// The visual publish options of a visual in a dashboard.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardVisualPublishOptionsProperty := &DashboardVisualPublishOptionsProperty{
//   	ExportHiddenFieldsOption: &ExportHiddenFieldsOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   }
//
type CfnDashboard_DashboardVisualPublishOptionsProperty struct {
	// Determines if hidden fields are included in an exported dashboard.
	ExportHiddenFieldsOption interface{} `field:"optional" json:"exportHiddenFieldsOption" yaml:"exportHiddenFieldsOption"`
}

