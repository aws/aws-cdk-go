package interfacesawsxray


// A reference to a ResourcePolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcePolicyReference := &ResourcePolicyReference{
//   	PolicyName: jsii.String("policyName"),
//   }
//
type ResourcePolicyReference struct {
	// The PolicyName of the ResourcePolicy resource.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

