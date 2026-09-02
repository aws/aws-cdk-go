package awsagentregistry


// The A2A agent card descriptor, populated when the record type is AGENT.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   a2aAgentCardDescriptorProperty := &A2aAgentCardDescriptorProperty{
//   	Data: jsii.String("data"),
//   	DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   	Source: &DescriptorSourceProperty{
//   		FromUrl: &DescriptorSourceFromUrlProperty{
//   			CredentialProviderConfigurations: []interface{}{
//   				&RegistryRecordCredentialProviderConfigurationProperty{
//   					CredentialProvider: &RegistryRecordCredentialProviderUnionProperty{
//   						IamCredentialProvider: &RegistryRecordIamCredentialProviderProperty{
//   							Region: jsii.String("region"),
//   							RoleArn: jsii.String("roleArn"),
//   							Service: jsii.String("service"),
//   						},
//   						OauthCredentialProvider: &RegistryRecordOAuthCredentialProviderProperty{
//   							CustomParameters: map[string]*string{
//   								"customParametersKey": jsii.String("customParameters"),
//   							},
//   							GrantType: jsii.String("grantType"),
//   							ProviderArn: jsii.String("providerArn"),
//   							Scopes: []*string{
//   								jsii.String("scopes"),
//   							},
//   						},
//   					},
//   					CredentialProviderType: jsii.String("credentialProviderType"),
//   				},
//   			},
//   			Url: jsii.String("url"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-a2aagentcarddescriptor.html
//
type CfnRegistryRecordPropsMixin_A2aAgentCardDescriptorProperty struct {
	// Descriptor payload data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-a2aagentcarddescriptor.html#cfn-agentregistry-registryrecord-a2aagentcarddescriptor-data
	//
	Data *string `field:"optional" json:"data" yaml:"data"`
	// Version of the descriptor type schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-a2aagentcarddescriptor.html#cfn-agentregistry-registryrecord-a2aagentcarddescriptor-dataschemaversion
	//
	DataSchemaVersion *string `field:"optional" json:"dataSchemaVersion" yaml:"dataSchemaVersion"`
	// The source configuration that defines where descriptor content is retrieved from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-a2aagentcarddescriptor.html#cfn-agentregistry-registryrecord-a2aagentcarddescriptor-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

