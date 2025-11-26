package previewawsdatasyncmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAgentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentMixinProps := &CfnAgentMixinProps{
//   	ActivationKey: jsii.String("activationKey"),
//   	AgentName: jsii.String("agentName"),
//   	SecurityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//   	SubnetArns: []*string{
//   		jsii.String("subnetArns"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcEndpointId: jsii.String("vpcEndpointId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html
//
type CfnAgentMixinProps struct {
	// Specifies your DataSync agent's activation key.
	//
	// If you don't have an activation key, see [Activating your agent](https://docs.aws.amazon.com/datasync/latest/userguide/activate-agent.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-activationkey
	//
	ActivationKey *string `field:"optional" json:"activationKey" yaml:"activationKey"`
	// Specifies a name for your agent.
	//
	// We recommend specifying a name that you can remember.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-agentname
	//
	AgentName *string `field:"optional" json:"agentName" yaml:"agentName"`
	// The Amazon Resource Names (ARNs) of the security groups used to protect your data transfer task subnets.
	//
	// See [SecurityGroupArns](https://docs.aws.amazon.com/datasync/latest/userguide/API_Ec2Config.html#DataSync-Type-Ec2Config-SecurityGroupArns) .
	//
	// *Pattern* : `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-securitygrouparns
	//
	SecurityGroupArns *[]*string `field:"optional" json:"securityGroupArns" yaml:"securityGroupArns"`
	// Specifies the ARN of the subnet where your VPC service endpoint is located.
	//
	// You can only specify one ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-subnetarns
	//
	SubnetArns *[]*string `field:"optional" json:"subnetArns" yaml:"subnetArns"`
	// Specifies labels that help you categorize, filter, and search for your AWS resources.
	//
	// We recommend creating at least one tag for your agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the virtual private cloud (VPC) endpoint that the agent has access to.
	//
	// This is the client-side VPC endpoint, powered by AWS PrivateLink . If you don't have an AWS PrivateLink VPC endpoint, see [AWS PrivateLink and VPC endpoints](https://docs.aws.amazon.com//vpc/latest/userguide/endpoint-services-overview.html) in the *Amazon VPC User Guide* .
	//
	// For more information about activating your agent in a private network based on a VPC, see [Using AWS DataSync in a Virtual Private Cloud](https://docs.aws.amazon.com/datasync/latest/userguide/datasync-in-vpc.html) in the *AWS DataSync User Guide.*
	//
	// A VPC endpoint ID looks like this: `vpce-01234d5aff67890e1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-agent.html#cfn-datasync-agent-vpcendpointid
	//
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

