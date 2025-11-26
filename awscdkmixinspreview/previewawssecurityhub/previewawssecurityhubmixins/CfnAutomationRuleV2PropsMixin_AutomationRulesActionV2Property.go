package previewawssecurityhubmixins


// Allows you to configure automated responses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   automationRulesActionV2Property := &AutomationRulesActionV2Property{
//   	ExternalIntegrationConfiguration: &ExternalIntegrationConfigurationProperty{
//   		ConnectorArn: jsii.String("connectorArn"),
//   	},
//   	FindingFieldsUpdate: &AutomationRulesFindingFieldsUpdateV2Property{
//   		Comment: jsii.String("comment"),
//   		SeverityId: jsii.Number(123),
//   		StatusId: jsii.Number(123),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-automationrulesactionv2.html
//
type CfnAutomationRuleV2PropsMixin_AutomationRulesActionV2Property struct {
	// The settings for integrating automation rule actions with external systems or service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-automationrulesactionv2.html#cfn-securityhub-automationrulev2-automationrulesactionv2-externalintegrationconfiguration
	//
	ExternalIntegrationConfiguration interface{} `field:"optional" json:"externalIntegrationConfiguration" yaml:"externalIntegrationConfiguration"`
	// Specifies that the automation rule action is an update to a finding field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-automationrulesactionv2.html#cfn-securityhub-automationrulev2-automationrulesactionv2-findingfieldsupdate
	//
	FindingFieldsUpdate interface{} `field:"optional" json:"findingFieldsUpdate" yaml:"findingFieldsUpdate"`
	// Specifies the type of action that Security Hub takes when a finding matches the defined criteria of a rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-automationrulesactionv2.html#cfn-securityhub-automationrulev2-automationrulesactionv2-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

