package awssecurityagent


// The provider-specific input required to create the integration.
//
// Exactly one of the provider input objects must be specified.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   providerInputProperty := &ProviderInputProperty{
//   	Bitbucket: &BitbucketInputProperty{
//   		Code: jsii.String("code"),
//   		InstallationId: jsii.String("installationId"),
//   		State: jsii.String("state"),
//   		Workspace: jsii.String("workspace"),
//   	},
//   	Confluence: &ConfluenceInputProperty{
//   		Code: jsii.String("code"),
//   		InstallationId: jsii.String("installationId"),
//   		SiteUrl: jsii.String("siteUrl"),
//   		State: jsii.String("state"),
//   	},
//   	Github: &GitHubInputProperty{
//   		Code: jsii.String("code"),
//   		State: jsii.String("state"),
//
//   		// the properties below are optional
//   		InstallationId: jsii.String("installationId"),
//   		OrganizationName: jsii.String("organizationName"),
//   		TargetUrl: jsii.String("targetUrl"),
//   	},
//   	Gitlab: &GitLabInputProperty{
//   		AccessToken: jsii.String("accessToken"),
//   		TokenType: jsii.String("tokenType"),
//
//   		// the properties below are optional
//   		GroupId: jsii.String("groupId"),
//   		TargetUrl: jsii.String("targetUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-providerinput.html
//
type CfnIntegration_ProviderInputProperty struct {
	// The input required to create a Bitbucket integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-providerinput.html#cfn-securityagent-integration-providerinput-bitbucket
	//
	Bitbucket interface{} `field:"optional" json:"bitbucket" yaml:"bitbucket"`
	// The input required to create a Confluence integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-providerinput.html#cfn-securityagent-integration-providerinput-confluence
	//
	Confluence interface{} `field:"optional" json:"confluence" yaml:"confluence"`
	// The input required to create a GitHub integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-providerinput.html#cfn-securityagent-integration-providerinput-github
	//
	Github interface{} `field:"optional" json:"github" yaml:"github"`
	// The input required to create a GitLab integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-providerinput.html#cfn-securityagent-integration-providerinput-gitlab
	//
	Gitlab interface{} `field:"optional" json:"gitlab" yaml:"gitlab"`
}

