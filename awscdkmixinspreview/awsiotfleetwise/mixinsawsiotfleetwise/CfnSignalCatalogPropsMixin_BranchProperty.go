package mixinsawsiotfleetwise


// A group of signals that are defined in a hierarchical structure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   branchProperty := &BranchProperty{
//   	Description: jsii.String("description"),
//   	FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-branch.html
//
type CfnSignalCatalogPropsMixin_BranchProperty struct {
	// A brief description of the branch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-branch.html#cfn-iotfleetwise-signalcatalog-branch-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The fully qualified name of the branch.
	//
	// For example, the fully qualified name of a branch might be `Vehicle.Body.Engine` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-branch.html#cfn-iotfleetwise-signalcatalog-branch-fullyqualifiedname
	//
	FullyQualifiedName *string `field:"optional" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
}

