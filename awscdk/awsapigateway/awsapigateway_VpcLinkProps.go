package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

// Properties for a VpcLink.
//
// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   nlb := elbv2.NewNetworkLoadBalancer(this, jsii.String("NLB"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   })
//   link := apigateway.NewVpcLink(this, jsii.String("link"), &vpcLinkProps{
//   	targets: []iNetworkLoadBalancer{
//   		nlb,
//   	},
//   })
//
//   integration := apigateway.NewIntegration(&integrationProps{
//   	type: apigateway.integrationType_HTTP_PROXY,
//   	options: &integrationOptions{
//   		connectionType: apigateway.connectionType_VPC_LINK,
//   		vpcLink: link,
//   	},
//   })
//
// Experimental.
type VpcLinkProps struct {
	// The description of the VPC link.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The network load balancers of the VPC targeted by the VPC link.
	//
	// The network load balancers must be owned by the same AWS account of the API owner.
	// Experimental.
	Targets *[]awselasticloadbalancingv2.INetworkLoadBalancer `field:"optional" json:"targets" yaml:"targets"`
	// The name used to label and identify the VPC link.
	// Experimental.
	VpcLinkName *string `field:"optional" json:"vpcLinkName" yaml:"vpcLinkName"`
}

