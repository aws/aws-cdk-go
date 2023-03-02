package awsmediaconnect


// Properties for defining a `CfnFlowVpcInterface`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowVpcInterfaceProps := &cfnFlowVpcInterfaceProps{
//   	flowArn: jsii.String("flowArn"),
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetId: jsii.String("subnetId"),
//   }
//
type CfnFlowVpcInterfaceProps struct {
	// The Amazon Resource Name (ARN) of the flow.
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The name of the VPC Interface.
	//
	// This value must be unique within the current flow.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the role that you created when you set up MediaConnect as a trusted service.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The VPC security groups that you want MediaConnect to use for your VPC configuration.
	//
	// You must include at least one security group in the request.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The subnet IDs that you want to use for your VPC interface.
	//
	// A range of IP addresses in your VPC. When you create your VPC, you specify a range of IPv4 addresses for the VPC in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16. This is the primary CIDR block for your VPC. When you create a subnet for your VPC, you specify the CIDR block for the subnet, which is a subset of the VPC CIDR block.
	//
	// The subnets that you use across all VPC interfaces on the flow must be in the same Availability Zone as the flow.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

