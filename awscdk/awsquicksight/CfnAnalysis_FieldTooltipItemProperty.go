package awsquicksight


// The tooltip item for the fields.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldTooltipItemProperty := &FieldTooltipItemProperty{
//   	FieldId: jsii.String("fieldId"),
//
//   	// the properties below are optional
//   	Label: jsii.String("label"),
//   	TooltipTarget: jsii.String("tooltipTarget"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-fieldtooltipitem.html
//
type CfnAnalysis_FieldTooltipItemProperty struct {
	// The unique ID of the field that is targeted by the tooltip.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-fieldtooltipitem.html#cfn-quicksight-analysis-fieldtooltipitem-fieldid
	//
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// The label of the tooltip item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-fieldtooltipitem.html#cfn-quicksight-analysis-fieldtooltipitem-label
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-fieldtooltipitem.html#cfn-quicksight-analysis-fieldtooltipitem-tooltiptarget
	//
	TooltipTarget *string `field:"optional" json:"tooltipTarget" yaml:"tooltipTarget"`
	// The visibility of the tooltip item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-fieldtooltipitem.html#cfn-quicksight-analysis-fieldtooltipitem-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

