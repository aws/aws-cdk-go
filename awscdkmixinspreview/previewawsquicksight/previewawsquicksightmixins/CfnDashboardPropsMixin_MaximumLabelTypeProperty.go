package previewawsquicksightmixins


// The maximum label of a data path label.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   maximumLabelTypeProperty := &MaximumLabelTypeProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-maximumlabeltype.html
//
type CfnDashboardPropsMixin_MaximumLabelTypeProperty struct {
	// The visibility of the maximum label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-maximumlabeltype.html#cfn-quicksight-dashboard-maximumlabeltype-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

