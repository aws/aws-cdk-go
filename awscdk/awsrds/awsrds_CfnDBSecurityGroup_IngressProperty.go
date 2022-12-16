package awsrds


// The `Ingress` property type specifies an individual ingress rule within an `AWS::RDS::DBSecurityGroup` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressProperty := &ingressProperty{
//   	cidrip: jsii.String("cidrip"),
//   	ec2SecurityGroupId: jsii.String("ec2SecurityGroupId"),
//   	ec2SecurityGroupName: jsii.String("ec2SecurityGroupName"),
//   	ec2SecurityGroupOwnerId: jsii.String("ec2SecurityGroupOwnerId"),
//   }
//
type CfnDBSecurityGroup_IngressProperty struct {
	// The IP range to authorize.
	Cidrip *string `field:"optional" json:"cidrip" yaml:"cidrip"`
	// Id of the EC2 security group to authorize.
	//
	// For VPC DB security groups, `EC2SecurityGroupId` must be provided. Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.
	Ec2SecurityGroupId *string `field:"optional" json:"ec2SecurityGroupId" yaml:"ec2SecurityGroupId"`
	// Name of the EC2 security group to authorize.
	//
	// For VPC DB security groups, `EC2SecurityGroupId` must be provided. Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.
	Ec2SecurityGroupName *string `field:"optional" json:"ec2SecurityGroupName" yaml:"ec2SecurityGroupName"`
	// AWS account number of the owner of the EC2 security group specified in the `EC2SecurityGroupName` parameter.
	//
	// The AWS access key ID isn't an acceptable value. For VPC DB security groups, `EC2SecurityGroupId` must be provided. Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.
	Ec2SecurityGroupOwnerId *string `field:"optional" json:"ec2SecurityGroupOwnerId" yaml:"ec2SecurityGroupOwnerId"`
}

