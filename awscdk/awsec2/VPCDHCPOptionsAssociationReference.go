package awsec2


// A reference to a VPCDHCPOptionsAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCDHCPOptionsAssociationReference := &VPCDHCPOptionsAssociationReference{
//   	DhcpOptionsId: jsii.String("dhcpOptionsId"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
type VPCDHCPOptionsAssociationReference struct {
	// The DhcpOptionsId of the VPCDHCPOptionsAssociation resource.
	DhcpOptionsId *string `field:"required" json:"dhcpOptionsId" yaml:"dhcpOptionsId"`
	// The VpcId of the VPCDHCPOptionsAssociation resource.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

