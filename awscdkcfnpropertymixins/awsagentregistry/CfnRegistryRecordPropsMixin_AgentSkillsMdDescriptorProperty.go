package awsagentregistry


// Markdown-format descriptor containing an agent skills document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   agentSkillsMdDescriptorProperty := &AgentSkillsMdDescriptorProperty{
//   	Data: jsii.String("data"),
//   	DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   	Source: &SkillMdSourceProperty{
//   		FromUrl: &SkillMdSourceFromUrlProperty{
//   			Url: jsii.String("url"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-agentskillsmddescriptor.html
//
type CfnRegistryRecordPropsMixin_AgentSkillsMdDescriptorProperty struct {
	// Descriptor payload data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-agentskillsmddescriptor.html#cfn-agentregistry-registryrecord-agentskillsmddescriptor-data
	//
	Data *string `field:"optional" json:"data" yaml:"data"`
	// Version of the descriptor type schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-agentskillsmddescriptor.html#cfn-agentregistry-registryrecord-agentskillsmddescriptor-dataschemaversion
	//
	DataSchemaVersion *string `field:"optional" json:"dataSchemaVersion" yaml:"dataSchemaVersion"`
	// Source configuration for a SkillMd document.
	//
	// Unlike MCP/A2A sources, SkillMd does not support credential providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-agentskillsmddescriptor.html#cfn-agentregistry-registryrecord-agentskillsmddescriptor-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

