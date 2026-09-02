package awsagentregistry

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRegistryRecordPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnRegistryRecordMixinProps := &CfnRegistryRecordMixinProps{
//   	Description: jsii.String("description"),
//   	Descriptors: &DescriptorsProperty{
//   		A2AAgentCard: &A2aAgentCardDescriptorProperty{
//   			Data: jsii.String("data"),
//   			DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   			Source: &DescriptorSourceProperty{
//   				FromUrl: &DescriptorSourceFromUrlProperty{
//   					CredentialProviderConfigurations: []interface{}{
//   						&RegistryRecordCredentialProviderConfigurationProperty{
//   							CredentialProvider: &RegistryRecordCredentialProviderUnionProperty{
//   								IamCredentialProvider: &RegistryRecordIamCredentialProviderProperty{
//   									Region: jsii.String("region"),
//   									RoleArn: jsii.String("roleArn"),
//   									Service: jsii.String("service"),
//   								},
//   								OauthCredentialProvider: &RegistryRecordOAuthCredentialProviderProperty{
//   									CustomParameters: map[string]*string{
//   										"customParametersKey": jsii.String("customParameters"),
//   									},
//   									GrantType: jsii.String("grantType"),
//   									ProviderArn: jsii.String("providerArn"),
//   									Scopes: []*string{
//   										jsii.String("scopes"),
//   									},
//   								},
//   							},
//   							CredentialProviderType: jsii.String("credentialProviderType"),
//   						},
//   					},
//   					Url: jsii.String("url"),
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
//   					CredentialProviderConfigurations: []interface{}{
//   						&RegistryRecordCredentialProviderConfigurationProperty{
//   							CredentialProvider: &RegistryRecordCredentialProviderUnionProperty{
//   								IamCredentialProvider: &RegistryRecordIamCredentialProviderProperty{
//   									Region: jsii.String("region"),
//   									RoleArn: jsii.String("roleArn"),
//   									Service: jsii.String("service"),
//   								},
//   								OauthCredentialProvider: &RegistryRecordOAuthCredentialProviderProperty{
//   									CustomParameters: map[string]*string{
//   										"customParametersKey": jsii.String("customParameters"),
//   									},
//   									GrantType: jsii.String("grantType"),
//   									ProviderArn: jsii.String("providerArn"),
//   									Scopes: []*string{
//   										jsii.String("scopes"),
//   									},
//   								},
//   							},
//   							CredentialProviderType: jsii.String("credentialProviderType"),
//   						},
//   					},
//   					Url: jsii.String("url"),
//   				},
//   			},
//   		},
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	Name: jsii.String("name"),
//   	RecordType: jsii.String("recordType"),
//   	RecordVersion: jsii.String("recordVersion"),
//   	RegistryId: jsii.String("registryId"),
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
type CfnRegistryRecordMixinProps struct {
	// The description of the registry record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The typed set of descriptors for a registry record.
	//
	// Exactly one descriptor field is populated based on the record type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-descriptors
	//
	Descriptors interface{} `field:"optional" json:"descriptors" yaml:"descriptors"`
	// The human-readable display name of the registry record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The name of the registry record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of the registry record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-recordtype
	//
	RecordType *string `field:"optional" json:"recordType" yaml:"recordType"`
	// The version of the registry record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-recordversion
	//
	RecordVersion *string `field:"optional" json:"recordVersion" yaml:"recordVersion"`
	// The identifier of the registry containing the record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-registryid
	//
	RegistryId *string `field:"optional" json:"registryId" yaml:"registryId"`
	// Tags to assign to the registry record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html#cfn-agentregistry-registryrecord-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

