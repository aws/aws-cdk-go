package interfacesawsopensearchserverless


// A reference to a SecurityPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityPolicyReference := &SecurityPolicyReference{
//   	SecurityPolicyName: jsii.String("securityPolicyName"),
//   	Type: jsii.String("type"),
//   }
//
type SecurityPolicyReference struct {
	// The Name of the SecurityPolicy resource.
	SecurityPolicyName *string `field:"required" json:"securityPolicyName" yaml:"securityPolicyName"`
	// The Type of the SecurityPolicy resource.
	Type *string `field:"required" json:"type" yaml:"type"`
}

