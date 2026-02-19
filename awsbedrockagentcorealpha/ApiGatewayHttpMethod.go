package awsbedrockagentcorealpha


// HTTP methods supported by API Gateway.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   api := apigateway.NewRestApi(this, jsii.String("MyApi"), &RestApiProps{
//   	RestApiName: jsii.String("my-api"),
//   })
//
//   // Uses IAM authorization for outbound auth by default
//   apiGatewayTarget := gateway.AddApiGatewayTarget(jsii.String("MyApiGatewayTarget"), &AddApiGatewayTargetOptions{
//   	RestApi: api,
//   	ApiGatewayToolConfiguration: &ApiGatewayToolConfiguration{
//   		ToolFilters: []ApiGatewayToolFilter{
//   			&ApiGatewayToolFilter{
//   				FilterPath: jsii.String("/pets/*"),
//   				Methods: []ApiGatewayHttpMethod{
//   					agentcore.ApiGatewayHttpMethod_GET,
//   				},
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type ApiGatewayHttpMethod string

const (
	// GET method.
	// Experimental.
	ApiGatewayHttpMethod_GET ApiGatewayHttpMethod = "GET"
	// POST method.
	// Experimental.
	ApiGatewayHttpMethod_POST ApiGatewayHttpMethod = "POST"
	// PUT method.
	// Experimental.
	ApiGatewayHttpMethod_PUT ApiGatewayHttpMethod = "PUT"
	// DELETE method.
	// Experimental.
	ApiGatewayHttpMethod_DELETE ApiGatewayHttpMethod = "DELETE"
	// PATCH method.
	// Experimental.
	ApiGatewayHttpMethod_PATCH ApiGatewayHttpMethod = "PATCH"
	// HEAD method.
	// Experimental.
	ApiGatewayHttpMethod_HEAD ApiGatewayHttpMethod = "HEAD"
	// OPTIONS method.
	// Experimental.
	ApiGatewayHttpMethod_OPTIONS ApiGatewayHttpMethod = "OPTIONS"
)

