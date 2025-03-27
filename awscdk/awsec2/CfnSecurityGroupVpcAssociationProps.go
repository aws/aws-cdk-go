package awsec2


// Properties for defining a `CfnSecurityGroupVpcAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityGroupVpcAssociationProps := &CfnSecurityGroupVpcAssociationProps{
//   	GroupId: jsii.String("groupId"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupvpcassociation.html
//
type CfnSecurityGroupVpcAssociationProps struct {
	// The association's security group ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupvpcassociation.html#cfn-ec2-securitygroupvpcassociation-groupid
	//
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The association's VPC ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupvpcassociation.html#cfn-ec2-securitygroupvpcassociation-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

