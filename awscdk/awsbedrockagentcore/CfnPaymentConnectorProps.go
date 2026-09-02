package awsbedrockagentcore


// Properties for defining a `CfnPaymentConnector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPaymentConnectorProps := &CfnPaymentConnectorProps{
//   	ConnectorName: jsii.String("connectorName"),
//   	ConnectorType: jsii.String("connectorType"),
//   	PaymentManagerId: jsii.String("paymentManagerId"),
//
//   	// the properties below are optional
//   	CredentialProviderConfigurations: []interface{}{
//   		&CredentialsProviderConfigurationProperty{
//   			CoinbaseCdp: &PaymentCredentialProviderConfigurationProperty{
//   				CredentialProviderArn: jsii.String("credentialProviderArn"),
//   			},
//   			StripePrivy: &PaymentCredentialProviderConfigurationProperty{
//   				CredentialProviderArn: jsii.String("credentialProviderArn"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ProvisionMode: jsii.String("provisionMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html
//
type CfnPaymentConnectorProps struct {
	// The name of the payment connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html#cfn-bedrockagentcore-paymentconnector-connectorname
	//
	ConnectorName *string `field:"required" json:"connectorName" yaml:"connectorName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html#cfn-bedrockagentcore-paymentconnector-connectortype
	//
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// The identifier of the parent payment manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html#cfn-bedrockagentcore-paymentconnector-paymentmanagerid
	//
	PaymentManagerId *string `field:"required" json:"paymentManagerId" yaml:"paymentManagerId"`
	// The credential provider configurations for the connector.
	//
	// Required when ProvisionMode is MANUAL or not specified. Empty for QUICK_CREATE until provisioning completes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html#cfn-bedrockagentcore-paymentconnector-credentialproviderconfigurations
	//
	CredentialProviderConfigurations interface{} `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// A description of the payment connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html#cfn-bedrockagentcore-paymentconnector-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html#cfn-bedrockagentcore-paymentconnector-provisionmode
	//
	ProvisionMode *string `field:"optional" json:"provisionMode" yaml:"provisionMode"`
}

