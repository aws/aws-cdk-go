package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipelineactions"
)

// Options for GitHub sources.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   pipelines.codePipelineSource.gitHub(jsii.String("org/repo"), jsii.String("branch"), &gitHubSourceOptions{
//   	// This is optional
//   	authentication: cdk.secretValue.secretsManager(jsii.String("my-token")),
//   })
//
type GitHubSourceOptions struct {
	// The action name used for this source in the CodePipeline.
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// A GitHub OAuth token to use for authentication.
	//
	// It is recommended to use a Secrets Manager `Secret` to obtain the token:
	//
	// ```ts
	// const oauth = cdk.SecretValue.secretsManager('my-github-token');
	// ```
	//
	// The GitHub Personal Access Token should have these scopes:
	//
	// * **repo** - to read the repository
	// * **admin:repo_hook** - if you plan to use webhooks (true by default).
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/GitHub-create-personal-token-CLI.html
	//
	Authentication awscdk.SecretValue `field:"optional" json:"authentication" yaml:"authentication"`
	// How AWS CodePipeline should be triggered.
	//
	// With the default value "WEBHOOK", a webhook is created in GitHub that triggers the action.
	// With "POLL", CodePipeline periodically checks the source for changes.
	// With "None", the action is not triggered through changes in the source.
	//
	// To use `WEBHOOK`, your GitHub Personal Access Token should have
	// **admin:repo_hook** scope (in addition to the regular **repo** scope).
	Trigger awscodepipelineactions.GitHubTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}

