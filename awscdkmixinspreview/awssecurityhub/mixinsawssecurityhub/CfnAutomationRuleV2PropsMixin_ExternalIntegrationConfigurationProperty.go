package mixinsawssecurityhub


// The settings for integrating automation rule actions with external systems or service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   externalIntegrationConfigurationProperty := &ExternalIntegrationConfigurationProperty{
//   	ConnectorArn: jsii.String("connectorArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-externalintegrationconfiguration.html
//
type CfnAutomationRuleV2PropsMixin_ExternalIntegrationConfigurationProperty struct {
	// The ARN of the connector that establishes the integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-externalintegrationconfiguration.html#cfn-securityhub-automationrulev2-externalintegrationconfiguration-connectorarn
	//
	ConnectorArn *string `field:"optional" json:"connectorArn" yaml:"connectorArn"`
}

