package awsec2


// Properties for defining a `CfnSubnetNetworkAclAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubnetNetworkAclAssociationProps := &CfnSubnetNetworkAclAssociationProps{
//   	NetworkAclId: jsii.String("networkAclId"),
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetnetworkaclassociation.html
//
type CfnSubnetNetworkAclAssociationProps struct {
	// The ID of the network ACL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetnetworkaclassociation.html#cfn-ec2-subnetnetworkaclassociation-networkaclid
	//
	NetworkAclId *string `field:"required" json:"networkAclId" yaml:"networkAclId"`
	// The ID of the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetnetworkaclassociation.html#cfn-ec2-subnetnetworkaclassociation-subnetid
	//
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

