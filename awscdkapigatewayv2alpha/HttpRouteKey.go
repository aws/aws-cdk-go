package awscdkapigatewayv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HTTP route in APIGateway is a combination of the HTTP method and the path component.
//
// This class models that combination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   httpRouteKey := apigatewayv2_alpha.HttpRouteKey_With(jsii.String("path"), apigatewayv2_alpha.HttpMethod_ANY)
//
// Experimental.
type HttpRouteKey interface {
	// The key to the RouteKey as recognized by APIGateway.
	// Experimental.
	Key() *string
	// The method of the route.
	// Experimental.
	Method() HttpMethod
	// The path part of this RouteKey.
	//
	// Returns `undefined` when `RouteKey.DEFAULT` is used.
	// Experimental.
	Path() *string
}

// The jsii proxy struct for HttpRouteKey
type jsiiProxy_HttpRouteKey struct {
	_ byte // padding
}

func (j *jsiiProxy_HttpRouteKey) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpRouteKey) Method() HttpMethod {
	var returns HttpMethod
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpRouteKey) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}


// Create a route key with the combination of the path and the method.
// Experimental.
func HttpRouteKey_With(path *string, method HttpMethod) HttpRouteKey {
	_init_.Initialize()

	if err := validateHttpRouteKey_WithParameters(path); err != nil {
		panic(err)
	}
	var returns HttpRouteKey

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRouteKey",
		"with",
		[]interface{}{path, method},
		&returns,
	)

	return returns
}

func HttpRouteKey_DEFAULT() HttpRouteKey {
	_init_.Initialize()
	var returns HttpRouteKey
	_jsii_.StaticGet(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRouteKey",
		"DEFAULT",
		&returns,
	)
	return returns
}

