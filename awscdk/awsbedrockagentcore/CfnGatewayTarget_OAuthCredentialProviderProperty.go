package awsbedrockagentcore


// The OAuth credential provider for the gateway target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oAuthCredentialProviderProperty := &OAuthCredentialProviderProperty{
//   	ProviderArn: jsii.String("providerArn"),
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//
//   	// the properties below are optional
//   	CustomParameters: map[string]*string{
//   		"customParametersKey": jsii.String("customParameters"),
//   	},
//   	DefaultReturnUrl: jsii.String("defaultReturnUrl"),
//   	GrantType: jsii.String("grantType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-oauthcredentialprovider.html
//
type CfnGatewayTarget_OAuthCredentialProviderProperty struct {
	// The provider ARN for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-oauthcredentialprovider.html#cfn-bedrockagentcore-gatewaytarget-oauthcredentialprovider-providerarn
	//
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// The OAuth credential provider scopes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-oauthcredentialprovider.html#cfn-bedrockagentcore-gatewaytarget-oauthcredentialprovider-scopes
	//
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The OAuth credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-oauthcredentialprovider.html#cfn-bedrockagentcore-gatewaytarget-oauthcredentialprovider-customparameters
	//
	CustomParameters interface{} `field:"optional" json:"customParameters" yaml:"customParameters"`
	// Return URL for OAuth callback.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-oauthcredentialprovider.html#cfn-bedrockagentcore-gatewaytarget-oauthcredentialprovider-defaultreturnurl
	//
	DefaultReturnUrl *string `field:"optional" json:"defaultReturnUrl" yaml:"defaultReturnUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-oauthcredentialprovider.html#cfn-bedrockagentcore-gatewaytarget-oauthcredentialprovider-granttype
	//
	GrantType *string `field:"optional" json:"grantType" yaml:"grantType"`
}

