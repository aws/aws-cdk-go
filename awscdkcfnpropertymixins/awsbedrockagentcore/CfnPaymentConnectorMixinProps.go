package awsbedrockagentcore


// Properties for CfnPaymentConnectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnPaymentConnectorMixinProps := &CfnPaymentConnectorMixinProps{
//   	ConnectorName: jsii.String("connectorName"),
//   	ConnectorType: jsii.String("connectorType"),
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
//   	PaymentManagerId: jsii.String("paymentManagerId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html
//
type CfnPaymentConnectorMixinProps struct {
	// The name of the payment connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html#cfn-bedrockagentcore-paymentconnector-connectorname
	//
	ConnectorName *string `field:"optional" json:"connectorName" yaml:"connectorName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html#cfn-bedrockagentcore-paymentconnector-connectortype
	//
	ConnectorType *string `field:"optional" json:"connectorType" yaml:"connectorType"`
	// The credential provider configurations for the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html#cfn-bedrockagentcore-paymentconnector-credentialproviderconfigurations
	//
	CredentialProviderConfigurations interface{} `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// A description of the payment connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html#cfn-bedrockagentcore-paymentconnector-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identifier of the parent payment manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html#cfn-bedrockagentcore-paymentconnector-paymentmanagerid
	//
	PaymentManagerId *string `field:"optional" json:"paymentManagerId" yaml:"paymentManagerId"`
}

