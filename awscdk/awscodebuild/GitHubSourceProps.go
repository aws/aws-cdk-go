package awscodebuild


// Construction properties for `GitHubSource` and `GitHubEnterpriseSource`.
//
// Example:
//   project := codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
//   	BuildSpec: codebuild.BuildSpec_FromSourceFilename(jsii.String("my-buildspec.yml")),
//   	Source: codebuild.Source_GitHub(&GitHubSourceProps{
//   		Owner: jsii.String("awslabs"),
//   		Repo: jsii.String("aws-cdk"),
//   	}),
//   })
//
type GitHubSourceProps struct {
	// The source identifier.
	//
	// This property is required on secondary sources.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// The GitHub account/user that owns the repo.
	//
	// Example:
	//   "awslabs"
	//
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The name of the repo (without the username).
	//
	// Example:
	//   "aws-cdk"
	//
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// The commit ID, pull request ID, branch name, or tag name that corresponds to the version of the source code you want to build.
	//
	// Example:
	//   "mybranch"
	//
	// Default: the default branch's HEAD commit ID is used.
	//
	BranchOrRef *string `field:"optional" json:"branchOrRef" yaml:"branchOrRef"`
	// This parameter is used for the `context` parameter in the GitHub commit status.
	//
	// Can use built-in CodeBuild variables, like $AWS_REGION.
	//
	// Example:
	//   "My build #$CODEBUILD_BUILD_NUMBER"
	//
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-env-vars.html
	//
	// Default: "AWS CodeBuild $AWS_REGION ($PROJECT_NAME)".
	//
	BuildStatusContext *string `field:"optional" json:"buildStatusContext" yaml:"buildStatusContext"`
	// The URL that the build will report back to the source provider.
	//
	// Can use built-in CodeBuild variables, like $AWS_REGION.
	//
	// Example:
	//   "$CODEBUILD_PUBLIC_BUILD_URL"
	//
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-env-vars.html
	//
	// Default: - link to the AWS Console for CodeBuild to a particular build execution.
	//
	BuildStatusUrl *string `field:"optional" json:"buildStatusUrl" yaml:"buildStatusUrl"`
	// The depth of history to download.
	//
	// Minimum value is 0.
	// If this value is 0, greater than 25, or not provided,
	// then the full history is downloaded with each build of the project.
	CloneDepth *float64 `field:"optional" json:"cloneDepth" yaml:"cloneDepth"`
	// Whether to fetch submodules while cloning git repo.
	// Default: false.
	//
	FetchSubmodules *bool `field:"optional" json:"fetchSubmodules" yaml:"fetchSubmodules"`
	// Whether to send notifications on your build's start and end.
	// Default: true.
	//
	ReportBuildStatus *bool `field:"optional" json:"reportBuildStatus" yaml:"reportBuildStatus"`
	// Whether to create a webhook that will trigger a build every time an event happens in the repository.
	// Default: true if any `webhookFilters` were provided, false otherwise.
	//
	Webhook *bool `field:"optional" json:"webhook" yaml:"webhook"`
	// A list of webhook filters that can constraint what events in the repository will trigger a build.
	//
	// A build is triggered if any of the provided filter groups match.
	// Only valid if `webhook` was not provided as false.
	// Default: every push and every Pull Request (create or update) triggers a build.
	//
	WebhookFilters *[]FilterGroup `field:"optional" json:"webhookFilters" yaml:"webhookFilters"`
	// Trigger a batch build from a webhook instead of a standard one.
	//
	// Enabling this will enable batch builds on the CodeBuild project.
	// Default: false.
	//
	WebhookTriggersBatchBuild *bool `field:"optional" json:"webhookTriggersBatchBuild" yaml:"webhookTriggersBatchBuild"`
}

