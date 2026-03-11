package awsec2


// Properties for CfnSubnetNetworkAclAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSubnetNetworkAclAssociationMixinProps := &CfnSubnetNetworkAclAssociationMixinProps{
//   	NetworkAclId: jsii.String("networkAclId"),
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetnetworkaclassociation.html
//
type CfnSubnetNetworkAclAssociationMixinProps struct {
	// The ID of the network ACL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetnetworkaclassociation.html#cfn-ec2-subnetnetworkaclassociation-networkaclid
	//
	NetworkAclId interface{} `field:"optional" json:"networkAclId" yaml:"networkAclId"`
	// The ID of the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetnetworkaclassociation.html#cfn-ec2-subnetnetworkaclassociation-subnetid
	//
	SubnetId interface{} `field:"optional" json:"subnetId" yaml:"subnetId"`
}

