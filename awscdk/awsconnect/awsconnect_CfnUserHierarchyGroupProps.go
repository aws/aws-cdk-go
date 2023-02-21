package awsconnect


// Properties for defining a `CfnUserHierarchyGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserHierarchyGroupProps := &cfnUserHierarchyGroupProps{
//   	instanceArn: jsii.String("instanceArn"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	parentGroupArn: jsii.String("parentGroupArn"),
//   }
//
type CfnUserHierarchyGroupProps struct {
	// The Amazon Resource Name (ARN) of the user hierarchy group.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the user hierarchy group.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the parent group.
	ParentGroupArn *string `field:"optional" json:"parentGroupArn" yaml:"parentGroupArn"`
}

