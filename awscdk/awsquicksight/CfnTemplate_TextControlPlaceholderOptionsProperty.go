package awsquicksight


// The configuration of the placeholder options in a text control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textControlPlaceholderOptionsProperty := &TextControlPlaceholderOptionsProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-textcontrolplaceholderoptions.html
//
type CfnTemplate_TextControlPlaceholderOptionsProperty struct {
	// The visibility configuration of the placeholder options in a text control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-textcontrolplaceholderoptions.html#cfn-quicksight-template-textcontrolplaceholderoptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

