package awscodepipeline


// Git push filter for trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitPushFilter := &GitPushFilter{
//   	BranchesExcludes: []*string{
//   		jsii.String("branchesExcludes"),
//   	},
//   	BranchesIncludes: []*string{
//   		jsii.String("branchesIncludes"),
//   	},
//   	FilePathsExcludes: []*string{
//   		jsii.String("filePathsExcludes"),
//   	},
//   	FilePathsIncludes: []*string{
//   		jsii.String("filePathsIncludes"),
//   	},
//   	TagsExcludes: []*string{
//   		jsii.String("tagsExcludes"),
//   	},
//   	TagsIncludes: []*string{
//   		jsii.String("tagsIncludes"),
//   	},
//   }
//
type GitPushFilter struct {
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
	// The list of patterns of Git tags that, when pushed, are to be excluded from starting the pipeline.
	//
	// You can filter with glob patterns. The `tagsExcludes` takes priority
	// over the `tagsIncludes`.
	//
	// Maximum length of this array is 8.
	// Default: - no tags.
	//
	TagsExcludes *[]*string `field:"optional" json:"tagsExcludes" yaml:"tagsExcludes"`
	// The list of patterns of Git tags that, when pushed, are to be included as criteria that starts the pipeline.
	//
	// You can filter with glob patterns. The `tagsExcludes` takes priority
	// over the `tagsIncludes`.
	//
	// Maximum length of this array is 8.
	// Default: - no tags.
	//
	TagsIncludes *[]*string `field:"optional" json:"tagsIncludes" yaml:"tagsIncludes"`
}

