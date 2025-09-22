package awsb2bi


// A reference to a Capability resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capabilityReference := &CapabilityReference{
//   	CapabilityArn: jsii.String("capabilityArn"),
//   	CapabilityId: jsii.String("capabilityId"),
//   }
//
type CapabilityReference struct {
	// The ARN of the Capability resource.
	CapabilityArn *string `field:"required" json:"capabilityArn" yaml:"capabilityArn"`
	// The CapabilityId of the Capability resource.
	CapabilityId *string `field:"required" json:"capabilityId" yaml:"capabilityId"`
}

