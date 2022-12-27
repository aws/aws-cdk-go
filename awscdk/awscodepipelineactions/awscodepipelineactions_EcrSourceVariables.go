package awscodepipelineactions


// The CodePipeline variables emitted by the ECR source Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecrSourceVariables := &ecrSourceVariables{
//   	imageDigest: jsii.String("imageDigest"),
//   	imageTag: jsii.String("imageTag"),
//   	imageUri: jsii.String("imageUri"),
//   	registryId: jsii.String("registryId"),
//   	repositoryName: jsii.String("repositoryName"),
//   }
//
type EcrSourceVariables struct {
	// The digest of the current image, in the form '<digest type>:<digest value>'.
	ImageDigest *string `field:"required" json:"imageDigest" yaml:"imageDigest"`
	// The Docker tag of the current image.
	ImageTag *string `field:"required" json:"imageTag" yaml:"imageTag"`
	// The full ECR Docker URI of the current image.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// The identifier of the registry.
	//
	// In ECR, this is usually the ID of the AWS account owning it.
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
	// The physical name of the repository that this action tracks.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
}

