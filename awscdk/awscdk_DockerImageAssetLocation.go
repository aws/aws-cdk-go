// An experiment to bundle the entire CDK into a single module
package awscdk


// The location of the published docker image.
//
// This is where the image can be
// consumed at runtime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerImageAssetLocation := &DockerImageAssetLocation{
//   	ImageUri: jsii.String("imageUri"),
//   	RepositoryName: jsii.String("repositoryName"),
//   }
//
// Experimental.
type DockerImageAssetLocation struct {
	// The URI of the image in Amazon ECR.
	// Experimental.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// The name of the ECR repository.
	// Experimental.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
}

