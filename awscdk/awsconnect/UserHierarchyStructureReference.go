package awsconnect


// A reference to a UserHierarchyStructure resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userHierarchyStructureReference := &UserHierarchyStructureReference{
//   	UserHierarchyStructureArn: jsii.String("userHierarchyStructureArn"),
//   }
//
type UserHierarchyStructureReference struct {
	// The UserHierarchyStructureArn of the UserHierarchyStructure resource.
	UserHierarchyStructureArn *string `field:"required" json:"userHierarchyStructureArn" yaml:"userHierarchyStructureArn"`
}

