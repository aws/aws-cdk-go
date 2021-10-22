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
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaIHttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpAlbIntegrationProps",
		reflect.TypeOf((*HttpAlbIntegrationProps)(nil)).Elem(),
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
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaIHttpRouteIntegration)
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
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpProxyIntegration",
		reflect.TypeOf((*HttpProxyIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpProxyIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaIHttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpProxyIntegrationProps",
		reflect.TypeOf((*HttpProxyIntegrationProps)(nil)).Elem(),
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
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaIHttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpServiceDiscoveryIntegrationProps",
		reflect.TypeOf((*HttpServiceDiscoveryIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.LambdaProxyIntegration",
		reflect.TypeOf((*LambdaProxyIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaProxyIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaIHttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.LambdaProxyIntegrationProps",
		reflect.TypeOf((*LambdaProxyIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.LambdaWebSocketIntegration",
		reflect.TypeOf((*LambdaWebSocketIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaWebSocketIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkapigatewayv2alphaIWebSocketRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.LambdaWebSocketIntegrationProps",
		reflect.TypeOf((*LambdaWebSocketIntegrationProps)(nil)).Elem(),
	)
}
