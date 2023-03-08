package awsec2


// Properties for defining a `CfnVPCDHCPOptionsAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCDHCPOptionsAssociationProps := &cfnVPCDHCPOptionsAssociationProps{
//   	dhcpOptionsId: jsii.String("dhcpOptionsId"),
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnVPCDHCPOptionsAssociationProps struct {
	// The ID of the DHCP options set, or `default` to associate no DHCP options with the VPC.
	DhcpOptionsId *string `field:"required" json:"dhcpOptionsId" yaml:"dhcpOptionsId"`
	// The ID of the VPC.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

