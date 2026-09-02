package awsagentregistry


// The MCP server descriptor, populated when the record type is MCP.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mcpServerDescriptorProperty := &McpServerDescriptorProperty{
//   	AdditionalData: &McpServerAdditionalDataProperty{
//   		Tools: &McpToolsDescriptorProperty{
//   			Data: jsii.String("data"),
//   			DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   		},
//   	},
//   	Data: jsii.String("data"),
//   	DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   	Source: &DescriptorSourceProperty{
//   		FromUrl: &DescriptorSourceFromUrlProperty{
//   			Url: jsii.String("url"),
//
//   			// the properties below are optional
//   			CredentialProviderConfigurations: []interface{}{
//   				&RegistryRecordCredentialProviderConfigurationProperty{
//   					CredentialProvider: &RegistryRecordCredentialProviderUnionProperty{
//   						IamCredentialProvider: &RegistryRecordIamCredentialProviderProperty{
//   							Region: jsii.String("region"),
//   							RoleArn: jsii.String("roleArn"),
//   							Service: jsii.String("service"),
//   						},
//   						OauthCredentialProvider: &RegistryRecordOAuthCredentialProviderProperty{
//   							ProviderArn: jsii.String("providerArn"),
//
//   							// the properties below are optional
//   							CustomParameters: map[string]*string{
//   								"customParametersKey": jsii.String("customParameters"),
//   							},
//   							GrantType: jsii.String("grantType"),
//   							Scopes: []*string{
//   								jsii.String("scopes"),
//   							},
//   						},
//   					},
//   					CredentialProviderType: jsii.String("credentialProviderType"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-mcpserverdescriptor.html
//
type CfnRegistryRecord_McpServerDescriptorProperty struct {
	// Additional data associated with an MCP server descriptor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-mcpserverdescriptor.html#cfn-agentregistry-registryrecord-mcpserverdescriptor-additionaldata
	//
	AdditionalData interface{} `field:"optional" json:"additionalData" yaml:"additionalData"`
	// Descriptor payload data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-mcpserverdescriptor.html#cfn-agentregistry-registryrecord-mcpserverdescriptor-data
	//
	Data *string `field:"optional" json:"data" yaml:"data"`
	// Version of the descriptor type schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-mcpserverdescriptor.html#cfn-agentregistry-registryrecord-mcpserverdescriptor-dataschemaversion
	//
	DataSchemaVersion *string `field:"optional" json:"dataSchemaVersion" yaml:"dataSchemaVersion"`
	// The source configuration that defines where descriptor content is retrieved from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-mcpserverdescriptor.html#cfn-agentregistry-registryrecord-mcpserverdescriptor-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

