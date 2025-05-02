package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties to reference an existing load balancer.
//
// Example:
//   // Create an Accelerator
//   accelerator := globalaccelerator.NewAccelerator(this, jsii.String("Accelerator"))
//
//   // Create a Listener
//   listener := accelerator.AddListener(jsii.String("Listener"), &ListenerOptions{
//   	PortRanges: []portRange{
//   		&portRange{
//   			FromPort: jsii.Number(80),
//   		},
//   		&portRange{
//   			FromPort: jsii.Number(443),
//   		},
//   	},
//   })
//
//   // Import the Load Balancers
//   nlb1 := elbv2.NetworkLoadBalancer_FromNetworkLoadBalancerAttributes(this, jsii.String("NLB1"), &NetworkLoadBalancerAttributes{
//   	LoadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:us-west-2:111111111111:loadbalancer/app/my-load-balancer1/e16bef66805b"),
//   })
//   nlb2 := elbv2.NetworkLoadBalancer_FromNetworkLoadBalancerAttributes(this, jsii.String("NLB2"), &NetworkLoadBalancerAttributes{
//   	LoadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:ap-south-1:111111111111:loadbalancer/app/my-load-balancer2/5513dc2ea8a1"),
//   })
//
//   // Add one EndpointGroup for each Region we are targeting
//   listener.AddEndpointGroup(jsii.String("Group1"), &EndpointGroupOptions{
//   	Endpoints: []iEndpoint{
//   		ga_endpoints.NewNetworkLoadBalancerEndpoint(nlb1),
//   	},
//   })
//   listener.AddEndpointGroup(jsii.String("Group2"), &EndpointGroupOptions{
//   	// Imported load balancers automatically calculate their Region from the ARN.
//   	// If you are load balancing to other resources, you must also pass a `region`
//   	// parameter here.
//   	Endpoints: []*iEndpoint{
//   		ga_endpoints.NewNetworkLoadBalancerEndpoint(nlb2),
//   	},
//   })
//
type NetworkLoadBalancerAttributes struct {
	// ARN of the load balancer.
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// The canonical hosted zone ID of this load balancer.
	// Default: - When not provided, LB cannot be used as Route53 Alias target.
	//
	LoadBalancerCanonicalHostedZoneId *string `field:"optional" json:"loadBalancerCanonicalHostedZoneId" yaml:"loadBalancerCanonicalHostedZoneId"`
	// The DNS name of this load balancer.
	// Default: - When not provided, LB cannot be used as Route53 Alias target.
	//
	LoadBalancerDnsName *string `field:"optional" json:"loadBalancerDnsName" yaml:"loadBalancerDnsName"`
	// Security groups to associate with this load balancer.
	// Default: - No security groups associated with the load balancer.
	//
	LoadBalancerSecurityGroups *[]*string `field:"optional" json:"loadBalancerSecurityGroups" yaml:"loadBalancerSecurityGroups"`
	// The VPC to associate with the load balancer.
	// Default: - When not provided, listeners cannot be created on imported load
	// balancers.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

