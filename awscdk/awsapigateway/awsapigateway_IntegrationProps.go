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
type IntegrationProps struct {
	// Specifies an API method integration type.
	Type IntegrationType `field:"required" json:"type" yaml:"type"`
	// The integration's HTTP method type.
	//
	// Required unless you use a MOCK integration.
	IntegrationHttpMethod *string `field:"optional" json:"integrationHttpMethod" yaml:"integrationHttpMethod"`
	// Integration options.
	Options *IntegrationOptions `field:"optional" json:"options" yaml:"options"`
	// The Uniform Resource Identifier (URI) for the integration.
	//
	// - If you specify HTTP for the `type` property, specify the API endpoint URL.
	// - If you specify MOCK for the `type` property, don't specify this property.
	// - If you specify AWS for the `type` property, specify an AWS service that
	//    follows this form: `arn:partition:apigateway:region:subdomain.service|service:path|action/service_api.`
	//    For example, a Lambda function URI follows this form:
	//    arn:partition:apigateway:region:lambda:path/path. The path is usually in the
	//    form /2015-03-31/functions/LambdaFunctionARN/invocations.
	// See: https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#uri
	//
	Uri interface{} `field:"optional" json:"uri" yaml:"uri"`
}

