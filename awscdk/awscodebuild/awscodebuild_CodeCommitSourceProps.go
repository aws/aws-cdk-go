package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodecommit"
)

// Construction properties for {@link CodeCommitSource}.
//
// Example:
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"
//   var repo repository
//   var bucket bucket
//
//
//   project := codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
//   	secondarySources: []iSource{
//   		codebuild.source.codeCommit(&codeCommitSourceProps{
//   			identifier: jsii.String("source2"),
//   			repository: repo,
//   		}),
//   	},
//   	secondaryArtifacts: []iArtifacts{
//   		codebuild.artifacts.s3(&s3ArtifactsProps{
//   			identifier: jsii.String("artifact2"),
//   			bucket: bucket,
//   			path: jsii.String("some/path"),
//   			name: jsii.String("file.zip"),
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
	BranchOrRef *string `field:"optional" json:"branchOrRef" yaml:"branchOrRef"`
	// The depth of history to download.
	//
	// Minimum value is 0.
	// If this value is 0, greater than 25, or not provided,
	// then the full history is downloaded with each build of the project.
	CloneDepth *float64 `field:"optional" json:"cloneDepth" yaml:"cloneDepth"`
	// Whether to fetch submodules while cloning git repo.
	FetchSubmodules *bool `field:"optional" json:"fetchSubmodules" yaml:"fetchSubmodules"`
}

