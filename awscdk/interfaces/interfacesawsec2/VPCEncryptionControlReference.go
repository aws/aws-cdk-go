package interfacesawsec2


// A reference to a VPCEncryptionControl resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCEncryptionControlReference := &VPCEncryptionControlReference{
//   	VpcEncryptionControlId: jsii.String("vpcEncryptionControlId"),
//   }
//
type VPCEncryptionControlReference struct {
	// The VpcEncryptionControlId of the VPCEncryptionControl resource.
	VpcEncryptionControlId *string `field:"required" json:"vpcEncryptionControlId" yaml:"vpcEncryptionControlId"`
}

