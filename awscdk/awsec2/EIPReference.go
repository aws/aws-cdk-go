package awsec2


// A reference to a EIP resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eIPReference := &EIPReference{
//   	AllocationId: jsii.String("allocationId"),
//   	PublicIp: jsii.String("publicIp"),
//   }
//
type EIPReference struct {
	// The AllocationId of the EIP resource.
	AllocationId *string `field:"required" json:"allocationId" yaml:"allocationId"`
	// The PublicIp of the EIP resource.
	PublicIp *string `field:"required" json:"publicIp" yaml:"publicIp"`
}

