package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ipv4CidrBlockAssociationProperty := &Ipv4CidrBlockAssociationProperty{
//   	AssociationId: jsii.String("associationId"),
//   	CidrBlock: jsii.String("cidrBlock"),
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-secondarynetwork-ipv4cidrblockassociation.html
//
type CfnSecondaryNetworkPropsMixin_Ipv4CidrBlockAssociationProperty struct {
	// The association ID for the IPv4 CIDR block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-secondarynetwork-ipv4cidrblockassociation.html#cfn-ec2-secondarynetwork-ipv4cidrblockassociation-associationid
	//
	AssociationId *string `field:"optional" json:"associationId" yaml:"associationId"`
	// The IPv4 CIDR block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-secondarynetwork-ipv4cidrblockassociation.html#cfn-ec2-secondarynetwork-ipv4cidrblockassociation-cidrblock
	//
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// The state of the CIDR block association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-secondarynetwork-ipv4cidrblockassociation.html#cfn-ec2-secondarynetwork-ipv4cidrblockassociation-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

