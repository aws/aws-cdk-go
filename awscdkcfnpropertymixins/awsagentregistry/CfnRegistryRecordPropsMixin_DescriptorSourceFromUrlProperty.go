package awsagentregistry


// URL-based descriptor source configuration, with credential provider configurations for authenticated URL retrieval.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   descriptorSourceFromUrlProperty := &DescriptorSourceFromUrlProperty{
//   	CredentialProviderConfigurations: []interface{}{
//   		&RegistryRecordCredentialProviderConfigurationProperty{
//   			CredentialProvider: &RegistryRecordCredentialProviderUnionProperty{
//   				IamCredentialProvider: &RegistryRecordIamCredentialProviderProperty{
//   					Region: jsii.String("region"),
//   					RoleArn: jsii.String("roleArn"),
//   					Service: jsii.String("service"),
//   				},
//   				OauthCredentialProvider: &RegistryRecordOAuthCredentialProviderProperty{
//   					CustomParameters: map[string]*string{
//   						"customParametersKey": jsii.String("customParameters"),
//   					},
//   					GrantType: jsii.String("grantType"),
//   					ProviderArn: jsii.String("providerArn"),
//   					Scopes: []*string{
//   						jsii.String("scopes"),
//   					},
//   				},
//   			},
//   			CredentialProviderType: jsii.String("credentialProviderType"),
//   		},
//   	},
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-descriptorsourcefromurl.html
//
type CfnRegistryRecordPropsMixin_DescriptorSourceFromUrlProperty struct {
	// The credential providers used to authenticate when fetching descriptor content from the source URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-descriptorsourcefromurl.html#cfn-agentregistry-registryrecord-descriptorsourcefromurl-credentialproviderconfigurations
	//
	CredentialProviderConfigurations interface{} `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// URL source for descriptor content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-descriptorsourcefromurl.html#cfn-agentregistry-registryrecord-descriptorsourcefromurl-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}

