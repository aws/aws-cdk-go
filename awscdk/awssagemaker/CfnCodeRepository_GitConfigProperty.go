package awssagemaker


// Specifies configuration details for a Git repository in your AWS account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitConfigProperty := &GitConfigProperty{
//   	RepositoryUrl: jsii.String("repositoryUrl"),
//
//   	// the properties below are optional
//   	Branch: jsii.String("branch"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-coderepository-gitconfig.html
//
type CfnCodeRepository_GitConfigProperty struct {
	// The URL where the Git repository is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-coderepository-gitconfig.html#cfn-sagemaker-coderepository-gitconfig-repositoryurl
	//
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// The default branch for the Git repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-coderepository-gitconfig.html#cfn-sagemaker-coderepository-gitconfig-branch
	//
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the credentials used to access the git repository.
	//
	// The secret must have a staging label of `AWSCURRENT` and must be in the following format:
	//
	// `{"username": *UserName* , "password": *Password* }`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-coderepository-gitconfig.html#cfn-sagemaker-coderepository-gitconfig-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

