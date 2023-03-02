package awsapigateway


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
type ConnectionType string

const (
	// For connections through the public routable internet.
	// Experimental.
	ConnectionType_INTERNET ConnectionType = "INTERNET"
	// For private connections between API Gateway and a network load balancer in a VPC.
	// Experimental.
	ConnectionType_VPC_LINK ConnectionType = "VPC_LINK"
)

