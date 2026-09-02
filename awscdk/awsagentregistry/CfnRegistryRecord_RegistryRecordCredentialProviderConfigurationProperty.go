package awsagentregistry


// A credential provider configuration used for authenticated descriptor retrieval.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registryRecordCredentialProviderConfigurationProperty := &RegistryRecordCredentialProviderConfigurationProperty{
//   	CredentialProvider: &RegistryRecordCredentialProviderUnionProperty{
//   		IamCredentialProvider: &RegistryRecordIamCredentialProviderProperty{
//   			Region: jsii.String("region"),
//   			RoleArn: jsii.String("roleArn"),
//   			Service: jsii.String("service"),
//   		},
//   		OauthCredentialProvider: &RegistryRecordOAuthCredentialProviderProperty{
//   			ProviderArn: jsii.String("providerArn"),
//
//   			// the properties below are optional
//   			CustomParameters: map[string]*string{
//   				"customParametersKey": jsii.String("customParameters"),
//   			},
//   			GrantType: jsii.String("grantType"),
//   			Scopes: []*string{
//   				jsii.String("scopes"),
//   			},
//   		},
//   	},
//   	CredentialProviderType: jsii.String("credentialProviderType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-registryrecordcredentialproviderconfiguration.html
//
type CfnRegistryRecord_RegistryRecordCredentialProviderConfigurationProperty struct {
	// The credential provider details.
	//
	// Specify exactly one member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-registryrecordcredentialproviderconfiguration.html#cfn-agentregistry-registryrecord-registryrecordcredentialproviderconfiguration-credentialprovider
	//
	CredentialProvider interface{} `field:"required" json:"credentialProvider" yaml:"credentialProvider"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-registryrecordcredentialproviderconfiguration.html#cfn-agentregistry-registryrecord-registryrecordcredentialproviderconfiguration-credentialprovidertype
	//
	CredentialProviderType *string `field:"required" json:"credentialProviderType" yaml:"credentialProviderType"`
}

