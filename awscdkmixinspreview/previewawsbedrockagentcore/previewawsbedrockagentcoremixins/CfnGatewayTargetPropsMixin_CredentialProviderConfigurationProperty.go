package previewawsbedrockagentcoremixins


// The credential provider configuration for the gateway target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   credentialProviderConfigurationProperty := &CredentialProviderConfigurationProperty{
//   	CredentialProvider: &CredentialProviderProperty{
//   		ApiKeyCredentialProvider: &ApiKeyCredentialProviderProperty{
//   			CredentialLocation: jsii.String("credentialLocation"),
//   			CredentialParameterName: jsii.String("credentialParameterName"),
//   			CredentialPrefix: jsii.String("credentialPrefix"),
//   			ProviderArn: jsii.String("providerArn"),
//   		},
//   		OauthCredentialProvider: &OAuthCredentialProviderProperty{
//   			CustomParameters: map[string]*string{
//   				"customParametersKey": jsii.String("customParameters"),
//   			},
//   			ProviderArn: jsii.String("providerArn"),
//   			Scopes: []*string{
//   				jsii.String("scopes"),
//   			},
//   		},
//   	},
//   	CredentialProviderType: jsii.String("credentialProviderType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-credentialproviderconfiguration.html
//
type CfnGatewayTargetPropsMixin_CredentialProviderConfigurationProperty struct {
	// The credential provider for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-credentialproviderconfiguration.html#cfn-bedrockagentcore-gatewaytarget-credentialproviderconfiguration-credentialprovider
	//
	CredentialProvider interface{} `field:"optional" json:"credentialProvider" yaml:"credentialProvider"`
	// The credential provider type for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-credentialproviderconfiguration.html#cfn-bedrockagentcore-gatewaytarget-credentialproviderconfiguration-credentialprovidertype
	//
	CredentialProviderType *string `field:"optional" json:"credentialProviderType" yaml:"credentialProviderType"`
}

