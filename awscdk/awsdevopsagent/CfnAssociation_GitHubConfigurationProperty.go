package awsdevopsagent


// GitHub repository integration configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitHubConfigurationProperty := &GitHubConfigurationProperty{
//   	Owner: jsii.String("owner"),
//   	OwnerType: jsii.String("ownerType"),
//   	RepoId: jsii.String("repoId"),
//   	RepoName: jsii.String("repoName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-githubconfiguration.html
//
type CfnAssociation_GitHubConfigurationProperty struct {
	// Repository owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-githubconfiguration.html#cfn-devopsagent-association-githubconfiguration-owner
	//
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// Type of repository owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-githubconfiguration.html#cfn-devopsagent-association-githubconfiguration-ownertype
	//
	OwnerType *string `field:"required" json:"ownerType" yaml:"ownerType"`
	// Associated Github repo ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-githubconfiguration.html#cfn-devopsagent-association-githubconfiguration-repoid
	//
	RepoId *string `field:"required" json:"repoId" yaml:"repoId"`
	// Associated Github repo name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-githubconfiguration.html#cfn-devopsagent-association-githubconfiguration-reponame
	//
	RepoName *string `field:"required" json:"repoName" yaml:"repoName"`
}

