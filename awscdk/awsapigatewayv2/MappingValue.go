package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a Mapping Value.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var lb ApplicationLoadBalancer
//
//   listener := lb.AddListener(jsii.String("listener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(80),
//   })
//   listener.AddTargets(jsii.String("target"), &AddApplicationTargetsProps{
//   	Port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &HttpApiProps{
//   	DefaultIntegration: awscdk.NewHttpAlbIntegration(jsii.String("DefaultIntegration"), listener, &HttpAlbIntegrationProps{
//   		ParameterMapping: apigwv2.NewParameterMapping().AppendHeader(jsii.String("header2"), apigwv2.MappingValue_RequestHeader(jsii.String("header1"))).RemoveHeader(jsii.String("header1")),
//   	}),
//   })
//
type MappingValue interface {
	IMappingValue
	// Represents a Mapping Value.
	Value() *string
}

// The jsii proxy struct for MappingValue
type jsiiProxy_MappingValue struct {
	jsiiProxy_IMappingValue
}

func (j *jsiiProxy_MappingValue) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewMappingValue(value *string) MappingValue {
	_init_.Initialize()

	if err := validateNewMappingValueParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_MappingValue{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.MappingValue",
		[]interface{}{value},
		&j,
	)

	return &j
}

func NewMappingValue_Override(m MappingValue, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.MappingValue",
		[]interface{}{value},
		m,
	)
}

// Creates a context variable mapping value.
func MappingValue_ContextVariable(variableName *string) MappingValue {
	_init_.Initialize()

	if err := validateMappingValue_ContextVariableParameters(variableName); err != nil {
		panic(err)
	}
	var returns MappingValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.MappingValue",
		"contextVariable",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Creates a custom mapping value.
func MappingValue_Custom(value *string) MappingValue {
	_init_.Initialize()

	if err := validateMappingValue_CustomParameters(value); err != nil {
		panic(err)
	}
	var returns MappingValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.MappingValue",
		"custom",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Creates a request body mapping value.
func MappingValue_RequestBody(name *string) MappingValue {
	_init_.Initialize()

	if err := validateMappingValue_RequestBodyParameters(name); err != nil {
		panic(err)
	}
	var returns MappingValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.MappingValue",
		"requestBody",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Creates a header mapping value.
func MappingValue_RequestHeader(name *string) MappingValue {
	_init_.Initialize()

	if err := validateMappingValue_RequestHeaderParameters(name); err != nil {
		panic(err)
	}
	var returns MappingValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.MappingValue",
		"requestHeader",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Creates a request path mapping value.
func MappingValue_RequestPath() MappingValue {
	_init_.Initialize()

	var returns MappingValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.MappingValue",
		"requestPath",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Creates a request path parameter mapping value.
func MappingValue_RequestPathParam(name *string) MappingValue {
	_init_.Initialize()

	if err := validateMappingValue_RequestPathParamParameters(name); err != nil {
		panic(err)
	}
	var returns MappingValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.MappingValue",
		"requestPathParam",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Creates a query string mapping value.
func MappingValue_RequestQueryString(name *string) MappingValue {
	_init_.Initialize()

	if err := validateMappingValue_RequestQueryStringParameters(name); err != nil {
		panic(err)
	}
	var returns MappingValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.MappingValue",
		"requestQueryString",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Creates a stage variable mapping value.
func MappingValue_StageVariable(variableName *string) MappingValue {
	_init_.Initialize()

	if err := validateMappingValue_StageVariableParameters(variableName); err != nil {
		panic(err)
	}
	var returns MappingValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.MappingValue",
		"stageVariable",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

func MappingValue_NONE() MappingValue {
	_init_.Initialize()
	var returns MappingValue
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.MappingValue",
		"NONE",
		&returns,
	)
	return returns
}

