package previewawsdevopsagentmixins


// GitHub repository integration configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnAssociationPropsMixin_GitHubConfigurationProperty struct {
	// Repository owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-githubconfiguration.html#cfn-devopsagent-association-githubconfiguration-owner
	//
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Type of repository owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-githubconfiguration.html#cfn-devopsagent-association-githubconfiguration-ownertype
	//
	OwnerType *string `field:"optional" json:"ownerType" yaml:"ownerType"`
	// Associated Github repo ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-githubconfiguration.html#cfn-devopsagent-association-githubconfiguration-repoid
	//
	RepoId *string `field:"optional" json:"repoId" yaml:"repoId"`
	// Associated Github repo name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-githubconfiguration.html#cfn-devopsagent-association-githubconfiguration-reponame
	//
	RepoName *string `field:"optional" json:"repoName" yaml:"repoName"`
}

