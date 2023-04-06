package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
)

// Construction properties of the `GitHubSourceAction GitHub source action`.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // Read the secret from Secrets Manager
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewGitHubSourceAction(&GitHubSourceActionProps{
//   	ActionName: jsii.String("GitHub_Source"),
//   	Owner: jsii.String("awslabs"),
//   	Repo: jsii.String("aws-cdk"),
//   	OauthToken: awscdk.SecretValue_SecretsManager(jsii.String("my-github-token")),
//   	Output: sourceOutput,
//   	Branch: jsii.String("develop"),
//   })
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Source"),
//   	Actions: []iAction{
//   		sourceAction,
//   	},
//   })
//
type GitHubSourceActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// A GitHub OAuth token to use for authentication.
	//
	// It is recommended to use a Secrets Manager `Secret` to obtain the token:
	//
	//   const oauth = cdk.SecretValue.secretsManager('my-github-token');
	//   new GitHubSourceAction(this, 'GitHubAction', { oauthToken: oauth, ... });
	//
	// If you rotate the value in the Secret, you must also change at least one property
	// of the CodePipeline to force CloudFormation to re-read the secret.
	//
	// The GitHub Personal Access Token should have these scopes:
	//
	// * **repo** - to read the repository
	// * **admin:repo_hook** - if you plan to use webhooks (true by default).
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/appendix-github-oauth.html#GitHub-create-personal-token-CLI
	//
	OauthToken awscdk.SecretValue `field:"required" json:"oauthToken" yaml:"oauthToken"`
	Output awscodepipeline.Artifact `field:"required" json:"output" yaml:"output"`
	// The GitHub account/user that owns the repo.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The name of the repo, without the username.
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// The branch to use.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// How AWS CodePipeline should be triggered.
	//
	// With the default value "WEBHOOK", a webhook is created in GitHub that triggers the action
	// With "POLL", CodePipeline periodically checks the source for changes
	// With "None", the action is not triggered through changes in the source
	//
	// To use `WEBHOOK`, your GitHub Personal Access Token should have
	// **admin:repo_hook** scope (in addition to the regular **repo** scope).
	Trigger GitHubTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}

