package interfacesawseks


// A reference to a Capability resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capabilityReference := &CapabilityReference{
//   	CapabilityArn: jsii.String("capabilityArn"),
//   }
//
type CapabilityReference struct {
	// The Arn of the Capability resource.
	CapabilityArn *string `field:"required" json:"capabilityArn" yaml:"capabilityArn"`
}

