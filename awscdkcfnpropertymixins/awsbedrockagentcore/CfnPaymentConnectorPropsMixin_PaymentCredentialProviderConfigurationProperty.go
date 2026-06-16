package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   paymentCredentialProviderConfigurationProperty := &PaymentCredentialProviderConfigurationProperty{
//   	CredentialProviderArn: jsii.String("credentialProviderArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentconnector-paymentcredentialproviderconfiguration.html
//
type CfnPaymentConnectorPropsMixin_PaymentCredentialProviderConfigurationProperty struct {
	// The ARN of the payment credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentconnector-paymentcredentialproviderconfiguration.html#cfn-bedrockagentcore-paymentconnector-paymentcredentialproviderconfiguration-credentialproviderarn
	//
	CredentialProviderArn *string `field:"optional" json:"credentialProviderArn" yaml:"credentialProviderArn"`
}

