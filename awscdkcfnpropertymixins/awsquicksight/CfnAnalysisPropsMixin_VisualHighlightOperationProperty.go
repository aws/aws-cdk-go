package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   visualHighlightOperationProperty := &VisualHighlightOperationProperty{
//   	Trigger: jsii.String("trigger"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualhighlightoperation.html
//
type CfnAnalysisPropsMixin_VisualHighlightOperationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualhighlightoperation.html#cfn-quicksight-analysis-visualhighlightoperation-trigger
	//
	Trigger *string `field:"optional" json:"trigger" yaml:"trigger"`
}

