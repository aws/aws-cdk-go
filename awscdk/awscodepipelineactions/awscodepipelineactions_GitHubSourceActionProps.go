package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

// Construction properties of the {@link GitHubSourceAction GitHub source action}.
//
// Example:
//   var sourceOutput artifact
//   var project pipelineProject
//
//
//   sourceAction := codepipeline_actions.NewGitHubSourceAction(&gitHubSourceActionProps{
//   	actionName: jsii.String("Github_Source"),
//   	output: sourceOutput,
//   	owner: jsii.String("my-owner"),
//   	repo: jsii.String("my-repo"),
//   	oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
//   	variablesNamespace: jsii.String("MyNamespace"),
//   })
//
//   // later:
//
//   // later:
//   codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   	actionName: jsii.String("CodeBuild"),
//   	project: project,
//   	input: sourceOutput,
//   	environmentVariables: map[string]buildEnvironmentVariable{
//   		"COMMIT_URL": &buildEnvironmentVariable{
//   			"value": sourceAction.variables.commitUrl,
//   		},
//   	},
//   })
//
// Experimental.
type GitHubSourceActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// A GitHub OAuth token to use for authentication.
	//
	// It is recommended to use a Secrets Manager `Secret` to obtain the token:
	//
	//    const oauth = cdk.SecretValue.secretsManager('my-github-token');
	//    new GitHubSource(this, 'GitHubAction', { oauthToken: oauth, ... });
	//
	// The GitHub Personal Access Token should have these scopes:
	//
	// * **repo** - to read the repository
	// * **admin:repo_hook** - if you plan to use webhooks (true by default).
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/appendix-github-oauth.html#GitHub-create-personal-token-CLI
	//
	// Experimental.
	OauthToken awscdk.SecretValue `field:"required" json:"oauthToken" yaml:"oauthToken"`
	// Experimental.
	Output awscodepipeline.Artifact `field:"required" json:"output" yaml:"output"`
	// The GitHub account/user that owns the repo.
	// Experimental.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The name of the repo, without the username.
	// Experimental.
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// The branch to use.
	// Experimental.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// How AWS CodePipeline should be triggered.
	//
	// With the default value "WEBHOOK", a webhook is created in GitHub that triggers the action
	// With "POLL", CodePipeline periodically checks the source for changes
	// With "None", the action is not triggered through changes in the source
	//
	// To use `WEBHOOK`, your GitHub Personal Access Token should have
	// **admin:repo_hook** scope (in addition to the regular **repo** scope).
	// Experimental.
	Trigger GitHubTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}

