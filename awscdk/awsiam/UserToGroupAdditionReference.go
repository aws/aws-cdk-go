package awsiam


// A reference to a UserToGroupAddition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userToGroupAdditionReference := &UserToGroupAdditionReference{
//   	UserToGroupAdditionId: jsii.String("userToGroupAdditionId"),
//   }
//
type UserToGroupAdditionReference struct {
	// The Id of the UserToGroupAddition resource.
	UserToGroupAdditionId *string `field:"required" json:"userToGroupAdditionId" yaml:"userToGroupAdditionId"`
}

