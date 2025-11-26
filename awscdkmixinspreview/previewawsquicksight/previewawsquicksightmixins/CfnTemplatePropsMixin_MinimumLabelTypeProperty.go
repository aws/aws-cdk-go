package previewawsquicksightmixins


// The minimum label of a data path label.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   minimumLabelTypeProperty := &MinimumLabelTypeProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-minimumlabeltype.html
//
type CfnTemplatePropsMixin_MinimumLabelTypeProperty struct {
	// The visibility of the minimum label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-minimumlabeltype.html#cfn-quicksight-template-minimumlabeltype-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

