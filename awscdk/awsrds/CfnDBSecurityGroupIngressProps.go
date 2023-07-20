package awsrds


// Properties for defining a `CfnDBSecurityGroupIngress`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBSecurityGroupIngressProps := &CfnDBSecurityGroupIngressProps{
//   	DbSecurityGroupName: jsii.String("dbSecurityGroupName"),
//
//   	// the properties below are optional
//   	Cidrip: jsii.String("cidrip"),
//   	Ec2SecurityGroupId: jsii.String("ec2SecurityGroupId"),
//   	Ec2SecurityGroupName: jsii.String("ec2SecurityGroupName"),
//   	Ec2SecurityGroupOwnerId: jsii.String("ec2SecurityGroupOwnerId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroupingress.html
//
type CfnDBSecurityGroupIngressProps struct {
	// The name of the DB security group to add authorization to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroupingress.html#cfn-rds-dbsecuritygroupingress-dbsecuritygroupname
	//
	DbSecurityGroupName *string `field:"required" json:"dbSecurityGroupName" yaml:"dbSecurityGroupName"`
	// The IP range to authorize.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroupingress.html#cfn-rds-dbsecuritygroupingress-cidrip
	//
	Cidrip *string `field:"optional" json:"cidrip" yaml:"cidrip"`
	// Id of the EC2 security group to authorize.
	//
	// For VPC DB security groups, `EC2SecurityGroupId` must be provided. Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroupingress.html#cfn-rds-dbsecuritygroupingress-ec2securitygroupid
	//
	Ec2SecurityGroupId *string `field:"optional" json:"ec2SecurityGroupId" yaml:"ec2SecurityGroupId"`
	// Name of the EC2 security group to authorize.
	//
	// For VPC DB security groups, `EC2SecurityGroupId` must be provided. Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroupingress.html#cfn-rds-dbsecuritygroupingress-ec2securitygroupname
	//
	Ec2SecurityGroupName *string `field:"optional" json:"ec2SecurityGroupName" yaml:"ec2SecurityGroupName"`
	// AWS account number of the owner of the EC2 security group specified in the `EC2SecurityGroupName` parameter.
	//
	// The AWS access key ID isn't an acceptable value. For VPC DB security groups, `EC2SecurityGroupId` must be provided. Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroupingress.html#cfn-rds-dbsecuritygroupingress-ec2securitygroupownerid
	//
	Ec2SecurityGroupOwnerId *string `field:"optional" json:"ec2SecurityGroupOwnerId" yaml:"ec2SecurityGroupOwnerId"`
}

