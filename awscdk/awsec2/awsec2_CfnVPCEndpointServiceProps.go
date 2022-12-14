package awsec2


// Properties for defining a `CfnVPCEndpointService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCEndpointServiceProps := &cfnVPCEndpointServiceProps{
//   	acceptanceRequired: jsii.Boolean(false),
//   	contributorInsightsEnabled: jsii.Boolean(false),
//   	gatewayLoadBalancerArns: []*string{
//   		jsii.String("gatewayLoadBalancerArns"),
//   	},
//   	networkLoadBalancerArns: []*string{
//   		jsii.String("networkLoadBalancerArns"),
//   	},
//   	payerResponsibility: jsii.String("payerResponsibility"),
//   }
//
type CfnVPCEndpointServiceProps struct {
	// Indicates whether requests from service consumers to create an endpoint to your service must be accepted.
	AcceptanceRequired interface{} `field:"optional" json:"acceptanceRequired" yaml:"acceptanceRequired"`
	// `AWS::EC2::VPCEndpointService.ContributorInsightsEnabled`.
	ContributorInsightsEnabled interface{} `field:"optional" json:"contributorInsightsEnabled" yaml:"contributorInsightsEnabled"`
	// The Amazon Resource Names (ARNs) of one or more Gateway Load Balancers.
	GatewayLoadBalancerArns *[]*string `field:"optional" json:"gatewayLoadBalancerArns" yaml:"gatewayLoadBalancerArns"`
	// The Amazon Resource Names (ARNs) of one or more Network Load Balancers for your service.
	NetworkLoadBalancerArns *[]*string `field:"optional" json:"networkLoadBalancerArns" yaml:"networkLoadBalancerArns"`
	// The entity that is responsible for the endpoint costs.
	//
	// The default is the endpoint owner. If you set the payer responsibility to the service owner, you cannot set it back to the endpoint owner.
	PayerResponsibility *string `field:"optional" json:"payerResponsibility" yaml:"payerResponsibility"`
}

