package awsecr


// A reference to a PublicRepository resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicRepositoryReference := &PublicRepositoryReference{
//   	PublicRepositoryArn: jsii.String("publicRepositoryArn"),
//   	RepositoryName: jsii.String("repositoryName"),
//   }
//
type PublicRepositoryReference struct {
	// The ARN of the PublicRepository resource.
	PublicRepositoryArn *string `field:"required" json:"publicRepositoryArn" yaml:"publicRepositoryArn"`
	// The RepositoryName of the PublicRepository resource.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
}

