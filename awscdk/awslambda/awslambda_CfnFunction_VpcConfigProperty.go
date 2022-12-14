package awslambda


// The VPC security groups and subnets that are attached to a Lambda function.
//
// When you connect a function to a VPC, Lambda creates an elastic network interface for each combination of security group and subnet in the function's VPC configuration. The function can only access resources and the internet through that VPC. For more information, see [VPC Settings](https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html) .
//
// > When you delete a function, AWS CloudFormation monitors the state of its network interfaces and waits for Lambda to delete them before proceeding. If the VPC is defined in the same stack, the network interfaces need to be deleted by Lambda before AWS CloudFormation can delete the VPC's resources.
// >
// > To monitor network interfaces, AWS CloudFormation needs the `ec2:DescribeNetworkInterfaces` permission. It obtains this from the user or role that modifies the stack. If you don't provide this permission, AWS CloudFormation does not wait for network interfaces to be deleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigProperty := &vpcConfigProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnFunction_VpcConfigProperty struct {
	// A list of VPC security groups IDs.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of VPC subnet IDs.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

