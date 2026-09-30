package awschime


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   keywordMatchConfigurationProperty := &KeywordMatchConfigurationProperty{
//   	Keywords: []*string{
//   		jsii.String("keywords"),
//   	},
//   	Negate: jsii.Boolean(false),
//   	RuleName: jsii.String("ruleName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-keywordmatchconfiguration.html
//
type CfnMediaInsightsPipelineConfigurationPropsMixin_KeywordMatchConfigurationProperty struct {
	// The keywords or phrases to match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-keywordmatchconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-keywordmatchconfiguration-keywords
	//
	Keywords *[]*string `field:"optional" json:"keywords" yaml:"keywords"`
	// Matches keywords on their presence or absence.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-keywordmatchconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-keywordmatchconfiguration-negate
	//
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// The name of the keyword match rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-keywordmatchconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-keywordmatchconfiguration-rulename
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
}

