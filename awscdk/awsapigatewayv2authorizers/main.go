package awsapigatewayv2authorizers

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpIamAuthorizer",
		reflect.TypeOf((*HttpIamAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authorizationType", GoGetter: "AuthorizationType"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpIamAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2IHttpRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpJwtAuthorizer",
		reflect.TypeOf((*HttpJwtAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authorizationType", GoGetter: "AuthorizationType"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerId", GoGetter: "AuthorizerId"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpJwtAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2IHttpRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpJwtAuthorizerProps",
		reflect.TypeOf((*HttpJwtAuthorizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpLambdaAuthorizer",
		reflect.TypeOf((*HttpLambdaAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authorizationType", GoGetter: "AuthorizationType"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerId", GoGetter: "AuthorizerId"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpLambdaAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2IHttpRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpLambdaAuthorizerProps",
		reflect.TypeOf((*HttpLambdaAuthorizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpLambdaResponseType",
		reflect.TypeOf((*HttpLambdaResponseType)(nil)).Elem(),
		map[string]interface{}{
			"SIMPLE": HttpLambdaResponseType_SIMPLE,
			"IAM": HttpLambdaResponseType_IAM,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpUserPoolAuthorizer",
		reflect.TypeOf((*HttpUserPoolAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authorizationType", GoGetter: "AuthorizationType"},
			_jsii_.MemberProperty{JsiiProperty: "authorizerId", GoGetter: "AuthorizerId"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpUserPoolAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2IHttpRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpUserPoolAuthorizerProps",
		reflect.TypeOf((*HttpUserPoolAuthorizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.WebSocketIamAuthorizer",
		reflect.TypeOf((*WebSocketIamAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_WebSocketIamAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2IWebSocketRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.WebSocketLambdaAuthorizer",
		reflect.TypeOf((*WebSocketLambdaAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_WebSocketLambdaAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2IWebSocketRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.WebSocketLambdaAuthorizerProps",
		reflect.TypeOf((*WebSocketLambdaAuthorizerProps)(nil)).Elem(),
	)
}
