package previewawsquicksightmixins


// The URL operation that opens a link to another webpage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customActionURLOperationProperty := &CustomActionURLOperationProperty{
//   	UrlTarget: jsii.String("urlTarget"),
//   	UrlTemplate: jsii.String("urlTemplate"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customactionurloperation.html
//
type CfnTemplatePropsMixin_CustomActionURLOperationProperty struct {
	// The target of the `CustomActionURLOperation` .
	//
	// Valid values are defined as follows:
	//
	// - `NEW_TAB` : Opens the target URL in a new browser tab.
	// - `NEW_WINDOW` : Opens the target URL in a new browser window.
	// - `SAME_TAB` : Opens the target URL in the same browser tab.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customactionurloperation.html#cfn-quicksight-template-customactionurloperation-urltarget
	//
	UrlTarget *string `field:"optional" json:"urlTarget" yaml:"urlTarget"`
	// THe URL link of the `CustomActionURLOperation` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customactionurloperation.html#cfn-quicksight-template-customactionurloperation-urltemplate
	//
	UrlTemplate *string `field:"optional" json:"urlTemplate" yaml:"urlTemplate"`
}

