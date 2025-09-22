package awsmemorydb


// A reference to a User resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userReference := &UserReference{
//   	UserArn: jsii.String("userArn"),
//   	UserName: jsii.String("userName"),
//   }
//
type UserReference struct {
	// The ARN of the User resource.
	UserArn *string `field:"required" json:"userArn" yaml:"userArn"`
	// The UserName of the User resource.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
}

