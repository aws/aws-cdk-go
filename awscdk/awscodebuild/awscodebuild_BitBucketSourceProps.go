package awscodebuild


// Construction properties for {@link BitBucketSource}.
//
// Example:
//   bbSource := codebuild.source.bitBucket(&bitBucketSourceProps{
//   	owner: jsii.String("owner"),
//   	repo: jsii.String("repo"),
//   })
//
type BitBucketSourceProps struct {
	// The source identifier.
	//
	// This property is required on secondary sources.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// The BitBucket account/user that owns the repo.
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
	BranchOrRef *string `field:"optional" json:"branchOrRef" yaml:"branchOrRef"`
	// This parameter is used for the `name` parameter in the Bitbucket commit status.
	//
	// Can use built-in CodeBuild variables, like $AWS_REGION.
	//
	// Example:
	//   "My build #$CODEBUILD_BUILD_NUMBER"
	//
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-env-vars.html
	//
	BuildStatusName *string `field:"optional" json:"buildStatusName" yaml:"buildStatusName"`
	// The URL that the build will report back to the source provider.
	//
	// Can use built-in CodeBuild variables, like $AWS_REGION.
	//
	// Example:
	//   "$CODEBUILD_PUBLIC_BUILD_URL"
	//
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-env-vars.html
	//
	BuildStatusUrl *string `field:"optional" json:"buildStatusUrl" yaml:"buildStatusUrl"`
	// The depth of history to download.
	//
	// Minimum value is 0.
	// If this value is 0, greater than 25, or not provided,
	// then the full history is downloaded with each build of the project.
	CloneDepth *float64 `field:"optional" json:"cloneDepth" yaml:"cloneDepth"`
	// Whether to fetch submodules while cloning git repo.
	FetchSubmodules *bool `field:"optional" json:"fetchSubmodules" yaml:"fetchSubmodules"`
	// Whether to send notifications on your build's start and end.
	ReportBuildStatus *bool `field:"optional" json:"reportBuildStatus" yaml:"reportBuildStatus"`
	// Whether to create a webhook that will trigger a build every time an event happens in the repository.
	Webhook *bool `field:"optional" json:"webhook" yaml:"webhook"`
	// A list of webhook filters that can constraint what events in the repository will trigger a build.
	//
	// A build is triggered if any of the provided filter groups match.
	// Only valid if `webhook` was not provided as false.
	WebhookFilters *[]FilterGroup `field:"optional" json:"webhookFilters" yaml:"webhookFilters"`
	// Trigger a batch build from a webhook instead of a standard one.
	//
	// Enabling this will enable batch builds on the CodeBuild project.
	WebhookTriggersBatchBuild *bool `field:"optional" json:"webhookTriggersBatchBuild" yaml:"webhookTriggersBatchBuild"`
}

