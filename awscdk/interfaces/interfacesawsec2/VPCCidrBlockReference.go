package interfacesawsec2


// A reference to a VPCCidrBlock resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCCidrBlockReference := &VPCCidrBlockReference{
//   	VpcCidrBlockId: jsii.String("vpcCidrBlockId"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
type VPCCidrBlockReference struct {
	// The Id of the VPCCidrBlock resource.
	VpcCidrBlockId *string `field:"required" json:"vpcCidrBlockId" yaml:"vpcCidrBlockId"`
	// The VpcId of the VPCCidrBlock resource.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

