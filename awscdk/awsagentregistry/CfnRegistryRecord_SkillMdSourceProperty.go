package awsagentregistry


// Source configuration for a SkillMd document.
//
// Unlike MCP/A2A sources, SkillMd does not support credential providers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   skillMdSourceProperty := &SkillMdSourceProperty{
//   	FromUrl: &SkillMdSourceFromUrlProperty{
//   		Url: jsii.String("url"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-skillmdsource.html
//
type CfnRegistryRecord_SkillMdSourceProperty struct {
	// URL-based source for SkillMd content (sync is skipped;
	//
	// content is provided inline via Data).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-skillmdsource.html#cfn-agentregistry-registryrecord-skillmdsource-fromurl
	//
	FromUrl interface{} `field:"optional" json:"fromUrl" yaml:"fromUrl"`
}

