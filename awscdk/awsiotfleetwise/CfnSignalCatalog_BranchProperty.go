package awsiotfleetwise


// A group of signals that are defined in a hierarchical structure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   branchProperty := &BranchProperty{
//   	FullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
type CfnSignalCatalog_BranchProperty struct {
	// The fully qualified name of the branch.
	//
	// For example, the fully qualified name of a branch might be `Vehicle.Body.Engine` .
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// (Optional) A brief description of the branch.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

