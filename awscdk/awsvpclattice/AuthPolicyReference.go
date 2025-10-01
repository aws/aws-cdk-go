package awsvpclattice


// A reference to a AuthPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authPolicyReference := &AuthPolicyReference{
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   }
//
type AuthPolicyReference struct {
	// The ResourceIdentifier of the AuthPolicy resource.
	ResourceIdentifier *string `field:"required" json:"resourceIdentifier" yaml:"resourceIdentifier"`
}

