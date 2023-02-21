package awscodepipelineactions


// The CodePipeline variables emitted by CodeStar source Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeStarSourceVariables := &codeStarSourceVariables{
//   	authorDate: jsii.String("authorDate"),
//   	branchName: jsii.String("branchName"),
//   	commitId: jsii.String("commitId"),
//   	commitMessage: jsii.String("commitMessage"),
//   	connectionArn: jsii.String("connectionArn"),
//   	fullRepositoryName: jsii.String("fullRepositoryName"),
//   }
//
// Experimental.
type CodeStarSourceVariables struct {
	// The date the currently last commit on the tracked branch was authored, in ISO-8601 format.
	// Experimental.
	AuthorDate *string `field:"required" json:"authorDate" yaml:"authorDate"`
	// The name of the branch this action tracks.
	// Experimental.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// The SHA1 hash of the currently last commit on the tracked branch.
	// Experimental.
	CommitId *string `field:"required" json:"commitId" yaml:"commitId"`
	// The message of the currently last commit on the tracked branch.
	// Experimental.
	CommitMessage *string `field:"required" json:"commitMessage" yaml:"commitMessage"`
	// The connection ARN this source uses.
	// Experimental.
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// The name of the repository this action points to.
	// Experimental.
	FullRepositoryName *string `field:"required" json:"fullRepositoryName" yaml:"fullRepositoryName"`
}

