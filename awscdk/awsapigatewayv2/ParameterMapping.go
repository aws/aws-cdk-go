package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a Parameter Mapping.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var lb applicationLoadBalancer
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
type ParameterMapping interface {
	// Represents all created parameter mappings.
	Mappings() *map[string]*string
	// Creates a mapping to append a header.
	AppendHeader(name *string, value MappingValue) ParameterMapping
	// Creates a mapping to append a query string.
	AppendQueryString(name *string, value MappingValue) ParameterMapping
	// Creates a custom mapping.
	Custom(key *string, value *string) ParameterMapping
	// Creates a mapping to overwrite a header.
	OverwriteHeader(name *string, value MappingValue) ParameterMapping
	// Creates a mapping to overwrite a path.
	OverwritePath(value MappingValue) ParameterMapping
	// Creates a mapping to overwrite a querystring.
	OverwriteQueryString(name *string, value MappingValue) ParameterMapping
	// Creates a mapping to remove a header.
	RemoveHeader(name *string) ParameterMapping
	// Creates a mapping to remove a querystring.
	RemoveQueryString(name *string) ParameterMapping
}

// The jsii proxy struct for ParameterMapping
type jsiiProxy_ParameterMapping struct {
	_ byte // padding
}

func (j *jsiiProxy_ParameterMapping) Mappings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"mappings",
		&returns,
	)
	return returns
}


func NewParameterMapping() ParameterMapping {
	_init_.Initialize()

	j := jsiiProxy_ParameterMapping{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.ParameterMapping",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewParameterMapping_Override(p ParameterMapping) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.ParameterMapping",
		nil, // no parameters
		p,
	)
}

// Creates a mapping from an object.
func ParameterMapping_FromObject(obj *map[string]MappingValue) ParameterMapping {
	_init_.Initialize()

	if err := validateParameterMapping_FromObjectParameters(obj); err != nil {
		panic(err)
	}
	var returns ParameterMapping

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.ParameterMapping",
		"fromObject",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParameterMapping) AppendHeader(name *string, value MappingValue) ParameterMapping {
	if err := p.validateAppendHeaderParameters(name, value); err != nil {
		panic(err)
	}
	var returns ParameterMapping

	_jsii_.Invoke(
		p,
		"appendHeader",
		[]interface{}{name, value},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParameterMapping) AppendQueryString(name *string, value MappingValue) ParameterMapping {
	if err := p.validateAppendQueryStringParameters(name, value); err != nil {
		panic(err)
	}
	var returns ParameterMapping

	_jsii_.Invoke(
		p,
		"appendQueryString",
		[]interface{}{name, value},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParameterMapping) Custom(key *string, value *string) ParameterMapping {
	if err := p.validateCustomParameters(key, value); err != nil {
		panic(err)
	}
	var returns ParameterMapping

	_jsii_.Invoke(
		p,
		"custom",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParameterMapping) OverwriteHeader(name *string, value MappingValue) ParameterMapping {
	if err := p.validateOverwriteHeaderParameters(name, value); err != nil {
		panic(err)
	}
	var returns ParameterMapping

	_jsii_.Invoke(
		p,
		"overwriteHeader",
		[]interface{}{name, value},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParameterMapping) OverwritePath(value MappingValue) ParameterMapping {
	if err := p.validateOverwritePathParameters(value); err != nil {
		panic(err)
	}
	var returns ParameterMapping

	_jsii_.Invoke(
		p,
		"overwritePath",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParameterMapping) OverwriteQueryString(name *string, value MappingValue) ParameterMapping {
	if err := p.validateOverwriteQueryStringParameters(name, value); err != nil {
		panic(err)
	}
	var returns ParameterMapping

	_jsii_.Invoke(
		p,
		"overwriteQueryString",
		[]interface{}{name, value},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParameterMapping) RemoveHeader(name *string) ParameterMapping {
	if err := p.validateRemoveHeaderParameters(name); err != nil {
		panic(err)
	}
	var returns ParameterMapping

	_jsii_.Invoke(
		p,
		"removeHeader",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParameterMapping) RemoveQueryString(name *string) ParameterMapping {
	if err := p.validateRemoveQueryStringParameters(name); err != nil {
		panic(err)
	}
	var returns ParameterMapping

	_jsii_.Invoke(
		p,
		"removeQueryString",
		[]interface{}{name},
		&returns,
	)

	return returns
}

