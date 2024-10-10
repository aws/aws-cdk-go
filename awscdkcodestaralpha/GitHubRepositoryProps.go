package awscdkcodestaralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Construction properties of `GitHubRepository`.
//
// Example:
//   import codestar "github.com/aws/aws-cdk-go/awscdkcodestaralpha"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   codestar.NewGitHubRepository(this, jsii.String("GitHubRepo"), &GitHubRepositoryProps{
//   	Owner: jsii.String("aws"),
//   	RepositoryName: jsii.String("aws-cdk"),
//   	AccessToken: awscdk.SecretValue_SecretsManager(jsii.String("my-github-token"), &SecretsManagerSecretOptions{
//   		JsonField: jsii.String("token"),
//   	}),
//   	ContentsBucket: s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("amzn-s3-demo-bucket")),
//   	ContentsKey: jsii.String("import.zip"),
//   })
//
// Experimental.
type GitHubRepositoryProps struct {
	// The GitHub user's personal access token for the GitHub repository.
	// Experimental.
	AccessToken awscdk.SecretValue `field:"required" json:"accessToken" yaml:"accessToken"`
	// The name of the Amazon S3 bucket that contains the ZIP file with the content to be committed to the new repository.
	// Experimental.
	ContentsBucket awss3.IBucket `field:"required" json:"contentsBucket" yaml:"contentsBucket"`
	// The S3 object key or file name for the ZIP file.
	// Experimental.
	ContentsKey *string `field:"required" json:"contentsKey" yaml:"contentsKey"`
	// The GitHub user name for the owner of the GitHub repository to be created.
	//
	// If this
	// repository should be owned by a GitHub organization, provide its name.
	// Experimental.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The name of the repository you want to create in GitHub with AWS CloudFormation stack creation.
	// Experimental.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// The object version of the ZIP file, if versioning is enabled for the Amazon S3 bucket.
	// Default: - not specified.
	//
	// Experimental.
	ContentsS3Version *string `field:"optional" json:"contentsS3Version" yaml:"contentsS3Version"`
	// A comment or description about the new repository.
	//
	// This description is displayed in GitHub after the repository
	// is created.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether to enable issues for the GitHub repository.
	//
	// You can use GitHub issues to track information
	// and bugs for your repository.
	// Default: true.
	//
	// Experimental.
	EnableIssues *bool `field:"optional" json:"enableIssues" yaml:"enableIssues"`
	// Indicates whether the GitHub repository is a private repository.
	//
	// If so, you choose who can see and commit to
	// this repository.
	// Default: RepositoryVisibility.PUBLIC
	//
	// Experimental.
	Visibility RepositoryVisibility `field:"optional" json:"visibility" yaml:"visibility"`
}

