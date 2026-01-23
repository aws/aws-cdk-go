package awsbedrockagentcorealpha


// Options for configuring an interceptor.
//
// Example:
//   // Create Lambda functions for interceptors
//   requestInterceptorFn := lambda.NewFunction(this, jsii.String("RequestInterceptor"), &FunctionProps{
//   	Runtime: lambda.Runtime_PYTHON_3_12(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String(`
//   	def handler(event, context):
//   	    # Validate and transform request
//   	    return {
//   	        "interceptorOutputVersion": "1.0",
//   	        "mcp": {
//   	            "transformedGatewayRequest": event["mcp"]["gatewayRequest"]
//   	        }
//   	    }
//   	  `)),
//   })
//
//   responseInterceptorFn := lambda.NewFunction(this, jsii.String("ResponseInterceptor"), &FunctionProps{
//   	Runtime: lambda.Runtime_PYTHON_3_12(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_*FromInline(jsii.String(`
//   	def handler(event, context):
//   	    # Filter or transform response
//   	    return {
//   	        "interceptorOutputVersion": "1.0",
//   	        "mcp": {
//   	            "transformedGatewayResponse": event["mcp"]["gatewayResponse"]
//   	        }
//   	    }
//   	  `)),
//   })
//
//   // Create gateway with interceptors
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   	InterceptorConfigurations: []IInterceptor{
//   		agentcore.LambdaInterceptor_ForRequest(requestInterceptorFn, &InterceptorOptions{
//   			PassRequestHeaders: jsii.Boolean(true),
//   		}),
//   		agentcore.LambdaInterceptor_ForResponse(responseInterceptorFn),
//   	},
//   })
//
// Experimental.
type InterceptorOptions struct {
	// Whether to pass request headers to the interceptor Lambda function.
	//
	// **Security Warning**: Request headers can contain sensitive information such as
	// authentication tokens and credentials. Only enable this if your interceptor needs
	// access to headers and you have verified that sensitive information is not logged
	// or exposed.
	// Default: false - Headers are not passed to interceptor for security.
	//
	// Experimental.
	PassRequestHeaders *bool `field:"optional" json:"passRequestHeaders" yaml:"passRequestHeaders"`
}

