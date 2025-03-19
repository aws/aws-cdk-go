package awscdk


// The location of the published docker image.
//
// This is where the image can be
// consumed at runtime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerImageAssetLocation := &DockerImageAssetLocation{
//   	ImageUri: jsii.String("imageUri"),
//   	RepositoryName: jsii.String("repositoryName"),
//
//   	// the properties below are optional
//   	ImageTag: jsii.String("imageTag"),
//   }
//
type DockerImageAssetLocation struct {
	// The URI of the image in Amazon ECR (including a tag).
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// The name of the ECR repository.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// The tag of the image in Amazon ECR.
	// Default: - the hash of the asset, or the `dockerTagPrefix` concatenated with the asset hash if a `dockerTagPrefix` is specified in the stack synthesizer.
	//
	ImageTag *string `field:"optional" json:"imageTag" yaml:"imageTag"`
}

