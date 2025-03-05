package awsquicksight


// The general visual interactions setup for visual publish options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var contextMenuOption interface{}
//   var visualMenuOption interface{}
//
//   visualInteractionOptionsProperty := &VisualInteractionOptionsProperty{
//   	ContextMenuOption: contextMenuOption,
//   	VisualMenuOption: visualMenuOption,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualinteractionoptions.html
//
type CfnAnalysis_VisualInteractionOptionsProperty struct {
	// The context menu options for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualinteractionoptions.html#cfn-quicksight-analysis-visualinteractionoptions-contextmenuoption
	//
	ContextMenuOption interface{} `field:"optional" json:"contextMenuOption" yaml:"contextMenuOption"`
	// The on-visual menu options for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualinteractionoptions.html#cfn-quicksight-analysis-visualinteractionoptions-visualmenuoption
	//
	VisualMenuOption interface{} `field:"optional" json:"visualMenuOption" yaml:"visualMenuOption"`
}

