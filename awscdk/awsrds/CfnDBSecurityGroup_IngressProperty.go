package awsrds


// The `Ingress` property type specifies an individual ingress rule within an `AWS::RDS::DBSecurityGroup` resource.
//
// > EC2-Classic was retired on August 15, 2022. If you haven't migrated from EC2-Classic to a VPC, we recommend that you migrate as soon as possible. For more information, see [Migrate from EC2-Classic to a VPC](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html) in the *Amazon EC2 User Guide* , the blog [EC2-Classic Networking is Retiring – Here’s How to Prepare](https://docs.aws.amazon.com/aws/ec2-classic-is-retiring-heres-how-to-prepare/) , and [Moving a DB instance not in a VPC into a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Non-VPC2VPC.html) in the *Amazon RDS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressProperty := &IngressProperty{
//   	Cidrip: jsii.String("cidrip"),
//   	Ec2SecurityGroupId: jsii.String("ec2SecurityGroupId"),
//   	Ec2SecurityGroupName: jsii.String("ec2SecurityGroupName"),
//   	Ec2SecurityGroupOwnerId: jsii.String("ec2SecurityGroupOwnerId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbsecuritygroup-ingress.html
//
type CfnDBSecurityGroup_IngressProperty struct {
	// The IP range to authorize.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbsecuritygroup-ingress.html#cfn-rds-dbsecuritygroup-ingress-cidrip
	//
	Cidrip *string `field:"optional" json:"cidrip" yaml:"cidrip"`
	// Id of the EC2 security group to authorize.
	//
	// For VPC DB security groups, `EC2SecurityGroupId` must be provided. Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbsecuritygroup-ingress.html#cfn-rds-dbsecuritygroup-ingress-ec2securitygroupid
	//
	Ec2SecurityGroupId *string `field:"optional" json:"ec2SecurityGroupId" yaml:"ec2SecurityGroupId"`
	// Name of the EC2 security group to authorize.
	//
	// For VPC DB security groups, `EC2SecurityGroupId` must be provided. Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbsecuritygroup-ingress.html#cfn-rds-dbsecuritygroup-ingress-ec2securitygroupname
	//
	Ec2SecurityGroupName *string `field:"optional" json:"ec2SecurityGroupName" yaml:"ec2SecurityGroupName"`
	// AWS account number of the owner of the EC2 security group specified in the `EC2SecurityGroupName` parameter.
	//
	// The AWS access key ID isn't an acceptable value. For VPC DB security groups, `EC2SecurityGroupId` must be provided. Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbsecuritygroup-ingress.html#cfn-rds-dbsecuritygroup-ingress-ec2securitygroupownerid
	//
	Ec2SecurityGroupOwnerId *string `field:"optional" json:"ec2SecurityGroupOwnerId" yaml:"ec2SecurityGroupOwnerId"`
}

