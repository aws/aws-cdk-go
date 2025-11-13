package interfacesawsiam


// A reference to a GroupPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupPolicyReference := &GroupPolicyReference{
//   	GroupName: jsii.String("groupName"),
//   	PolicyName: jsii.String("policyName"),
//   }
//
type GroupPolicyReference struct {
	// The GroupName of the GroupPolicy resource.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The PolicyName of the GroupPolicy resource.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

