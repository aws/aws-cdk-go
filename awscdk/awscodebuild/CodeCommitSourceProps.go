package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodecommit"
)

// Construction properties for `CodeCommitSource`.
//
// Example:
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"
//   var repo repository
//   var bucket bucket
//
//
//   project := codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
//   	SecondarySources: []iSource{
//   		codebuild.Source_CodeCommit(&CodeCommitSourceProps{
//   			Identifier: jsii.String("source2"),
//   			Repository: repo,
//   		}),
//   	},
//   	SecondaryArtifacts: []iArtifacts{
//   		codebuild.Artifacts_S3(&S3ArtifactsProps{
//   			Identifier: jsii.String("artifact2"),
//   			Bucket: bucket,
//   			Path: jsii.String("some/path"),
//   			Name: jsii.String("file.zip"),
//   		}),
//   	},
//   })
//
type CodeCommitSourceProps struct {
	// The source identifier.
	//
	// This property is required on secondary sources.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	Repository awscodecommit.IRepository `field:"required" json:"repository" yaml:"repository"`
	// The commit ID, pull request ID, branch name, or tag name that corresponds to the version of the source code you want to build.
	//
	// Example:
	//   "mybranch"
	//
	// Default: the default branch's HEAD commit ID is used.
	//
	BranchOrRef *string `field:"optional" json:"branchOrRef" yaml:"branchOrRef"`
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
}

