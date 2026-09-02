package awsagentregistry


// The credential provider details.
//
// Specify exactly one member.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   registryRecordCredentialProviderUnionProperty := &RegistryRecordCredentialProviderUnionProperty{
//   	IamCredentialProvider: &RegistryRecordIamCredentialProviderProperty{
//   		Region: jsii.String("region"),
//   		RoleArn: jsii.String("roleArn"),
//   		Service: jsii.String("service"),
//   	},
//   	OauthCredentialProvider: &RegistryRecordOAuthCredentialProviderProperty{
//   		CustomParameters: map[string]*string{
//   			"customParametersKey": jsii.String("customParameters"),
//   		},
//   		GrantType: jsii.String("grantType"),
//   		ProviderArn: jsii.String("providerArn"),
//   		Scopes: []*string{
//   			jsii.String("scopes"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-registryrecordcredentialproviderunion.html
//
type CfnRegistryRecordPropsMixin_RegistryRecordCredentialProviderUnionProperty struct {
	// IAM credential provider configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-registryrecordcredentialproviderunion.html#cfn-agentregistry-registryrecord-registryrecordcredentialproviderunion-iamcredentialprovider
	//
	IamCredentialProvider interface{} `field:"optional" json:"iamCredentialProvider" yaml:"iamCredentialProvider"`
	// OAuth credential provider configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-registryrecordcredentialproviderunion.html#cfn-agentregistry-registryrecord-registryrecordcredentialproviderunion-oauthcredentialprovider
	//
	OauthCredentialProvider interface{} `field:"optional" json:"oauthCredentialProvider" yaml:"oauthCredentialProvider"`
}

