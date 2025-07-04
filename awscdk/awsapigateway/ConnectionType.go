package awsapigateway


// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   nlb := elbv2.NewNetworkLoadBalancer(this, jsii.String("NLB"), &NetworkLoadBalancerProps{
//   	Vpc: Vpc,
//   })
//   link := apigateway.NewVpcLink(this, jsii.String("link"), &VpcLinkProps{
//   	Targets: []iNetworkLoadBalancer{
//   		nlb,
//   	},
//   })
//
//   integration := apigateway.NewIntegration(&IntegrationProps{
//   	Type: apigateway.IntegrationType_HTTP_PROXY,
//   	IntegrationHttpMethod: jsii.String("ANY"),
//   	Options: &IntegrationOptions{
//   		ConnectionType: apigateway.ConnectionType_VPC_LINK,
//   		VpcLink: link,
//   	},
//   })
//
type ConnectionType string

const (
	// For connections through the public routable internet.
	ConnectionType_INTERNET ConnectionType = "INTERNET"
	// For private connections between API Gateway and a network load balancer in a VPC.
	ConnectionType_VPC_LINK ConnectionType = "VPC_LINK"
)

