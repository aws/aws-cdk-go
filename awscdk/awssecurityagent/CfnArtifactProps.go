package awssecurityagent


// Properties for defining a `CfnArtifact`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnArtifactProps := &CfnArtifactProps{
//   	AgentSpaceId: jsii.String("agentSpaceId"),
//   	ArtifactType: jsii.String("artifactType"),
//   	FileName: jsii.String("fileName"),
//
//   	// the properties below are optional
//   	ArtifactContent: jsii.String("artifactContent"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-artifact.html
//
type CfnArtifactProps struct {
	// The unique identifier of the agent space to add the artifact to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-artifact.html#cfn-securityagent-artifact-agentspaceid
	//
	AgentSpaceId *string `field:"required" json:"agentSpaceId" yaml:"agentSpaceId"`
	// The file type of the artifact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-artifact.html#cfn-securityagent-artifact-artifacttype
	//
	ArtifactType *string `field:"required" json:"artifactType" yaml:"artifactType"`
	// The file name of the artifact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-artifact.html#cfn-securityagent-artifact-filename
	//
	FileName *string `field:"required" json:"fileName" yaml:"fileName"`
	// The binary content of the artifact to upload, encoded as a Base64 string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-artifact.html#cfn-securityagent-artifact-artifactcontent
	//
	ArtifactContent *string `field:"optional" json:"artifactContent" yaml:"artifactContent"`
}

