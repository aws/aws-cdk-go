package awsecr


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryAttributes := &RepositoryAttributes{
//   	RepositoryArn: jsii.String("repositoryArn"),
//   	RepositoryName: jsii.String("repositoryName"),
//   }
//
type RepositoryAttributes struct {
	RepositoryArn *string `field:"required" json:"repositoryArn" yaml:"repositoryArn"`
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
}

