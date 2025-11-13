package interfacesawssynthetics


// A reference to a Canary resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   canaryReference := &CanaryReference{
//   	CanaryName: jsii.String("canaryName"),
//   }
//
type CanaryReference struct {
	// The Name of the Canary resource.
	CanaryName *string `field:"required" json:"canaryName" yaml:"canaryName"`
}

