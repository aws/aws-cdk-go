package awssecurityagent


// Properties for CfnArtifactPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnArtifactMixinProps := &CfnArtifactMixinProps{
//   	AgentSpaceId: jsii.String("agentSpaceId"),
//   	ArtifactContent: jsii.String("artifactContent"),
//   	ArtifactType: jsii.String("artifactType"),
//   	FileName: jsii.String("fileName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-artifact.html
//
type CfnArtifactMixinProps struct {
	// The unique identifier of the agent space to add the artifact to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-artifact.html#cfn-securityagent-artifact-agentspaceid
	//
	AgentSpaceId *string `field:"optional" json:"agentSpaceId" yaml:"agentSpaceId"`
	// The binary content of the artifact to upload, encoded as a Base64 string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-artifact.html#cfn-securityagent-artifact-artifactcontent
	//
	ArtifactContent *string `field:"optional" json:"artifactContent" yaml:"artifactContent"`
	// The file type of the artifact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-artifact.html#cfn-securityagent-artifact-artifacttype
	//
	ArtifactType *string `field:"optional" json:"artifactType" yaml:"artifactType"`
	// The file name of the artifact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-artifact.html#cfn-securityagent-artifact-filename
	//
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
}

