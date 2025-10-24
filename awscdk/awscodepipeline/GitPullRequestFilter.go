package awscodepipeline


// Git pull request filter for trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitPullRequestFilter := &GitPullRequestFilter{
//   	BranchesExcludes: []*string{
//   		jsii.String("branchesExcludes"),
//   	},
//   	BranchesIncludes: []*string{
//   		jsii.String("branchesIncludes"),
//   	},
//   	Events: []GitPullRequestEvent{
//   		awscdk.Aws_codepipeline.GitPullRequestEvent_OPEN,
//   	},
//   	FilePathsExcludes: []*string{
//   		jsii.String("filePathsExcludes"),
//   	},
//   	FilePathsIncludes: []*string{
//   		jsii.String("filePathsIncludes"),
//   	},
//   }
//
type GitPullRequestFilter struct {
	// The list of patterns of Git branches that, when pull request events occurs, are to be excluded from starting the pipeline.
	//
	// You can filter with glob patterns. The `branchesExcludes` takes priority
	// over the `branchesIncludes`.
	//
	// Maximum length of this array is 8.
	// Default: - no branches.
	//
	BranchesExcludes *[]*string `field:"optional" json:"branchesExcludes" yaml:"branchesExcludes"`
	// The list of patterns of Git branches that, when pull request events occurs, are to be included as criteria that starts the pipeline.
	//
	// You can filter with glob patterns. The `branchesExcludes` takes priority
	// over the `branchesIncludes`.
	//
	// Maximum length of this array is 8.
	// Default: - no branches.
	//
	BranchesIncludes *[]*string `field:"optional" json:"branchesIncludes" yaml:"branchesIncludes"`
	// The field that specifies which pull request events to filter on (opened, updated, closed) for the trigger configuration.
	// Default: - all events.
	//
	Events *[]GitPullRequestEvent `field:"optional" json:"events" yaml:"events"`
	// The list of patterns of Git repository file paths that, when pull request events occurs, are to be excluded from starting the pipeline.
	//
	// You can filter with glob patterns. The `filePathsExcludes` takes priority
	// over the `filePathsIncludes`.
	//
	// Maximum length of this array is 8.
	// Default: - no filePaths.
	//
	FilePathsExcludes *[]*string `field:"optional" json:"filePathsExcludes" yaml:"filePathsExcludes"`
	// The list of patterns of Git repository file paths that, when pull request events occurs, are to be included as criteria that starts the pipeline.
	//
	// You can filter with glob patterns. The `filePathsExcludes` takes priority
	// over the `filePathsIncludes`.
	//
	// Maximum length of this array is 8.
	// Default: - no filePaths.
	//
	FilePathsIncludes *[]*string `field:"optional" json:"filePathsIncludes" yaml:"filePathsIncludes"`
}

