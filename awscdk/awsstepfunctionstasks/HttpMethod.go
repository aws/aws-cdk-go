package awsstepfunctionstasks


// Http Methods that API Gateway supports.
//
// Example:
//   import apigateway "github.com/aws/aws-cdk-go/awscdk"
//   var api RestApi
//
//
//   tasks.CallApiGatewayRestApiEndpoint_Jsonata(this, jsii.String("Endpoint"), &CallApiGatewayRestApiEndpointJsonataProps{
//   	Api: Api,
//   	StageName: jsii.String("Stage"),
//   	Method: tasks.HttpMethod_PUT,
//   	IntegrationPattern: sfn.IntegrationPattern_WAIT_FOR_TASK_TOKEN,
//   	Headers: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"TaskToken": jsii.String("{% States.Array($states.context.taskToken) %}"),
//   	}),
//   })
//
type HttpMethod string

const (
	// Retrieve data from a server at the specified resource.
	HttpMethod_GET HttpMethod = "GET"
	// Send data to the API endpoint to create or udpate a resource.
	HttpMethod_POST HttpMethod = "POST"
	// Send data to the API endpoint to update or create a resource.
	HttpMethod_PUT HttpMethod = "PUT"
	// Delete the resource at the specified endpoint.
	HttpMethod_DELETE HttpMethod = "DELETE"
	// Apply partial modifications to the resource.
	HttpMethod_PATCH HttpMethod = "PATCH"
	// Retrieve data from a server at the specified resource without the response body.
	HttpMethod_HEAD HttpMethod = "HEAD"
	// Return data describing what other methods and operations the server supports.
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
)

