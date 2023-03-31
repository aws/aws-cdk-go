package awskinesisfirehose


// The details of the VPC of the Amazon ES destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigurationProperty := &vpcConfigurationProperty{
//   	roleArn: jsii.String("roleArn"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnDeliveryStream_VpcConfigurationProperty struct {
	// The ARN of the IAM role that you want the delivery stream to use to create endpoints in the destination VPC.
	//
	// You can use your existing Kinesis Data Firehose delivery role or you can specify a new role. In either case, make sure that the role trusts the Kinesis Data Firehose service principal and that it grants the following permissions:
	//
	// - `ec2:DescribeVpcs`
	// - `ec2:DescribeVpcAttribute`
	// - `ec2:DescribeSubnets`
	// - `ec2:DescribeSecurityGroups`
	// - `ec2:DescribeNetworkInterfaces`
	// - `ec2:CreateNetworkInterface`
	// - `ec2:CreateNetworkInterfacePermission`
	// - `ec2:DeleteNetworkInterface`
	//
	// If you revoke these permissions after you create the delivery stream, Kinesis Data Firehose can't scale out by creating more ENIs when necessary. You might therefore see a degradation in performance.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The IDs of the security groups that you want Kinesis Data Firehose to use when it creates ENIs in the VPC of the Amazon ES destination.
	//
	// You can use the same security group that the Amazon ES domain uses or different ones. If you specify different security groups here, ensure that they allow outbound HTTPS traffic to the Amazon ES domain's security group. Also ensure that the Amazon ES domain's security group allows HTTPS traffic from the security groups specified here. If you use the same security group for both your delivery stream and the Amazon ES domain, make sure the security group inbound rule allows HTTPS traffic.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The IDs of the subnets that Kinesis Data Firehose uses to create ENIs in the VPC of the Amazon ES destination.
	//
	// Make sure that the routing tables and inbound and outbound rules allow traffic to flow from the subnets whose IDs are specified here to the subnets that have the destination Amazon ES endpoints. Kinesis Data Firehose creates at least one ENI in each of the subnets that are specified here. Do not delete or modify these ENIs.
	//
	// The number of ENIs that Kinesis Data Firehose creates in the subnets specified here scales up and down automatically based on throughput. To enable Kinesis Data Firehose to scale up the number of ENIs to match throughput, ensure that you have sufficient quota. To help you calculate the quota you need, assume that Kinesis Data Firehose can create up to three ENIs for this delivery stream for each of the subnets specified here.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
}

