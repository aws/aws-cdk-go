package awscloud9


// The `Repository` property type specifies an AWS CodeCommit source code repository to be cloned into an AWS Cloud9 development environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryProperty := &repositoryProperty{
//   	pathComponent: jsii.String("pathComponent"),
//   	repositoryUrl: jsii.String("repositoryUrl"),
//   }
//
type CfnEnvironmentEC2_RepositoryProperty struct {
	// The path within the development environment's default file system location to clone the AWS CodeCommit repository into.
	//
	// For example, `/REPOSITORY_NAME` would clone the repository into the `/home/USER_NAME/environment/REPOSITORY_NAME` directory in the environment.
	PathComponent *string `field:"required" json:"pathComponent" yaml:"pathComponent"`
	// The clone URL of the AWS CodeCommit repository to be cloned.
	//
	// For example, for an AWS CodeCommit repository this might be `https://git-codecommit.us-east-2.amazonaws.com/v1/repos/REPOSITORY_NAME` .
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
}

