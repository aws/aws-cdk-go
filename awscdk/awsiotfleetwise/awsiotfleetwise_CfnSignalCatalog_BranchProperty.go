package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   branchProperty := &branchProperty{
//   	fullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnSignalCatalog_BranchProperty struct {
	// `CfnSignalCatalog.BranchProperty.FullyQualifiedName`.
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// `CfnSignalCatalog.BranchProperty.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

