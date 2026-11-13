package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visualCustomActionDefaultsProperty := &VisualCustomActionDefaultsProperty{
//   	HighlightOperation: &VisualHighlightOperationProperty{
//   		Trigger: jsii.String("trigger"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualcustomactiondefaults.html
//
type CfnAnalysis_VisualCustomActionDefaultsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualcustomactiondefaults.html#cfn-quicksight-analysis-visualcustomactiondefaults-highlightoperation
	//
	HighlightOperation interface{} `field:"optional" json:"highlightOperation" yaml:"highlightOperation"`
}

