package awsredshift


// Properties for defining a `CfnClusterSecurityGroupIngress`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterSecurityGroupIngressProps := &CfnClusterSecurityGroupIngressProps{
//   	ClusterSecurityGroupName: jsii.String("clusterSecurityGroupName"),
//
//   	// the properties below are optional
//   	Cidrip: jsii.String("cidrip"),
//   	Ec2SecurityGroupName: jsii.String("ec2SecurityGroupName"),
//   	Ec2SecurityGroupOwnerId: jsii.String("ec2SecurityGroupOwnerId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html
//
type CfnClusterSecurityGroupIngressProps struct {
	// The name of the security group to which the ingress rule is added.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-clustersecuritygroupname
	//
	ClusterSecurityGroupName *string `field:"required" json:"clusterSecurityGroupName" yaml:"clusterSecurityGroupName"`
	// The IP range to be added the Amazon Redshift security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-cidrip
	//
	Cidrip *string `field:"optional" json:"cidrip" yaml:"cidrip"`
	// The EC2 security group to be added the Amazon Redshift security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-ec2securitygroupname
	//
	Ec2SecurityGroupName *string `field:"optional" json:"ec2SecurityGroupName" yaml:"ec2SecurityGroupName"`
	// The AWS account number of the owner of the security group specified by the *EC2SecurityGroupName* parameter.
	//
	// The AWS Access Key ID is not an acceptable value.
	//
	// Example: `111122223333`
	//
	// Conditional. If you specify the `EC2SecurityGroupName` property, you must specify this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-ec2securitygroupownerid
	//
	Ec2SecurityGroupOwnerId *string `field:"optional" json:"ec2SecurityGroupOwnerId" yaml:"ec2SecurityGroupOwnerId"`
}

