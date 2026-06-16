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
//   	AppSecretConfig: &SecretReferenceProperty{
//   		JsonKey: jsii.String("jsonKey"),
//   		SecretId: jsii.String("secretId"),
//   	},
//   	AppSecretSource: jsii.String("appSecretSource"),
//   	AuthorizationId: jsii.String("authorizationId"),
//   	AuthorizationPrivateKey: jsii.String("authorizationPrivateKey"),
//   	AuthorizationPrivateKeyConfig: &SecretReferenceProperty{
//   		JsonKey: jsii.String("jsonKey"),
//   		SecretId: jsii.String("secretId"),
//   	},
//   	AuthorizationPrivateKeySource: jsii.String("authorizationPrivateKeySource"),
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
	// A reference to a customer-provided secret stored in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput-appsecretconfig
	//
	AppSecretConfig interface{} `field:"optional" json:"appSecretConfig" yaml:"appSecretConfig"`
	// The source of the secret.
	//
	// Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput-appsecretsource
	//
	AppSecretSource *string `field:"optional" json:"appSecretSource" yaml:"appSecretSource"`
	// The authorization ID for the Stripe Privy integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput-authorizationid
	//
	AuthorizationId *string `field:"optional" json:"authorizationId" yaml:"authorizationId"`
	// The authorization private key for the Stripe Privy integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput-authorizationprivatekey
	//
	AuthorizationPrivateKey *string `field:"optional" json:"authorizationPrivateKey" yaml:"authorizationPrivateKey"`
	// A reference to a customer-provided secret stored in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput-authorizationprivatekeyconfig
	//
	AuthorizationPrivateKeyConfig interface{} `field:"optional" json:"authorizationPrivateKeyConfig" yaml:"authorizationPrivateKeyConfig"`
	// The source of the secret.
	//
	// Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationinput-authorizationprivatekeysource
	//
	AuthorizationPrivateKeySource *string `field:"optional" json:"authorizationPrivateKeySource" yaml:"authorizationPrivateKeySource"`
}

