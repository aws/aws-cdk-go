package interfacesawscognito


// A reference to a UserPoolGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolGroupReference := &UserPoolGroupReference{
//   	GroupName: jsii.String("groupName"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type UserPoolGroupReference struct {
	// The GroupName of the UserPoolGroup resource.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The UserPoolId of the UserPoolGroup resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

