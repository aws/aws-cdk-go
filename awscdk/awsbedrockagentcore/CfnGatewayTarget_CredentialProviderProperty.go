package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   credentialProviderProperty := &CredentialProviderProperty{
//   	ApiKeyCredentialProvider: &ApiKeyCredentialProviderProperty{
//   		ProviderArn: jsii.String("providerArn"),
//
//   		// the properties below are optional
//   		CredentialLocation: jsii.String("credentialLocation"),
//   		CredentialParameterName: jsii.String("credentialParameterName"),
//   		CredentialPrefix: jsii.String("credentialPrefix"),
//   	},
//   	OauthCredentialProvider: &OAuthCredentialProviderProperty{
//   		ProviderArn: jsii.String("providerArn"),
//   		Scopes: []*string{
//   			jsii.String("scopes"),
//   		},
//
//   		// the properties below are optional
//   		CustomParameters: map[string]*string{
//   			"customParametersKey": jsii.String("customParameters"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-credentialprovider.html
//
type CfnGatewayTarget_CredentialProviderProperty struct {
	// The API key credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-credentialprovider.html#cfn-bedrockagentcore-gatewaytarget-credentialprovider-apikeycredentialprovider
	//
	ApiKeyCredentialProvider interface{} `field:"optional" json:"apiKeyCredentialProvider" yaml:"apiKeyCredentialProvider"`
	// The OAuth credential provider for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-credentialprovider.html#cfn-bedrockagentcore-gatewaytarget-credentialprovider-oauthcredentialprovider
	//
	OauthCredentialProvider interface{} `field:"optional" json:"oauthCredentialProvider" yaml:"oauthCredentialProvider"`
}

