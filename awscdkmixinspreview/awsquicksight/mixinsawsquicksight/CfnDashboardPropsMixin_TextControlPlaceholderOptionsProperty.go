package mixinsawsquicksight


// The configuration of the placeholder options in a text control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   textControlPlaceholderOptionsProperty := &TextControlPlaceholderOptionsProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-textcontrolplaceholderoptions.html
//
type CfnDashboardPropsMixin_TextControlPlaceholderOptionsProperty struct {
	// The visibility configuration of the placeholder options in a text control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-textcontrolplaceholderoptions.html#cfn-quicksight-dashboard-textcontrolplaceholderoptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

