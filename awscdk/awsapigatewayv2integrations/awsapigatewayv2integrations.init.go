package awsapigatewayv2integrations

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_apigatewayv2_integrations.HttpAlbIntegration",
		reflect.TypeOf((*HttpAlbIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectionType", GoGetter: "ConnectionType"},
			_jsii_.MemberProperty{JsiiProperty: "httpMethod", GoGetter: "HttpMethod"},
			_jsii_.MemberProperty{JsiiProperty: "integrationType", GoGetter: "IntegrationType"},
			_jsii_.MemberProperty{JsiiProperty: "payloadFormatVersion", GoGetter: "PayloadFormatVersion"},
		},
		func() interface{} {
			j := jsiiProxy_HttpAlbIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2IHttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apigatewayv2_integrations.HttpAlbIntegrationProps",
		reflect.TypeOf((*HttpAlbIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apigatewayv2_integrations.HttpNlbIntegration",
		reflect.TypeOf((*HttpNlbIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectionType", GoGetter: "ConnectionType"},
			_jsii_.MemberProperty{JsiiProperty: "httpMethod", GoGetter: "HttpMethod"},
			_jsii_.MemberProperty{JsiiProperty: "integrationType", GoGetter: "IntegrationType"},
			_jsii_.MemberProperty{JsiiProperty: "payloadFormatVersion", GoGetter: "PayloadFormatVersion"},
		},
		func() interface{} {
			j := jsiiProxy_HttpNlbIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2IHttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apigatewayv2_integrations.HttpNlbIntegrationProps",
		reflect.TypeOf((*HttpNlbIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apigatewayv2_integrations.HttpPrivateIntegrationOptions",
		reflect.TypeOf((*HttpPrivateIntegrationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apigatewayv2_integrations.HttpProxyIntegration",
		reflect.TypeOf((*HttpProxyIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpProxyIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2IHttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apigatewayv2_integrations.HttpProxyIntegrationProps",
		reflect.TypeOf((*HttpProxyIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apigatewayv2_integrations.HttpServiceDiscoveryIntegration",
		reflect.TypeOf((*HttpServiceDiscoveryIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectionType", GoGetter: "ConnectionType"},
			_jsii_.MemberProperty{JsiiProperty: "httpMethod", GoGetter: "HttpMethod"},
			_jsii_.MemberProperty{JsiiProperty: "integrationType", GoGetter: "IntegrationType"},
			_jsii_.MemberProperty{JsiiProperty: "payloadFormatVersion", GoGetter: "PayloadFormatVersion"},
		},
		func() interface{} {
			j := jsiiProxy_HttpServiceDiscoveryIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2IHttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apigatewayv2_integrations.HttpServiceDiscoveryIntegrationProps",
		reflect.TypeOf((*HttpServiceDiscoveryIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apigatewayv2_integrations.LambdaProxyIntegration",
		reflect.TypeOf((*LambdaProxyIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaProxyIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2IHttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apigatewayv2_integrations.LambdaProxyIntegrationProps",
		reflect.TypeOf((*LambdaProxyIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apigatewayv2_integrations.LambdaWebSocketIntegration",
		reflect.TypeOf((*LambdaWebSocketIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaWebSocketIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2IWebSocketRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apigatewayv2_integrations.LambdaWebSocketIntegrationProps",
		reflect.TypeOf((*LambdaWebSocketIntegrationProps)(nil)).Elem(),
	)
}
