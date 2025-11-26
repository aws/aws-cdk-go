package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   credentialProviderProperty := &CredentialProviderProperty{
//   	ApiKeyCredentialProvider: &ApiKeyCredentialProviderProperty{
//   		CredentialLocation: jsii.String("credentialLocation"),
//   		CredentialParameterName: jsii.String("credentialParameterName"),
//   		CredentialPrefix: jsii.String("credentialPrefix"),
//   		ProviderArn: jsii.String("providerArn"),
//   	},
//   	OauthCredentialProvider: &OAuthCredentialProviderProperty{
//   		CustomParameters: map[string]*string{
//   			"customParametersKey": jsii.String("customParameters"),
//   		},
//   		ProviderArn: jsii.String("providerArn"),
//   		Scopes: []*string{
//   			jsii.String("scopes"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-credentialprovider.html
//
type CfnGatewayTargetPropsMixin_CredentialProviderProperty struct {
	// The API key credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-credentialprovider.html#cfn-bedrockagentcore-gatewaytarget-credentialprovider-apikeycredentialprovider
	//
	ApiKeyCredentialProvider interface{} `field:"optional" json:"apiKeyCredentialProvider" yaml:"apiKeyCredentialProvider"`
	// The OAuth credential provider for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-credentialprovider.html#cfn-bedrockagentcore-gatewaytarget-credentialprovider-oauthcredentialprovider
	//
	OauthCredentialProvider interface{} `field:"optional" json:"oauthCredentialProvider" yaml:"oauthCredentialProvider"`
}

