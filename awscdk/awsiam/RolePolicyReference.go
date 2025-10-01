package awsiam


// A reference to a RolePolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rolePolicyReference := &RolePolicyReference{
//   	PolicyName: jsii.String("policyName"),
//   	RoleName: jsii.String("roleName"),
//   }
//
type RolePolicyReference struct {
	// The PolicyName of the RolePolicy resource.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The RoleName of the RolePolicy resource.
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
}

