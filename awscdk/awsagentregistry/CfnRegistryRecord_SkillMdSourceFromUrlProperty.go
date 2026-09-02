package awsagentregistry


// URL-based source for SkillMd content (sync is skipped;
//
// content is provided inline via Data).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   skillMdSourceFromUrlProperty := &SkillMdSourceFromUrlProperty{
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-skillmdsourcefromurl.html
//
type CfnRegistryRecord_SkillMdSourceFromUrlProperty struct {
	// URL source for the SkillMd document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-skillmdsourcefromurl.html#cfn-agentregistry-registryrecord-skillmdsourcefromurl-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
}

