package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAgent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAgentProps := &CfnAgentProps{
//   	ActivationKey: jsii.String("activationKey"),
//   	AgentName: jsii.String("agentName"),
//   	SecurityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//   	SubnetArns: []*string{
//   		jsii.String("subnetArns"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcEndpointId: jsii.String("vpcEndpointId"),
//   }
//
type CfnAgentProps struct {
	// Specifies your DataSync agent's activation key.
	//
	// If you don't have an activation key, see [Activate your agent](https://docs.aws.amazon.com/datasync/latest/userguide/activate-agent.html) .
	ActivationKey *string `field:"optional" json:"activationKey" yaml:"activationKey"`
	// Specifies a name for your agent.
	//
	// You can see this name in the DataSync console.
	AgentName *string `field:"optional" json:"agentName" yaml:"agentName"`
	// The Amazon Resource Names (ARNs) of the security groups used to protect your data transfer task subnets.
	//
	// See [SecurityGroupArns](https://docs.aws.amazon.com/datasync/latest/userguide/API_Ec2Config.html#DataSync-Type-Ec2Config-SecurityGroupArns) .
	//
	// *Pattern* : `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`
	SecurityGroupArns *[]*string `field:"optional" json:"securityGroupArns" yaml:"securityGroupArns"`
	// Specifies the ARN of the subnet where you want to run your DataSync task when using a VPC endpoint.
	//
	// This is the subnet where DataSync creates and manages the [network interfaces](https://docs.aws.amazon.com/datasync/latest/userguide/datasync-network.html#required-network-interfaces) for your transfer.
	SubnetArns *[]*string `field:"optional" json:"subnetArns" yaml:"subnetArns"`
	// Specifies labels that help you categorize, filter, and search for your AWS resources.
	//
	// We recommend creating at least one tag for your agent.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the virtual private cloud (VPC) endpoint that the agent has access to.
	//
	// This is the client-side VPC endpoint, powered by AWS PrivateLink . If you don't have an AWS PrivateLink VPC endpoint, see [AWS PrivateLink and VPC endpoints](https://docs.aws.amazon.com//vpc/latest/userguide/endpoint-services-overview.html) in the *Amazon VPC User Guide* .
	//
	// For more information about activating your agent in a private network based on a VPC, see [Using AWS DataSync in a Virtual Private Cloud](https://docs.aws.amazon.com/datasync/latest/userguide/datasync-in-vpc.html) in the *AWS DataSync User Guide.*
	//
	// A VPC endpoint ID looks like this: `vpce-01234d5aff67890e1` .
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

