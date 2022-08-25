package awscodepipelineactions


// The CodePipeline variables emitted by the CodeCommit source Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeCommitSourceVariables := &codeCommitSourceVariables{
//   	authorDate: jsii.String("authorDate"),
//   	branchName: jsii.String("branchName"),
//   	commitId: jsii.String("commitId"),
//   	commitMessage: jsii.String("commitMessage"),
//   	committerDate: jsii.String("committerDate"),
//   	repositoryName: jsii.String("repositoryName"),
//   }
//
type CodeCommitSourceVariables struct {
	// The date the currently last commit on the tracked branch was authored, in ISO-8601 format.
	AuthorDate *string `field:"required" json:"authorDate" yaml:"authorDate"`
	// The name of the branch this action tracks.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// The SHA1 hash of the currently last commit on the tracked branch.
	CommitId *string `field:"required" json:"commitId" yaml:"commitId"`
	// The message of the currently last commit on the tracked branch.
	CommitMessage *string `field:"required" json:"commitMessage" yaml:"commitMessage"`
	// The date the currently last commit on the tracked branch was committed, in ISO-8601 format.
	CommitterDate *string `field:"required" json:"committerDate" yaml:"committerDate"`
	// The name of the repository this action points to.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
}

