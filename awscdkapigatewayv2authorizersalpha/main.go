// Authorizers for AWS APIGateway V2
package awscdkapigatewayv2authorizersalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpIamAuthorizer",
		reflect.TypeOf((*HttpIamAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpIamAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaIHttpRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpJwtAuthorizer",
		reflect.TypeOf((*HttpJwtAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpJwtAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaIHttpRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpJwtAuthorizerProps",
		reflect.TypeOf((*HttpJwtAuthorizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpLambdaAuthorizer",
		reflect.TypeOf((*HttpLambdaAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpLambdaAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaIHttpRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpLambdaAuthorizerProps",
		reflect.TypeOf((*HttpLambdaAuthorizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpLambdaResponseType",
		reflect.TypeOf((*HttpLambdaResponseType)(nil)).Elem(),
		map[string]interface{}{
			"SIMPLE": HttpLambdaResponseType_SIMPLE,
			"IAM": HttpLambdaResponseType_IAM,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpUserPoolAuthorizer",
		reflect.TypeOf((*HttpUserPoolAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpUserPoolAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaIHttpRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpUserPoolAuthorizerProps",
		reflect.TypeOf((*HttpUserPoolAuthorizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.WebSocketIamAuthorizer",
		reflect.TypeOf((*WebSocketIamAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_WebSocketIamAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaIWebSocketRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.WebSocketLambdaAuthorizer",
		reflect.TypeOf((*WebSocketLambdaAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_WebSocketLambdaAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaIWebSocketRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.WebSocketLambdaAuthorizerProps",
		reflect.TypeOf((*WebSocketLambdaAuthorizerProps)(nil)).Elem(),
	)
}
