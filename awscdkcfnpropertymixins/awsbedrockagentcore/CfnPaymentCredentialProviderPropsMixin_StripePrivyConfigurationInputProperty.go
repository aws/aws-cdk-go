package awsbedrockagentcore


// Stripe Privy configuration with credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   stripePrivyConfigurationInputProperty := &StripePrivyConfigurationInputProperty{
//   	AppId: jsii.String("appId"),
//   	AppSecret: jsii.String("appSecret"),
//   	AuthorizationId: jsii.String("authorizationId"),
//   	AuthorizationPrivateKey: jsii.String("authorizationPrivateKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput.html
//
type CfnPaymentCredentialProviderPropsMixin_StripePrivyConfigurationInputProperty struct {
	// The app ID provided by Privy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput-appid
	//
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// The app secret provided by Privy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput-appsecret
	//
	AppSecret *string `field:"optional" json:"appSecret" yaml:"appSecret"`
	// The authorization ID for the Stripe Privy integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput-authorizationid
	//
	AuthorizationId *string `field:"optional" json:"authorizationId" yaml:"authorizationId"`
	// The authorization private key for the Stripe Privy integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput-authorizationprivatekey
	//
	AuthorizationPrivateKey *string `field:"optional" json:"authorizationPrivateKey" yaml:"authorizationPrivateKey"`
}

