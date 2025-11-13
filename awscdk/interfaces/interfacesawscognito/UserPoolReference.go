package interfacesawscognito


// A reference to a UserPool resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolReference := &UserPoolReference{
//   	UserPoolArn: jsii.String("userPoolArn"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type UserPoolReference struct {
	// The ARN of the UserPool resource.
	UserPoolArn *string `field:"required" json:"userPoolArn" yaml:"userPoolArn"`
	// The UserPoolId of the UserPool resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

