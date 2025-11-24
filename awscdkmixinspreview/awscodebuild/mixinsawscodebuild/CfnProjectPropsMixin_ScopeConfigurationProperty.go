package mixinsawscodebuild


// Contains configuration information about the scope for a webhook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scopeConfigurationProperty := &ScopeConfigurationProperty{
//   	Domain: jsii.String("domain"),
//   	Name: jsii.String("name"),
//   	Scope: jsii.String("scope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-scopeconfiguration.html
//
type CfnProjectPropsMixin_ScopeConfigurationProperty struct {
	// The domain of the GitHub Enterprise organization or the GitLab Self Managed group.
	//
	// Note that this parameter is only required if your project's source type is GITHUB_ENTERPRISE or GITLAB_SELF_MANAGED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-scopeconfiguration.html#cfn-codebuild-project-scopeconfiguration-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The name of either the enterprise or organization that will send webhook events to CodeBuild , depending on if the webhook is a global or organization webhook respectively.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-scopeconfiguration.html#cfn-codebuild-project-scopeconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of scope for a GitHub or GitLab webhook.
	//
	// The scope default is GITHUB_ORGANIZATION.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-scopeconfiguration.html#cfn-codebuild-project-scopeconfiguration-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

