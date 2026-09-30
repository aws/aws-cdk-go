package awschime


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keywordMatchConfigurationProperty := &KeywordMatchConfigurationProperty{
//   	Keywords: []*string{
//   		jsii.String("keywords"),
//   	},
//   	RuleName: jsii.String("ruleName"),
//
//   	// the properties below are optional
//   	Negate: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-keywordmatchconfiguration.html
//
type CfnMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty struct {
	// The keywords or phrases to match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-keywordmatchconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-keywordmatchconfiguration-keywords
	//
	Keywords *[]*string `field:"required" json:"keywords" yaml:"keywords"`
	// The name of the keyword match rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-keywordmatchconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-keywordmatchconfiguration-rulename
	//
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Matches keywords on their presence or absence.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-keywordmatchconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-keywordmatchconfiguration-negate
	//
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
}

