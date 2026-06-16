package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   credentialsProviderConfigurationProperty := &CredentialsProviderConfigurationProperty{
//   	CoinbaseCdp: &PaymentCredentialProviderConfigurationProperty{
//   		CredentialProviderArn: jsii.String("credentialProviderArn"),
//   	},
//   	StripePrivy: &PaymentCredentialProviderConfigurationProperty{
//   		CredentialProviderArn: jsii.String("credentialProviderArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentconnector-credentialsproviderconfiguration.html
//
type CfnPaymentConnector_CredentialsProviderConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentconnector-credentialsproviderconfiguration.html#cfn-bedrockagentcore-paymentconnector-credentialsproviderconfiguration-coinbasecdp
	//
	CoinbaseCdp interface{} `field:"optional" json:"coinbaseCdp" yaml:"coinbaseCdp"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentconnector-credentialsproviderconfiguration.html#cfn-bedrockagentcore-paymentconnector-credentialsproviderconfiguration-stripeprivy
	//
	StripePrivy interface{} `field:"optional" json:"stripePrivy" yaml:"stripePrivy"`
}

