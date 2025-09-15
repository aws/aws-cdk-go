package awsconnect


// A reference to a UserHierarchyGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userHierarchyGroupReference := &UserHierarchyGroupReference{
//   	UserHierarchyGroupArn: jsii.String("userHierarchyGroupArn"),
//   }
//
type UserHierarchyGroupReference struct {
	// The UserHierarchyGroupArn of the UserHierarchyGroup resource.
	UserHierarchyGroupArn *string `field:"required" json:"userHierarchyGroupArn" yaml:"userHierarchyGroupArn"`
}

