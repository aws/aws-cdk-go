package mixinsawsec2


// Properties for CfnVPCDHCPOptionsAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPCDHCPOptionsAssociationMixinProps := &CfnVPCDHCPOptionsAssociationMixinProps{
//   	DhcpOptionsId: jsii.String("dhcpOptionsId"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcdhcpoptionsassociation.html
//
type CfnVPCDHCPOptionsAssociationMixinProps struct {
	// The ID of the DHCP options set, or `default` to associate no DHCP options with the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcdhcpoptionsassociation.html#cfn-ec2-vpcdhcpoptionsassociation-dhcpoptionsid
	//
	DhcpOptionsId *string `field:"optional" json:"dhcpOptionsId" yaml:"dhcpOptionsId"`
	// The ID of the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcdhcpoptionsassociation.html#cfn-ec2-vpcdhcpoptionsassociation-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

