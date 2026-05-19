package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

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
//   					agentcore.ApiGatewayHttpMethod_GET(),
//   				},
//   			},
//   		},
//   	},
//   })
//
type ApiGatewayHttpMethod interface {
	// The HTTP method string value.
	Value() *string
	// Returns the string value.
	ToString() *string
}

// The jsii proxy struct for ApiGatewayHttpMethod
type jsiiProxy_ApiGatewayHttpMethod struct {
	_ byte // padding
}

func (j *jsiiProxy_ApiGatewayHttpMethod) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom HTTP method not yet defined in this class.
func ApiGatewayHttpMethod_Of(value *string) ApiGatewayHttpMethod {
	_init_.Initialize()

	if err := validateApiGatewayHttpMethod_OfParameters(value); err != nil {
		panic(err)
	}
	var returns ApiGatewayHttpMethod

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayHttpMethod",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ApiGatewayHttpMethod_DELETE() ApiGatewayHttpMethod {
	_init_.Initialize()
	var returns ApiGatewayHttpMethod
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayHttpMethod",
		"DELETE",
		&returns,
	)
	return returns
}

func ApiGatewayHttpMethod_GET() ApiGatewayHttpMethod {
	_init_.Initialize()
	var returns ApiGatewayHttpMethod
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayHttpMethod",
		"GET",
		&returns,
	)
	return returns
}

func ApiGatewayHttpMethod_HEAD() ApiGatewayHttpMethod {
	_init_.Initialize()
	var returns ApiGatewayHttpMethod
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayHttpMethod",
		"HEAD",
		&returns,
	)
	return returns
}

func ApiGatewayHttpMethod_OPTIONS() ApiGatewayHttpMethod {
	_init_.Initialize()
	var returns ApiGatewayHttpMethod
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayHttpMethod",
		"OPTIONS",
		&returns,
	)
	return returns
}

func ApiGatewayHttpMethod_PATCH() ApiGatewayHttpMethod {
	_init_.Initialize()
	var returns ApiGatewayHttpMethod
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayHttpMethod",
		"PATCH",
		&returns,
	)
	return returns
}

func ApiGatewayHttpMethod_POST() ApiGatewayHttpMethod {
	_init_.Initialize()
	var returns ApiGatewayHttpMethod
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayHttpMethod",
		"POST",
		&returns,
	)
	return returns
}

func ApiGatewayHttpMethod_PUT() ApiGatewayHttpMethod {
	_init_.Initialize()
	var returns ApiGatewayHttpMethod
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayHttpMethod",
		"PUT",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApiGatewayHttpMethod) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

