package awsagentregistry


// The typed set of descriptors for a registry record.
//
// Exactly one descriptor field is populated based on the record type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   descriptorsProperty := &DescriptorsProperty{
//   	A2AAgentCard: &A2aAgentCardDescriptorProperty{
//   		Data: jsii.String("data"),
//   		DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   		Source: &DescriptorSourceProperty{
//   			FromUrl: &DescriptorSourceFromUrlProperty{
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				CredentialProviderConfigurations: []interface{}{
//   					&RegistryRecordCredentialProviderConfigurationProperty{
//   						CredentialProvider: &RegistryRecordCredentialProviderUnionProperty{
//   							IamCredentialProvider: &RegistryRecordIamCredentialProviderProperty{
//   								Region: jsii.String("region"),
//   								RoleArn: jsii.String("roleArn"),
//   								Service: jsii.String("service"),
//   							},
//   							OauthCredentialProvider: &RegistryRecordOAuthCredentialProviderProperty{
//   								ProviderArn: jsii.String("providerArn"),
//
//   								// the properties below are optional
//   								CustomParameters: map[string]*string{
//   									"customParametersKey": jsii.String("customParameters"),
//   								},
//   								GrantType: jsii.String("grantType"),
//   								Scopes: []*string{
//   									jsii.String("scopes"),
//   								},
//   							},
//   						},
//   						CredentialProviderType: jsii.String("credentialProviderType"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	AgentSkillsDefinition: &AgentSkillsDefinitionDescriptorProperty{
//   		AdditionalData: &AgentSkillsAdditionalDataProperty{
//   			SkillMd: &AgentSkillsMdDescriptorProperty{
//   				Data: jsii.String("data"),
//   				DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   				Source: &SkillMdSourceProperty{
//   					FromUrl: &SkillMdSourceFromUrlProperty{
//   						Url: jsii.String("url"),
//   					},
//   				},
//   			},
//   		},
//   		Data: jsii.String("data"),
//   		DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   	},
//   	Custom: &CustomDescriptorProperty{
//   		Data: jsii.String("data"),
//   	},
//   	McpServer: &McpServerDescriptorProperty{
//   		AdditionalData: &McpServerAdditionalDataProperty{
//   			Tools: &McpToolsDescriptorProperty{
//   				Data: jsii.String("data"),
//   				DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   			},
//   		},
//   		Data: jsii.String("data"),
//   		DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   		Source: &DescriptorSourceProperty{
//   			FromUrl: &DescriptorSourceFromUrlProperty{
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				CredentialProviderConfigurations: []interface{}{
//   					&RegistryRecordCredentialProviderConfigurationProperty{
//   						CredentialProvider: &RegistryRecordCredentialProviderUnionProperty{
//   							IamCredentialProvider: &RegistryRecordIamCredentialProviderProperty{
//   								Region: jsii.String("region"),
//   								RoleArn: jsii.String("roleArn"),
//   								Service: jsii.String("service"),
//   							},
//   							OauthCredentialProvider: &RegistryRecordOAuthCredentialProviderProperty{
//   								ProviderArn: jsii.String("providerArn"),
//
//   								// the properties below are optional
//   								CustomParameters: map[string]*string{
//   									"customParametersKey": jsii.String("customParameters"),
//   								},
//   								GrantType: jsii.String("grantType"),
//   								Scopes: []*string{
//   									jsii.String("scopes"),
//   								},
//   							},
//   						},
//   						CredentialProviderType: jsii.String("credentialProviderType"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-descriptors.html
//
type CfnRegistryRecord_DescriptorsProperty struct {
	// The A2A agent card descriptor, populated when the record type is AGENT.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-descriptors.html#cfn-agentregistry-registryrecord-descriptors-a2aagentcard
	//
	A2AAgentCard interface{} `field:"optional" json:"a2AAgentCard" yaml:"a2AAgentCard"`
	// The agent skills definition descriptor, populated when the record type is SKILL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-descriptors.html#cfn-agentregistry-registryrecord-descriptors-agentskillsdefinition
	//
	AgentSkillsDefinition interface{} `field:"optional" json:"agentSkillsDefinition" yaml:"agentSkillsDefinition"`
	// The custom descriptor, populated when the record type is CUSTOM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-descriptors.html#cfn-agentregistry-registryrecord-descriptors-custom
	//
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// The MCP server descriptor, populated when the record type is MCP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-descriptors.html#cfn-agentregistry-registryrecord-descriptors-mcpserver
	//
	McpServer interface{} `field:"optional" json:"mcpServer" yaml:"mcpServer"`
}

