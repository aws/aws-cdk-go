package awsstepfunctionstasks


// Http Methods that API Gateway supports.
//
// Example:
//   import apigatewayv2 "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   httpApi := apigatewayv2.NewHttpApi(this, jsii.String("MyHttpApi"))
//
//   invokeTask := tasks.NewCallApiGatewayHttpApiEndpoint(this, jsii.String("Call HTTP API"), &CallApiGatewayHttpApiEndpointProps{
//   	ApiId: httpApi.ApiId,
//   	ApiStack: awscdk.*stack_Of(httpApi),
//   	Method: tasks.HttpMethod_GET,
//   })
//
type HttpMethod string

const (
	// Retreive data from a server at the specified resource.
	HttpMethod_GET HttpMethod = "GET"
	// Send data to the API endpoint to create or udpate a resource.
	HttpMethod_POST HttpMethod = "POST"
	// Send data to the API endpoint to update or create a resource.
	HttpMethod_PUT HttpMethod = "PUT"
	// Delete the resource at the specified endpoint.
	HttpMethod_DELETE HttpMethod = "DELETE"
	// Apply partial modifications to the resource.
	HttpMethod_PATCH HttpMethod = "PATCH"
	// Retreive data from a server at the specified resource without the response body.
	HttpMethod_HEAD HttpMethod = "HEAD"
	// Return data describing what other methods and operations the server supports.
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
)

