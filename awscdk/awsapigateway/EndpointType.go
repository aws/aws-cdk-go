package awsapigateway


// Example:
//   var apiDefinition apiDefinition
//
//
//   api := apigateway.NewSpecRestApi(this, jsii.String("ExampleRestApi"), &SpecRestApiProps{
//   	ApiDefinition: ApiDefinition,
//   	EndpointTypes: []endpointType{
//   		apigateway.*endpointType_PRIVATE,
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

