package awsstepfunctionstasks


// Http Methods that API Gateway supports.
//
// Example:
//   import apigatewayv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   httpApi := apigatewayv2.NewHttpApi(this, jsii.String("MyHttpApi"))
//
//   invokeTask := tasks.NewCallApiGatewayHttpApiEndpoint(this, jsii.String("Call HTTP API"), &callApiGatewayHttpApiEndpointProps{
//   	apiId: httpApi.apiId,
//   	apiStack: awscdk.*stack.of(httpApi),
//   	method: tasks.httpMethod_GET,
//   })
//
// Experimental.
type HttpMethod string

const (
	// Retreive data from a server at the specified resource.
	// Experimental.
	HttpMethod_GET HttpMethod = "GET"
	// Send data to the API endpoint to create or udpate a resource.
	// Experimental.
	HttpMethod_POST HttpMethod = "POST"
	// Send data to the API endpoint to update or create a resource.
	// Experimental.
	HttpMethod_PUT HttpMethod = "PUT"
	// Delete the resource at the specified endpoint.
	// Experimental.
	HttpMethod_DELETE HttpMethod = "DELETE"
	// Apply partial modifications to the resource.
	// Experimental.
	HttpMethod_PATCH HttpMethod = "PATCH"
	// Retreive data from a server at the specified resource without the response body.
	// Experimental.
	HttpMethod_HEAD HttpMethod = "HEAD"
	// Return data describing what other methods and operations the server supports.
	// Experimental.
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
)

