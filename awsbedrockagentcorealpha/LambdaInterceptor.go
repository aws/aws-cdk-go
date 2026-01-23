package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Lambda-based interceptor for Gateway.
//
// Interceptors allow you to run custom code during each invocation of your gateway:
// - REQUEST interceptors execute before calling the target
// - RESPONSE interceptors execute after the target responds.
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
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/gateway-interceptors.html
//
// Experimental.
type LambdaInterceptor interface {
	IInterceptor
	// The interception point (REQUEST or RESPONSE).
	// Experimental.
	InterceptionPoint() InterceptionPoint
	// Binds this Lambda interceptor to a Gateway.
	//
	// This method:
	// 1. Grants the Gateway's IAM role permission to invoke the Lambda function
	// 2. Returns the CloudFormation configuration for this interceptor
	//
	// Returns: Configuration for CloudFormation rendering.
	// Experimental.
	Bind(scope constructs.Construct, gateway IGateway) *InterceptorBindConfig
}

// The jsii proxy struct for LambdaInterceptor
type jsiiProxy_LambdaInterceptor struct {
	jsiiProxy_IInterceptor
}

func (j *jsiiProxy_LambdaInterceptor) InterceptionPoint() InterceptionPoint {
	var returns InterceptionPoint
	_jsii_.Get(
		j,
		"interceptionPoint",
		&returns,
	)
	return returns
}


// Create a REQUEST interceptor that executes before the gateway calls the target.
//
// **Important:** When this interceptor is added to a gateway, the gateway's IAM role
// will automatically be granted `lambda:InvokeFunction` permission on the specified
// Lambda function.
//
// Returns: A configured LambdaInterceptor for request interception.
// Experimental.
func LambdaInterceptor_ForRequest(lambdaFunction awslambda.IFunction, options *InterceptorOptions) LambdaInterceptor {
	_init_.Initialize()

	if err := validateLambdaInterceptor_ForRequestParameters(lambdaFunction, options); err != nil {
		panic(err)
	}
	var returns LambdaInterceptor

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LambdaInterceptor",
		"forRequest",
		[]interface{}{lambdaFunction, options},
		&returns,
	)

	return returns
}

// Create a RESPONSE interceptor that executes after the target responds.
//
// **Important:** When this interceptor is added to a gateway, the gateway's IAM role
// will automatically be granted `lambda:InvokeFunction` permission on the specified
// Lambda function.
//
// Returns: A configured LambdaInterceptor for response interception.
// Experimental.
func LambdaInterceptor_ForResponse(lambdaFunction awslambda.IFunction, options *InterceptorOptions) LambdaInterceptor {
	_init_.Initialize()

	if err := validateLambdaInterceptor_ForResponseParameters(lambdaFunction, options); err != nil {
		panic(err)
	}
	var returns LambdaInterceptor

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LambdaInterceptor",
		"forResponse",
		[]interface{}{lambdaFunction, options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInterceptor) Bind(scope constructs.Construct, gateway IGateway) *InterceptorBindConfig {
	if err := l.validateBindParameters(scope, gateway); err != nil {
		panic(err)
	}
	var returns *InterceptorBindConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, gateway},
		&returns,
	)

	return returns
}

