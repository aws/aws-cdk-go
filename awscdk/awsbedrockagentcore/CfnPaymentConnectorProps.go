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
//   	PaymentManagerId: jsii.String("paymentManagerId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
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
	// The credential provider configurations for the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html#cfn-bedrockagentcore-paymentconnector-credentialproviderconfigurations
	//
	CredentialProviderConfigurations interface{} `field:"required" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// The identifier of the parent payment manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html#cfn-bedrockagentcore-paymentconnector-paymentmanagerid
	//
	PaymentManagerId *string `field:"required" json:"paymentManagerId" yaml:"paymentManagerId"`
	// A description of the payment connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentconnector.html#cfn-bedrockagentcore-paymentconnector-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

