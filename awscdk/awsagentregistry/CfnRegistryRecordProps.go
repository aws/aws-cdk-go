package awsagentregistry

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRegistryRecord`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRegistryRecordProps := &CfnRegistryRecordProps{
//   	Descriptors: &DescriptorsProperty{
//   		A2AAgentCard: &A2aAgentCardDescriptorProperty{
//   			Data: jsii.String("data"),
//   			DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   			Source: &DescriptorSourceProperty{
//   				FromUrl: &DescriptorSourceFromUrlProperty{
//   					Url: jsii.String("url"),
//
//   					// the properties below are optional
//   					CredentialProviderConfigurations: []interface{}{
//   						&RegistryRecordCredentialProviderConfigurationProperty{
//   							CredentialProvider: &RegistryRecordCredentialProviderUnionProperty{
//   								IamCredentialProvider: &RegistryRecordIamCredentialProviderProperty{
//   									Region: jsii.String("region"),
//   									RoleArn: jsii.String("roleArn"),
//   									Service: jsii.String("service"),
//   								},
//   								OauthCredentialProvider: &RegistryRecordOAuthCredentialProviderProperty{
//   									ProviderArn: jsii.String("providerArn"),
//
//   									// the properties below are optional
//   									CustomParameters: map[string]*string{
//   										"customParametersKey": jsii.String("customParameters"),
//   									},
//   									GrantType: jsii.String("grantType"),
//   									Scopes: []*string{
//   										jsii.String("scopes"),
//   									},
//   								},
//   							},
//   							CredentialProviderType: jsii.String("credentialProviderType"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		AgentSkillsDefinition: &AgentSkillsDefinitionDescriptorProperty{
//   			AdditionalData: &AgentSkillsAdditionalDataProperty{
//   				SkillMd: &AgentSkillsMdDescriptorProperty{
//   					Data: jsii.String("data"),
//   					DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   					Source: &SkillMdSourceProperty{
//   						FromUrl: &SkillMdSourceFromUrlProperty{
//   							Url: jsii.String("url"),
//   						},
//   					},
//   				},
//   			},
//   			Data: jsii.String("data"),
//   			DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   		},
//   		Custom: &CustomDescriptorProperty{
//   			Data: jsii.String("data"),
//   		},
//   		McpServer: &McpServerDescriptorProperty{
//   			AdditionalData: &McpServerAdditionalDataProperty{
//   				Tools: &McpToolsDescriptorProperty{
//   					Data: jsii.String("data"),
//   					DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   				},
//   			},
//   			Data: jsii.String("data"),
//   			DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   			Source: &DescriptorSourceProperty{
//   				FromUrl: &DescriptorSourceFromUrlProperty{
//   					Url: jsii.String("url"),
//
//   					// the properties below are optional
//   					CredentialProviderConfigurations: []interface{}{
//   						&RegistryRecordCredentialProviderConfigurationProperty{
//   							CredentialProvider: &RegistryRecordCredentialProviderUnionProperty{
//   								IamCredentialProvider: &RegistryRecordIamCredentialProviderProperty{
//   									Region: jsii.String("region"),
//   									RoleArn: jsii.String("roleArn"),
//   									Service: jsii.String("service"),
//   								},
//   								OauthCredentialProvider: &RegistryRecordOAuthCredentialProviderProperty{
//   									ProviderArn: jsii.String("providerArn"),
//
//   									// the properties below are optional
//   									CustomParameters: map[string]*string{
//   										"customParametersKey": jsii.String("customParameters"),
//   									},
//   									GrantType: jsii.String("grantType"),
//   									Scopes: []*string{
//   										jsii.String("scopes"),
//   									},
//   								},
//   							},
//   							CredentialProviderType: jsii.String("credentialProviderType"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RecordType: jsii.String("recordType"),
//   	RegistryId: jsii.String("registryId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	RecordVersion: jsii.String("recordVersion"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html
//
type CfnRegistryRecordProps struct {
	// The typed set of descriptors for a registry record.
	//
	// Exactly one descriptor field is populated based on the record type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-descriptors
	//
	Descriptors interface{} `field:"required" json:"descriptors" yaml:"descriptors"`
	// The name of the registry record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the registry record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-recordtype
	//
	RecordType *string `field:"required" json:"recordType" yaml:"recordType"`
	// The identifier of the registry containing the record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-registryid
	//
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
	// The description of the registry record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The human-readable display name of the registry record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The version of the registry record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-recordversion
	//
	RecordVersion *string `field:"optional" json:"recordVersion" yaml:"recordVersion"`
	// Tags to assign to the registry record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

