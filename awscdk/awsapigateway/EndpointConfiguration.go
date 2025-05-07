package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The endpoint configuration of a REST API, including VPCs and endpoint types.
//
// EndpointConfiguration is a property of the AWS::ApiGateway::RestApi resource.
//
// Example:
//   var someEndpoint iVpcEndpoint
//
//
//   api := apigateway.NewRestApi(this, jsii.String("api"), &RestApiProps{
//   	EndpointConfiguration: &EndpointConfiguration{
//   		Types: []endpointType{
//   			apigateway.*endpointType_PRIVATE,
//   		},
//   		VpcEndpoints: []*iVpcEndpoint{
//   			someEndpoint,
//   		},
//   	},
//   })
//
type EndpointConfiguration struct {
	// A list of endpoint types of an API or its custom domain name.
	// Default: EndpointType.EDGE
	//
	Types *[]EndpointType `field:"required" json:"types" yaml:"types"`
	// The IP address types that can invoke the API.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-ip-address-type.html
	//
	// Default: undefined - AWS default is DUAL_STACK for private API, IPV4 for all other APIs.
	//
	IpAddressType IpAddressType `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// A list of VPC Endpoints against which to create Route53 ALIASes.
	// Default: - no ALIASes are created for the endpoint.
	//
	VpcEndpoints *[]awsec2.IVpcEndpoint `field:"optional" json:"vpcEndpoints" yaml:"vpcEndpoints"`
}

