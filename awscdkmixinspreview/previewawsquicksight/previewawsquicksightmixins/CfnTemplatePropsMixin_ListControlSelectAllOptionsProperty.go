package previewawsquicksightmixins


// The configuration of the `Select all` options in a list control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   listControlSelectAllOptionsProperty := &ListControlSelectAllOptionsProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-listcontrolselectalloptions.html
//
type CfnTemplatePropsMixin_ListControlSelectAllOptionsProperty struct {
	// The visibility configuration of the `Select all` options in a list control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-listcontrolselectalloptions.html#cfn-quicksight-template-listcontrolselectalloptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

