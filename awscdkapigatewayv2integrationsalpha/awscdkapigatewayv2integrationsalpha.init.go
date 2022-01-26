package awscdkapigatewayv2integrationsalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpAlbIntegration",
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
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaHttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpAlbIntegrationProps",
		reflect.TypeOf((*HttpAlbIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpLambdaIntegration",
		reflect.TypeOf((*HttpLambdaIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpLambdaIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaHttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpLambdaIntegrationProps",
		reflect.TypeOf((*HttpLambdaIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpNlbIntegration",
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
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaHttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpNlbIntegrationProps",
		reflect.TypeOf((*HttpNlbIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpPrivateIntegrationOptions",
		reflect.TypeOf((*HttpPrivateIntegrationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpServiceDiscoveryIntegration",
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
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaHttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpServiceDiscoveryIntegrationProps",
		reflect.TypeOf((*HttpServiceDiscoveryIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpUrlIntegration",
		reflect.TypeOf((*HttpUrlIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpUrlIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaHttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpUrlIntegrationProps",
		reflect.TypeOf((*HttpUrlIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.WebSocketLambdaIntegration",
		reflect.TypeOf((*WebSocketLambdaIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_WebSocketLambdaIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaWebSocketRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.WebSocketMockIntegration",
		reflect.TypeOf((*WebSocketMockIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_WebSocketMockIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaWebSocketRouteIntegration)
			return &j
		},
	)
}
