package awsagentregistry


// The source configuration that defines where descriptor content is retrieved from.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   descriptorSourceProperty := &DescriptorSourceProperty{
//   	FromUrl: &DescriptorSourceFromUrlProperty{
//   		CredentialProviderConfigurations: []interface{}{
//   			&RegistryRecordCredentialProviderConfigurationProperty{
//   				CredentialProvider: &RegistryRecordCredentialProviderUnionProperty{
//   					IamCredentialProvider: &RegistryRecordIamCredentialProviderProperty{
//   						Region: jsii.String("region"),
//   						RoleArn: jsii.String("roleArn"),
//   						Service: jsii.String("service"),
//   					},
//   					OauthCredentialProvider: &RegistryRecordOAuthCredentialProviderProperty{
//   						CustomParameters: map[string]*string{
//   							"customParametersKey": jsii.String("customParameters"),
//   						},
//   						GrantType: jsii.String("grantType"),
//   						ProviderArn: jsii.String("providerArn"),
//   						Scopes: []*string{
//   							jsii.String("scopes"),
//   						},
//   					},
//   				},
//   				CredentialProviderType: jsii.String("credentialProviderType"),
//   			},
//   		},
//   		Url: jsii.String("url"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-descriptorsource.html
//
type CfnRegistryRecordPropsMixin_DescriptorSourceProperty struct {
	// URL-based descriptor source configuration, with credential provider configurations for authenticated URL retrieval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-descriptorsource.html#cfn-agentregistry-registryrecord-descriptorsource-fromurl
	//
	FromUrl interface{} `field:"optional" json:"fromUrl" yaml:"fromUrl"`
}

