package interfacesawselasticache


// A reference to a UserGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userGroupReference := &UserGroupReference{
//   	UserGroupArn: jsii.String("userGroupArn"),
//   	UserGroupId: jsii.String("userGroupId"),
//   }
//
type UserGroupReference struct {
	// The ARN of the UserGroup resource.
	UserGroupArn *string `field:"required" json:"userGroupArn" yaml:"userGroupArn"`
	// The UserGroupId of the UserGroup resource.
	UserGroupId *string `field:"required" json:"userGroupId" yaml:"userGroupId"`
}

