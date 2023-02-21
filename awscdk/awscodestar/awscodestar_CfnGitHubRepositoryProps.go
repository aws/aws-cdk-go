package awscodestar


// Properties for defining a `CfnGitHubRepository`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGitHubRepositoryProps := &cfnGitHubRepositoryProps{
//   	repositoryName: jsii.String("repositoryName"),
//   	repositoryOwner: jsii.String("repositoryOwner"),
//
//   	// the properties below are optional
//   	code: &codeProperty{
//   		s3: &s3Property{
//   			bucket: jsii.String("bucket"),
//   			key: jsii.String("key"),
//
//   			// the properties below are optional
//   			objectVersion: jsii.String("objectVersion"),
//   		},
//   	},
//   	connectionArn: jsii.String("connectionArn"),
//   	enableIssues: jsii.Boolean(false),
//   	isPrivate: jsii.Boolean(false),
//   	repositoryAccessToken: jsii.String("repositoryAccessToken"),
//   	repositoryDescription: jsii.String("repositoryDescription"),
//   }
//
type CfnGitHubRepositoryProps struct {
	// The name of the repository you want to create in GitHub with AWS CloudFormation stack creation.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// The GitHub user name for the owner of the GitHub repository to be created.
	//
	// If this repository should be owned by a GitHub organization, provide its name.
	RepositoryOwner *string `field:"required" json:"repositoryOwner" yaml:"repositoryOwner"`
	// Information about code to be committed to a repository after it is created in an AWS CloudFormation stack.
	Code interface{} `field:"optional" json:"code" yaml:"code"`
	// `AWS::CodeStar::GitHubRepository.ConnectionArn`.
	ConnectionArn *string `field:"optional" json:"connectionArn" yaml:"connectionArn"`
	// Indicates whether to enable issues for the GitHub repository.
	//
	// You can use GitHub issues to track information and bugs for your repository.
	EnableIssues interface{} `field:"optional" json:"enableIssues" yaml:"enableIssues"`
	// Indicates whether the GitHub repository is a private repository.
	//
	// If so, you choose who can see and commit to this repository.
	IsPrivate interface{} `field:"optional" json:"isPrivate" yaml:"isPrivate"`
	// The GitHub user's personal access token for the GitHub repository.
	RepositoryAccessToken *string `field:"optional" json:"repositoryAccessToken" yaml:"repositoryAccessToken"`
	// A comment or description about the new repository.
	//
	// This description is displayed in GitHub after the repository is created.
	RepositoryDescription *string `field:"optional" json:"repositoryDescription" yaml:"repositoryDescription"`
}

