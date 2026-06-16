package awsbedrockagentcore


// Stripe Privy configuration output with secret ARNs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//
//   	// the properties below are optional
//   	AppSecretJsonKey: jsii.String("appSecretJsonKey"),
//   	AppSecretSource: jsii.String("appSecretSource"),
//   	AuthorizationPrivateKeyJsonKey: jsii.String("authorizationPrivateKeyJsonKey"),
//   	AuthorizationPrivateKeySource: jsii.String("authorizationPrivateKeySource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput.html
//
type CfnPaymentCredentialProvider_StripePrivyConfigurationOutputProperty struct {
	// The app ID provided by Privy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput-appid
	//
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Contains information about a secret in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput-appsecretarn
	//
	AppSecretArn interface{} `field:"required" json:"appSecretArn" yaml:"appSecretArn"`
	// The authorization ID for the Stripe Privy integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput-authorizationid
	//
	AuthorizationId *string `field:"required" json:"authorizationId" yaml:"authorizationId"`
	// Contains information about a secret in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput-authorizationprivatekeyarn
	//
	AuthorizationPrivateKeyArn interface{} `field:"required" json:"authorizationPrivateKeyArn" yaml:"authorizationPrivateKeyArn"`
	// The JSON key within the secret that contains the app secret value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput-appsecretjsonkey
	//
	AppSecretJsonKey *string `field:"optional" json:"appSecretJsonKey" yaml:"appSecretJsonKey"`
	// The source of the secret.
	//
	// Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput-appsecretsource
	//
	AppSecretSource *string `field:"optional" json:"appSecretSource" yaml:"appSecretSource"`
	// The JSON key within the secret that contains the authorization private key value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput-authorizationprivatekeyjsonkey
	//
	AuthorizationPrivateKeyJsonKey *string `field:"optional" json:"authorizationPrivateKeyJsonKey" yaml:"authorizationPrivateKeyJsonKey"`
	// The source of the secret.
	//
	// Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-stripeprivyconfigurationoutput-authorizationprivatekeysource
	//
	AuthorizationPrivateKeySource *string `field:"optional" json:"authorizationPrivateKeySource" yaml:"authorizationPrivateKeySource"`
}

