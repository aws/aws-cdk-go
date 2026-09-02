package awsagentregistry


// Additional data associated with an agent skills definition descriptor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentSkillsAdditionalDataProperty := &AgentSkillsAdditionalDataProperty{
//   	SkillMd: &AgentSkillsMdDescriptorProperty{
//   		Data: jsii.String("data"),
//   		DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   		Source: &SkillMdSourceProperty{
//   			FromUrl: &SkillMdSourceFromUrlProperty{
//   				Url: jsii.String("url"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-agentskillsadditionaldata.html
//
type CfnRegistryRecord_AgentSkillsAdditionalDataProperty struct {
	// Markdown-format descriptor containing an agent skills document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-agentskillsadditionaldata.html#cfn-agentregistry-registryrecord-agentskillsadditionaldata-skillmd
	//
	SkillMd interface{} `field:"optional" json:"skillMd" yaml:"skillMd"`
}

