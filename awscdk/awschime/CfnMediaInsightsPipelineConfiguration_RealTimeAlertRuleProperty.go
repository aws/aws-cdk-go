package awschime


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   realTimeAlertRuleProperty := &RealTimeAlertRuleProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	IssueDetectionConfiguration: &IssueDetectionConfigurationProperty{
//   		RuleName: jsii.String("ruleName"),
//   	},
//   	KeywordMatchConfiguration: &KeywordMatchConfigurationProperty{
//   		Keywords: []*string{
//   			jsii.String("keywords"),
//   		},
//   		RuleName: jsii.String("ruleName"),
//
//   		// the properties below are optional
//   		Negate: jsii.Boolean(false),
//   	},
//   	SentimentConfiguration: &SentimentConfigurationProperty{
//   		RuleName: jsii.String("ruleName"),
//   		SentimentType: jsii.String("sentimentType"),
//   		TimePeriod: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-realtimealertrule.html
//
type CfnMediaInsightsPipelineConfiguration_RealTimeAlertRuleProperty struct {
	// The type of alert rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-realtimealertrule.html#cfn-chime-mediainsightspipelineconfiguration-realtimealertrule-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-realtimealertrule.html#cfn-chime-mediainsightspipelineconfiguration-realtimealertrule-issuedetectionconfiguration
	//
	IssueDetectionConfiguration interface{} `field:"optional" json:"issueDetectionConfiguration" yaml:"issueDetectionConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-realtimealertrule.html#cfn-chime-mediainsightspipelineconfiguration-realtimealertrule-keywordmatchconfiguration
	//
	KeywordMatchConfiguration interface{} `field:"optional" json:"keywordMatchConfiguration" yaml:"keywordMatchConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-realtimealertrule.html#cfn-chime-mediainsightspipelineconfiguration-realtimealertrule-sentimentconfiguration
	//
	SentimentConfiguration interface{} `field:"optional" json:"sentimentConfiguration" yaml:"sentimentConfiguration"`
}

