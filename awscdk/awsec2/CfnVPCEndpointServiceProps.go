package awsec2


// Properties for defining a `CfnVPCEndpointService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCEndpointServiceProps := &CfnVPCEndpointServiceProps{
//   	AcceptanceRequired: jsii.Boolean(false),
//   	ContributorInsightsEnabled: jsii.Boolean(false),
//   	GatewayLoadBalancerArns: []*string{
//   		jsii.String("gatewayLoadBalancerArns"),
//   	},
//   	NetworkLoadBalancerArns: []*string{
//   		jsii.String("networkLoadBalancerArns"),
//   	},
//   	PayerResponsibility: jsii.String("payerResponsibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservice.html
//
type CfnVPCEndpointServiceProps struct {
	// Indicates whether requests from service consumers to create an endpoint to your service must be accepted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservice.html#cfn-ec2-vpcendpointservice-acceptancerequired
	//
	AcceptanceRequired interface{} `field:"optional" json:"acceptanceRequired" yaml:"acceptanceRequired"`
	// Indicates whether to enable the built-in Contributor Insights rules provided by AWS PrivateLink .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservice.html#cfn-ec2-vpcendpointservice-contributorinsightsenabled
	//
	ContributorInsightsEnabled interface{} `field:"optional" json:"contributorInsightsEnabled" yaml:"contributorInsightsEnabled"`
	// The Amazon Resource Names (ARNs) of the Gateway Load Balancers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservice.html#cfn-ec2-vpcendpointservice-gatewayloadbalancerarns
	//
	GatewayLoadBalancerArns *[]*string `field:"optional" json:"gatewayLoadBalancerArns" yaml:"gatewayLoadBalancerArns"`
	// The Amazon Resource Names (ARNs) of the Network Load Balancers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservice.html#cfn-ec2-vpcendpointservice-networkloadbalancerarns
	//
	NetworkLoadBalancerArns *[]*string `field:"optional" json:"networkLoadBalancerArns" yaml:"networkLoadBalancerArns"`
	// The entity that is responsible for the endpoint costs.
	//
	// The default is the endpoint owner. If you set the payer responsibility to the service owner, you cannot set it back to the endpoint owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservice.html#cfn-ec2-vpcendpointservice-payerresponsibility
	//
	PayerResponsibility *string `field:"optional" json:"payerResponsibility" yaml:"payerResponsibility"`
}

