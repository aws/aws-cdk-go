package awsagentregistry


// OAuth credential provider configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registryRecordOAuthCredentialProviderProperty := &RegistryRecordOAuthCredentialProviderProperty{
//   	ProviderArn: jsii.String("providerArn"),
//
//   	// the properties below are optional
//   	CustomParameters: map[string]*string{
//   		"customParametersKey": jsii.String("customParameters"),
//   	},
//   	GrantType: jsii.String("grantType"),
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-registryrecordoauthcredentialprovider.html
//
type CfnRegistryRecord_RegistryRecordOAuthCredentialProviderProperty struct {
	// The ARN of the OAuth credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-registryrecordoauthcredentialprovider.html#cfn-agentregistry-registryrecord-registryrecordoauthcredentialprovider-providerarn
	//
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// Additional custom parameters for the OAuth flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-registryrecordoauthcredentialprovider.html#cfn-agentregistry-registryrecord-registryrecordoauthcredentialprovider-customparameters
	//
	CustomParameters interface{} `field:"optional" json:"customParameters" yaml:"customParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-registryrecordoauthcredentialprovider.html#cfn-agentregistry-registryrecord-registryrecordoauthcredentialprovider-granttype
	//
	GrantType *string `field:"optional" json:"grantType" yaml:"grantType"`
	// OAuth scopes to request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-registryrecordoauthcredentialprovider.html#cfn-agentregistry-registryrecord-registryrecordoauthcredentialprovider-scopes
	//
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

