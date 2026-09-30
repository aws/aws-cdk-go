package awschime


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   issueDetectionConfigurationProperty := &IssueDetectionConfigurationProperty{
//   	RuleName: jsii.String("ruleName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-issuedetectionconfiguration.html
//
type CfnMediaInsightsPipelineConfigurationPropsMixin_IssueDetectionConfigurationProperty struct {
	// The name of the issue detection rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-issuedetectionconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-issuedetectionconfiguration-rulename
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
}

