package awscognito


// A reference to a UserPoolUser resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolUserReference := &UserPoolUserReference{
//   	Username: jsii.String("username"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type UserPoolUserReference struct {
	// The Username of the UserPoolUser resource.
	Username *string `field:"required" json:"username" yaml:"username"`
	// The UserPoolId of the UserPoolUser resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

