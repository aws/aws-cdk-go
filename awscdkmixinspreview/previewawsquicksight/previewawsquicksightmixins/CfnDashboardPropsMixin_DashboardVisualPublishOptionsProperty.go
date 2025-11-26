package previewawsquicksightmixins


// The visual publish options of a visual in a dashboard.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dashboardVisualPublishOptionsProperty := &DashboardVisualPublishOptionsProperty{
//   	ExportHiddenFieldsOption: &ExportHiddenFieldsOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardvisualpublishoptions.html
//
type CfnDashboardPropsMixin_DashboardVisualPublishOptionsProperty struct {
	// Determines if hidden fields are included in an exported dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardvisualpublishoptions.html#cfn-quicksight-dashboard-dashboardvisualpublishoptions-exporthiddenfieldsoption
	//
	ExportHiddenFieldsOption interface{} `field:"optional" json:"exportHiddenFieldsOption" yaml:"exportHiddenFieldsOption"`
}

