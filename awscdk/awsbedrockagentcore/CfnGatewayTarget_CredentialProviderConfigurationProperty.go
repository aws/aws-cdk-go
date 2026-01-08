package awsbedrockagentcore


// The credential provider configuration for the gateway target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   credentialProviderConfigurationProperty := &CredentialProviderConfigurationProperty{
//   	CredentialProviderType: jsii.String("credentialProviderType"),
//
//   	// the properties below are optional
//   	CredentialProvider: &CredentialProviderProperty{
//   		ApiKeyCredentialProvider: &ApiKeyCredentialProviderProperty{
//   			ProviderArn: jsii.String("providerArn"),
//
//   			// the properties below are optional
//   			CredentialLocation: jsii.String("credentialLocation"),
//   			CredentialParameterName: jsii.String("credentialParameterName"),
//   			CredentialPrefix: jsii.String("credentialPrefix"),
//   		},
//   		OauthCredentialProvider: &OAuthCredentialProviderProperty{
//   			ProviderArn: jsii.String("providerArn"),
//   			Scopes: []*string{
//   				jsii.String("scopes"),
//   			},
//
//   			// the properties below are optional
//   			CustomParameters: map[string]*string{
//   				"customParametersKey": jsii.String("customParameters"),
//   			},
//   			DefaultReturnUrl: jsii.String("defaultReturnUrl"),
//   			GrantType: jsii.String("grantType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-credentialproviderconfiguration.html
//
type CfnGatewayTarget_CredentialProviderConfigurationProperty struct {
	// The credential provider type for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-credentialproviderconfiguration.html#cfn-bedrockagentcore-gatewaytarget-credentialproviderconfiguration-credentialprovidertype
	//
	CredentialProviderType *string `field:"required" json:"credentialProviderType" yaml:"credentialProviderType"`
	// The credential provider for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-credentialproviderconfiguration.html#cfn-bedrockagentcore-gatewaytarget-credentialproviderconfiguration-credentialprovider
	//
	CredentialProvider interface{} `field:"optional" json:"credentialProvider" yaml:"credentialProvider"`
}

