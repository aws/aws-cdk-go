package awscodepipelineactions


// The CodePipeline variables emitted by the ECR build and publish Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecrBuildAndPublishVariables := &EcrBuildAndPublishVariables{
//   	EcrImageDigestId: jsii.String("ecrImageDigestId"),
//   	EcrRepositoryName: jsii.String("ecrRepositoryName"),
//   }
//
type EcrBuildAndPublishVariables struct {
	// The sha256 digest of the image manifest.
	EcrImageDigestId *string `field:"required" json:"ecrImageDigestId" yaml:"ecrImageDigestId"`
	// The name of the Amazon ECR repository where the image was pushed.
	EcrRepositoryName *string `field:"required" json:"ecrRepositoryName" yaml:"ecrRepositoryName"`
}

