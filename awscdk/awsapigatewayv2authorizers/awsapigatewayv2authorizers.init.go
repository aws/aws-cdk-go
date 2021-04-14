package awsapigatewayv2authorizers

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpJwtAuthorizer",
		reflect.TypeOf((*HttpJwtAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
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
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpUserPoolAuthorizer",
		reflect.TypeOf((*HttpUserPoolAuthorizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpUserPoolAuthorizer{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2IHttpRouteAuthorizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.UserPoolAuthorizerProps",
		reflect.TypeOf((*UserPoolAuthorizerProps)(nil)).Elem(),
	)
}
