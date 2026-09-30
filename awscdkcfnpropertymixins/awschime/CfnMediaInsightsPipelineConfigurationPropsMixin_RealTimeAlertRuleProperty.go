package awschime


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   realTimeAlertRuleProperty := &RealTimeAlertRuleProperty{
//   	IssueDetectionConfiguration: &IssueDetectionConfigurationProperty{
//   		RuleName: jsii.String("ruleName"),
//   	},
//   	KeywordMatchConfiguration: &KeywordMatchConfigurationProperty{
//   		Keywords: []*string{
//   			jsii.String("keywords"),
//   		},
//   		Negate: jsii.Boolean(false),
//   		RuleName: jsii.String("ruleName"),
//   	},
//   	SentimentConfiguration: &SentimentConfigurationProperty{
//   		RuleName: jsii.String("ruleName"),
//   		SentimentType: jsii.String("sentimentType"),
//   		TimePeriod: jsii.Number(123),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-realtimealertrule.html
//
type CfnMediaInsightsPipelineConfigurationPropsMixin_RealTimeAlertRuleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-realtimealertrule.html#cfn-chime-mediainsightspipelineconfiguration-realtimealertrule-issuedetectionconfiguration
	//
	IssueDetectionConfiguration interface{} `field:"optional" json:"issueDetectionConfiguration" yaml:"issueDetectionConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-realtimealertrule.html#cfn-chime-mediainsightspipelineconfiguration-realtimealertrule-keywordmatchconfiguration
	//
	KeywordMatchConfiguration interface{} `field:"optional" json:"keywordMatchConfiguration" yaml:"keywordMatchConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-realtimealertrule.html#cfn-chime-mediainsightspipelineconfiguration-realtimealertrule-sentimentconfiguration
	//
	SentimentConfiguration interface{} `field:"optional" json:"sentimentConfiguration" yaml:"sentimentConfiguration"`
	// The type of alert rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-realtimealertrule.html#cfn-chime-mediainsightspipelineconfiguration-realtimealertrule-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

