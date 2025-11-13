package interfacesawsec2


// A reference to a VPCPeeringConnection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCPeeringConnectionReference := &VPCPeeringConnectionReference{
//   	VpcPeeringConnectionId: jsii.String("vpcPeeringConnectionId"),
//   }
//
type VPCPeeringConnectionReference struct {
	// The Id of the VPCPeeringConnection resource.
	VpcPeeringConnectionId *string `field:"required" json:"vpcPeeringConnectionId" yaml:"vpcPeeringConnectionId"`
}

