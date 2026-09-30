package awschime


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   realTimeAlertConfigurationProperty := &RealTimeAlertConfigurationProperty{
//   	Disabled: jsii.Boolean(false),
//   	Rules: []interface{}{
//   		&RealTimeAlertRuleProperty{
//   			IssueDetectionConfiguration: &IssueDetectionConfigurationProperty{
//   				RuleName: jsii.String("ruleName"),
//   			},
//   			KeywordMatchConfiguration: &KeywordMatchConfigurationProperty{
//   				Keywords: []*string{
//   					jsii.String("keywords"),
//   				},
//   				Negate: jsii.Boolean(false),
//   				RuleName: jsii.String("ruleName"),
//   			},
//   			SentimentConfiguration: &SentimentConfigurationProperty{
//   				RuleName: jsii.String("ruleName"),
//   				SentimentType: jsii.String("sentimentType"),
//   				TimePeriod: jsii.Number(123),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-realtimealertconfiguration.html
//
type CfnMediaInsightsPipelineConfigurationPropsMixin_RealTimeAlertConfigurationProperty struct {
	// Turns off real-time alerts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-realtimealertconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-realtimealertconfiguration-disabled
	//
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// The rules in the alert.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-realtimealertconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-realtimealertconfiguration-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

