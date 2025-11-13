package interfacesawsorganizations


// A reference to a ResourcePolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcePolicyReference := &ResourcePolicyReference{
//   	ResourcePolicyArn: jsii.String("resourcePolicyArn"),
//   	ResourcePolicyId: jsii.String("resourcePolicyId"),
//   }
//
type ResourcePolicyReference struct {
	// The ARN of the ResourcePolicy resource.
	ResourcePolicyArn *string `field:"required" json:"resourcePolicyArn" yaml:"resourcePolicyArn"`
	// The Id of the ResourcePolicy resource.
	ResourcePolicyId *string `field:"required" json:"resourcePolicyId" yaml:"resourcePolicyId"`
}

