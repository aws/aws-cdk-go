package awsvpclattice


// A reference to a ResourcePolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcePolicyReference := &ResourcePolicyReference{
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
type ResourcePolicyReference struct {
	// The ResourceArn of the ResourcePolicy resource.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

