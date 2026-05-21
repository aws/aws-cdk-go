package awsbedrockagentcore


// Stripe Privy configuration output with secret ARNs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   stripePrivyConfigurationOutputProperty := &StripePrivyConfigurationOutputProperty{
//   	AppId: jsii.String("appId"),
//   	AppSecretArn: &SecretInfoProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	AuthorizationId: jsii.String("authorizationId"),
//   	AuthorizationPrivateKeyArn: &SecretInfoProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput.html
//
type CfnPaymentCredentialProviderPropsMixin_StripePrivyConfigurationOutputProperty struct {
	// The app ID provided by Privy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput-appid
	//
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// Contains information about a secret in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput-appsecretarn
	//
	AppSecretArn interface{} `field:"optional" json:"appSecretArn" yaml:"appSecretArn"`
	// The authorization ID for the Stripe Privy integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput-authorizationid
	//
	AuthorizationId *string `field:"optional" json:"authorizationId" yaml:"authorizationId"`
	// Contains information about a secret in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput-authorizationprivatekeyarn
	//
	AuthorizationPrivateKeyArn interface{} `field:"optional" json:"authorizationPrivateKeyArn" yaml:"authorizationPrivateKeyArn"`
}

