package awscodecommit


// A reference to a Repository resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryReference := &RepositoryReference{
//   	RepositoryArn: jsii.String("repositoryArn"),
//   	RepositoryId: jsii.String("repositoryId"),
//   }
//
type RepositoryReference struct {
	// The ARN of the Repository resource.
	RepositoryArn *string `field:"required" json:"repositoryArn" yaml:"repositoryArn"`
	// The Id of the Repository resource.
	RepositoryId *string `field:"required" json:"repositoryId" yaml:"repositoryId"`
}

