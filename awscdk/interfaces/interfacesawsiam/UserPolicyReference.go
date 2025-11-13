package interfacesawsiam


// A reference to a UserPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPolicyReference := &UserPolicyReference{
//   	PolicyName: jsii.String("policyName"),
//   	UserName: jsii.String("userName"),
//   }
//
type UserPolicyReference struct {
	// The PolicyName of the UserPolicy resource.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The UserName of the UserPolicy resource.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
}

