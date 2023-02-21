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
//   cfnAgentProps := &cfnAgentProps{
//   	activationKey: jsii.String("activationKey"),
//
//   	// the properties below are optional
//   	agentName: jsii.String("agentName"),
//   	securityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//   	subnetArns: []*string{
//   		jsii.String("subnetArns"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcEndpointId: jsii.String("vpcEndpointId"),
//   }
//
type CfnAgentProps struct {
	// Your agent activation key.
	//
	// You can get the activation key either by sending an HTTP GET request with redirects that enable you to get the agent IP address (port 80). Alternatively, you can get it from the DataSync console.
	//
	// The redirect URL returned in the response provides you the activation key for your agent in the query string parameter `activationKey` . It might also include other activation-related parameters; however, these are merely defaults. The arguments you pass to this API call determine the actual configuration of your agent.
	//
	// For more information, see [Creating and activating an agent](https://docs.aws.amazon.com/datasync/latest/userguide/activating-agent.html) in the *AWS DataSync User Guide.*
	ActivationKey *string `field:"required" json:"activationKey" yaml:"activationKey"`
	// The name you configured for your agent.
	//
	// This value is a text reference that is used to identify the agent in the console.
	AgentName *string `field:"optional" json:"agentName" yaml:"agentName"`
	// The Amazon Resource Names (ARNs) of the security groups used to protect your data transfer task subnets.
	//
	// See [SecurityGroupArns](https://docs.aws.amazon.com/datasync/latest/userguide/API_Ec2Config.html#DataSync-Type-Ec2Config-SecurityGroupArns) .
	//
	// *Pattern* : `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`
	SecurityGroupArns *[]*string `field:"optional" json:"securityGroupArns" yaml:"securityGroupArns"`
	// The Amazon Resource Names (ARNs) of the subnets in which DataSync will create elastic network interfaces for each data transfer task.
	//
	// The agent that runs a task must be private. When you start a task that is associated with an agent created in a VPC, or one that has access to an IP address in a VPC, then the task is also private. In this case, DataSync creates four network interfaces for each task in your subnet. For a data transfer to work, the agent must be able to route to all these four network interfaces.
	SubnetArns *[]*string `field:"optional" json:"subnetArns" yaml:"subnetArns"`
	// The key-value pair that represents the tag that you want to associate with the agent.
	//
	// The value can be an empty string. This value helps you manage, filter, and search for your agents.
	//
	// > Valid characters for key and value are letters, spaces, and numbers representable in UTF-8 format, and the following special characters: + - = . _ : / @.
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

