package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// The endpoint configuration of a REST API, including VPCs and endpoint types.
//
// EndpointConfiguration is a property of the AWS::ApiGateway::RestApi resource.
//
// Example:
//   api := apigateway.NewRestApi(this, jsii.String("api"), &restApiProps{
//   	endpointConfiguration: &endpointConfiguration{
//   		types: []endpointType{
//   			apigateway.*endpointType_EDGE,
//   		},
//   	},
//   })
//
// Experimental.
type EndpointConfiguration struct {
	// A list of endpoint types of an API or its custom domain name.
	// Experimental.
	Types *[]EndpointType `field:"required" json:"types" yaml:"types"`
	// A list of VPC Endpoints against which to create Route53 ALIASes.
	// Experimental.
	VpcEndpoints *[]awsec2.IVpcEndpoint `field:"optional" json:"vpcEndpoints" yaml:"vpcEndpoints"`
}

