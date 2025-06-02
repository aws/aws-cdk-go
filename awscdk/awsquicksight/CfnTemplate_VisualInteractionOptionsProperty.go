package awsquicksight


// The general visual interactions setup for visual publish options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visualInteractionOptionsProperty := &VisualInteractionOptionsProperty{
//   	ContextMenuOption: &ContextMenuOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	VisualMenuOption: &VisualMenuOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualinteractionoptions.html
//
type CfnTemplate_VisualInteractionOptionsProperty struct {
	// The context menu options for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualinteractionoptions.html#cfn-quicksight-template-visualinteractionoptions-contextmenuoption
	//
	ContextMenuOption interface{} `field:"optional" json:"contextMenuOption" yaml:"contextMenuOption"`
	// The on-visual menu options for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualinteractionoptions.html#cfn-quicksight-template-visualinteractionoptions-visualmenuoption
	//
	VisualMenuOption interface{} `field:"optional" json:"visualMenuOption" yaml:"visualMenuOption"`
}

