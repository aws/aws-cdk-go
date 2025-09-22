package awsamplify


// A reference to a Branch resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   branchReference := &BranchReference{
//   	BranchArn: jsii.String("branchArn"),
//   }
//
type BranchReference struct {
	// The Arn of the Branch resource.
	BranchArn *string `field:"required" json:"branchArn" yaml:"branchArn"`
}

