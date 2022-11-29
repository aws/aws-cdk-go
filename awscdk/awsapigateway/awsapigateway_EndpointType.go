package awsapigateway


// Example:
//   var apiDefinition apiDefinition
//
//
//   api := apigateway.NewSpecRestApi(this, jsii.String("ExampleRestApi"), &specRestApiProps{
//   	apiDefinition: apiDefinition,
//   	endpointTypes: []endpointType{
//   		apigateway.*endpointType_PRIVATE,
//   	},
//   })
//
// Experimental.
type EndpointType string

const (
	// For an edge-optimized API and its custom domain name.
	// Experimental.
	EndpointType_EDGE EndpointType = "EDGE"
	// For a regional API and its custom domain name.
	// Experimental.
	EndpointType_REGIONAL EndpointType = "REGIONAL"
	// For a private API and its custom domain name.
	// Experimental.
	EndpointType_PRIVATE EndpointType = "PRIVATE"
)

