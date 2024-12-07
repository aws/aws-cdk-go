package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetImageTooltipConfigurationProperty := &SheetImageTooltipConfigurationProperty{
//   	TooltipText: &SheetImageTooltipTextProperty{
//   		PlainText: jsii.String("plainText"),
//   	},
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimagetooltipconfiguration.html
//
type CfnTemplate_SheetImageTooltipConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimagetooltipconfiguration.html#cfn-quicksight-template-sheetimagetooltipconfiguration-tooltiptext
	//
	TooltipText interface{} `field:"optional" json:"tooltipText" yaml:"tooltipText"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimagetooltipconfiguration.html#cfn-quicksight-template-sheetimagetooltipconfiguration-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

