package awsquicksight


// Formatting configuration for icon set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionalFormattingIconSetProperty := &ConditionalFormattingIconSetProperty{
//   	Expression: jsii.String("expression"),
//
//   	// the properties below are optional
//   	IconSetType: jsii.String("iconSetType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingiconset.html
//
type CfnAnalysis_ConditionalFormattingIconSetProperty struct {
	// The expression that determines the formatting configuration for the icon set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingiconset.html#cfn-quicksight-analysis-conditionalformattingiconset-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Determines the icon set type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-conditionalformattingiconset.html#cfn-quicksight-analysis-conditionalformattingiconset-iconsettype
	//
	IconSetType *string `field:"optional" json:"iconSetType" yaml:"iconSetType"`
}

