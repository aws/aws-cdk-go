package awsagentregistry


// The agent skills definition descriptor, populated when the record type is SKILL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   agentSkillsDefinitionDescriptorProperty := &AgentSkillsDefinitionDescriptorProperty{
//   	AdditionalData: &AgentSkillsAdditionalDataProperty{
//   		SkillMd: &AgentSkillsMdDescriptorProperty{
//   			Data: jsii.String("data"),
//   			DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   			Source: &SkillMdSourceProperty{
//   				FromUrl: &SkillMdSourceFromUrlProperty{
//   					Url: jsii.String("url"),
//   				},
//   			},
//   		},
//   	},
//   	Data: jsii.String("data"),
//   	DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-agentskillsdefinitiondescriptor.html
//
type CfnRegistryRecordPropsMixin_AgentSkillsDefinitionDescriptorProperty struct {
	// Additional data associated with an agent skills definition descriptor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-agentskillsdefinitiondescriptor.html#cfn-agentregistry-registryrecord-agentskillsdefinitiondescriptor-additionaldata
	//
	AdditionalData interface{} `field:"optional" json:"additionalData" yaml:"additionalData"`
	// Descriptor payload data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-agentskillsdefinitiondescriptor.html#cfn-agentregistry-registryrecord-agentskillsdefinitiondescriptor-data
	//
	Data *string `field:"optional" json:"data" yaml:"data"`
	// Version of the descriptor type schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-agentskillsdefinitiondescriptor.html#cfn-agentregistry-registryrecord-agentskillsdefinitiondescriptor-dataschemaversion
	//
	DataSchemaVersion *string `field:"optional" json:"dataSchemaVersion" yaml:"dataSchemaVersion"`
}

