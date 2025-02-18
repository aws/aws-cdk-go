package awsapigateway


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
type EndpointType string

const (
	// For an edge-optimized API and its custom domain name.
	EndpointType_EDGE EndpointType = "EDGE"
	// For a regional API and its custom domain name.
	EndpointType_REGIONAL EndpointType = "REGIONAL"
	// For a private API and its custom domain name.
	EndpointType_PRIVATE EndpointType = "PRIVATE"
)

