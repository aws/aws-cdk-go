package awscodestarconnections


// A reference to a RepositoryLink resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryLinkReference := &RepositoryLinkReference{
//   	RepositoryLinkArn: jsii.String("repositoryLinkArn"),
//   }
//
type RepositoryLinkReference struct {
	// The RepositoryLinkArn of the RepositoryLink resource.
	RepositoryLinkArn *string `field:"required" json:"repositoryLinkArn" yaml:"repositoryLinkArn"`
}

