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
type IntegrationType string

const (
	// For integrating the API method request with an AWS service action, including the Lambda function-invoking action.
	//
	// With the Lambda
	// function-invoking action, this is referred to as the Lambda custom
	// integration. With any other AWS service action, this is known as AWS
	// integration.
	IntegrationType_AWS IntegrationType = "AWS"
	// For integrating the API method request with the Lambda function-invoking action with the client request passed through as-is.
	//
	// This integration is
	// also referred to as the Lambda proxy integration.
	IntegrationType_AWS_PROXY IntegrationType = "AWS_PROXY"
	// For integrating the API method request with an HTTP endpoint, including a private HTTP endpoint within a VPC.
	//
	// This integration is also referred to
	// as the HTTP custom integration.
	IntegrationType_HTTP IntegrationType = "HTTP"
	// For integrating the API method request with an HTTP endpoint, including a private HTTP endpoint within a VPC, with the client request passed through as-is.
	//
	// This is also referred to as the HTTP proxy integration.
	IntegrationType_HTTP_PROXY IntegrationType = "HTTP_PROXY"
	// For integrating the API method request with API Gateway as a "loop-back" endpoint without invoking any backend.
	IntegrationType_MOCK IntegrationType = "MOCK"
)

