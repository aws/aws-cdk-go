package interfacesawsec2


// A reference to a VPC resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCReference := &VPCReference{
//   	VpcId: jsii.String("vpcId"),
//   }
//
type VPCReference struct {
	// The VpcId of the VPC resource.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

