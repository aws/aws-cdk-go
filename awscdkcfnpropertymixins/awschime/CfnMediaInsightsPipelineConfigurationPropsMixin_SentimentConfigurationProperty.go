package awschime


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sentimentConfigurationProperty := &SentimentConfigurationProperty{
//   	RuleName: jsii.String("ruleName"),
//   	SentimentType: jsii.String("sentimentType"),
//   	TimePeriod: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-sentimentconfiguration.html
//
type CfnMediaInsightsPipelineConfigurationPropsMixin_SentimentConfigurationProperty struct {
	// The name of the sentiment rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-sentimentconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-sentimentconfiguration-rulename
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The type of sentiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-sentimentconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-sentimentconfiguration-sentimenttype
	//
	SentimentType *string `field:"optional" json:"sentimentType" yaml:"sentimentType"`
	// The analysis interval in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-sentimentconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-sentimentconfiguration-timeperiod
	//
	TimePeriod *float64 `field:"optional" json:"timePeriod" yaml:"timePeriod"`
}

