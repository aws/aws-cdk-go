// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a Mapping Value.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   var lb applicationLoadBalancer
//
//   listener := lb.addListener(jsii.String("listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   listener.addTargets(jsii.String("target"), &addApplicationTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpAlbIntegration(jsii.String("DefaultIntegration"), listener, &httpAlbIntegrationProps{
//   		parameterMapping: apigwv2.NewParameterMapping().appendHeader(jsii.String("header2"), apigwv2.mappingValue.requestHeader(jsii.String("header1"))).removeHeader(jsii.String("header1")),
//   	}),
//   })
//
// Experimental.
type MappingValue interface {
	IMappingValue
	// Represents a Mapping Value.
	// Experimental.
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


// Experimental.
func NewMappingValue(value *string) MappingValue {
	_init_.Initialize()

	j := jsiiProxy_MappingValue{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.MappingValue",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Experimental.
func NewMappingValue_Override(m MappingValue, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.MappingValue",
		[]interface{}{value},
		m,
	)
}

// Creates a context variable mapping value.
// Experimental.
func MappingValue_ContextVariable(variableName *string) MappingValue {
	_init_.Initialize()

	var returns MappingValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.MappingValue",
		"contextVariable",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Creates a custom mapping value.
// Experimental.
func MappingValue_Custom(value *string) MappingValue {
	_init_.Initialize()

	var returns MappingValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.MappingValue",
		"custom",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Creates a request body mapping value.
// Experimental.
func MappingValue_RequestBody(name *string) MappingValue {
	_init_.Initialize()

	var returns MappingValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.MappingValue",
		"requestBody",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Creates a header mapping value.
// Experimental.
func MappingValue_RequestHeader(name *string) MappingValue {
	_init_.Initialize()

	var returns MappingValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.MappingValue",
		"requestHeader",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Creates a request path mapping value.
// Experimental.
func MappingValue_RequestPath() MappingValue {
	_init_.Initialize()

	var returns MappingValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.MappingValue",
		"requestPath",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Creates a request path parameter mapping value.
// Experimental.
func MappingValue_RequestPathParam(name *string) MappingValue {
	_init_.Initialize()

	var returns MappingValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.MappingValue",
		"requestPathParam",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Creates a query string mapping value.
// Experimental.
func MappingValue_RequestQueryString(name *string) MappingValue {
	_init_.Initialize()

	var returns MappingValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.MappingValue",
		"requestQueryString",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Creates a stage variable mapping value.
// Experimental.
func MappingValue_StageVariable(variableName *string) MappingValue {
	_init_.Initialize()

	var returns MappingValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.MappingValue",
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
		"@aws-cdk/aws-apigatewayv2-alpha.MappingValue",
		"NONE",
		&returns,
	)
	return returns
}

