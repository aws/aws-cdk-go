package previewawsquicksightmixins


// The visual display options for a data zoom scroll bar.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scrollBarOptionsProperty := &ScrollBarOptionsProperty{
//   	Visibility: jsii.String("visibility"),
//   	VisibleRange: &VisibleRangeOptionsProperty{
//   		PercentRange: &PercentVisibleRangeProperty{
//   			From: jsii.Number(123),
//   			To: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scrollbaroptions.html
//
type CfnDashboardPropsMixin_ScrollBarOptionsProperty struct {
	// The visibility of the data zoom scroll bar.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scrollbaroptions.html#cfn-quicksight-dashboard-scrollbaroptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// The visibility range for the data zoom scroll bar.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scrollbaroptions.html#cfn-quicksight-dashboard-scrollbaroptions-visiblerange
	//
	VisibleRange interface{} `field:"optional" json:"visibleRange" yaml:"visibleRange"`
}

