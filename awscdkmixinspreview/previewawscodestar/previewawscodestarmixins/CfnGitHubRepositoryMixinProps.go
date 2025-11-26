package previewawscodestarmixins


// Properties for CfnGitHubRepositoryPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGitHubRepositoryMixinProps := &CfnGitHubRepositoryMixinProps{
//   	Code: &CodeProperty{
//   		S3: &S3Property{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   			ObjectVersion: jsii.String("objectVersion"),
//   		},
//   	},
//   	ConnectionArn: jsii.String("connectionArn"),
//   	EnableIssues: jsii.Boolean(false),
//   	IsPrivate: jsii.Boolean(false),
//   	RepositoryAccessToken: jsii.String("repositoryAccessToken"),
//   	RepositoryDescription: jsii.String("repositoryDescription"),
//   	RepositoryName: jsii.String("repositoryName"),
//   	RepositoryOwner: jsii.String("repositoryOwner"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestar-githubrepository.html
//
type CfnGitHubRepositoryMixinProps struct {
	// Information about code to be committed to a repository after it is created in an CloudFormation stack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestar-githubrepository.html#cfn-codestar-githubrepository-code
	//
	Code interface{} `field:"optional" json:"code" yaml:"code"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestar-githubrepository.html#cfn-codestar-githubrepository-connectionarn
	//
	ConnectionArn *string `field:"optional" json:"connectionArn" yaml:"connectionArn"`
	// Indicates whether to enable issues for the GitHub repository.
	//
	// You can use GitHub issues to track information and bugs for your repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestar-githubrepository.html#cfn-codestar-githubrepository-enableissues
	//
	EnableIssues interface{} `field:"optional" json:"enableIssues" yaml:"enableIssues"`
	// Indicates whether the GitHub repository is a private repository.
	//
	// If so, you choose who can see and commit to this repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestar-githubrepository.html#cfn-codestar-githubrepository-isprivate
	//
	IsPrivate interface{} `field:"optional" json:"isPrivate" yaml:"isPrivate"`
	// The GitHub user's personal access token for the GitHub repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestar-githubrepository.html#cfn-codestar-githubrepository-repositoryaccesstoken
	//
	RepositoryAccessToken *string `field:"optional" json:"repositoryAccessToken" yaml:"repositoryAccessToken"`
	// A comment or description about the new repository.
	//
	// This description is displayed in GitHub after the repository is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestar-githubrepository.html#cfn-codestar-githubrepository-repositorydescription
	//
	RepositoryDescription *string `field:"optional" json:"repositoryDescription" yaml:"repositoryDescription"`
	// The name of the repository you want to create in GitHub with CloudFormation stack creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestar-githubrepository.html#cfn-codestar-githubrepository-repositoryname
	//
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// The GitHub user name for the owner of the GitHub repository to be created.
	//
	// If this repository should be owned by a GitHub organization, provide its name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestar-githubrepository.html#cfn-codestar-githubrepository-repositoryowner
	//
	RepositoryOwner *string `field:"optional" json:"repositoryOwner" yaml:"repositoryOwner"`
}

