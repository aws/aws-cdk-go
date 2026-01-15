package awsapigatewayv2integrations

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpAlbIntegration",
		reflect.TypeOf((*HttpAlbIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "completeBind", GoMethod: "CompleteBind"},
			_jsii_.MemberProperty{JsiiProperty: "connectionType", GoGetter: "ConnectionType"},
			_jsii_.MemberProperty{JsiiProperty: "httpMethod", GoGetter: "HttpMethod"},
			_jsii_.MemberProperty{JsiiProperty: "integrationType", GoGetter: "IntegrationType"},
			_jsii_.MemberProperty{JsiiProperty: "payloadFormatVersion", GoGetter: "PayloadFormatVersion"},
		},
		func() interface{} {
			j := jsiiProxy_HttpAlbIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2HttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpAlbIntegrationProps",
		reflect.TypeOf((*HttpAlbIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpEventBridgeIntegration",
		reflect.TypeOf((*HttpEventBridgeIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "completeBind", GoMethod: "CompleteBind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpEventBridgeIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2HttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpEventBridgeIntegrationProps",
		reflect.TypeOf((*HttpEventBridgeIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpLambdaIntegration",
		reflect.TypeOf((*HttpLambdaIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "completeBind", GoMethod: "CompleteBind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpLambdaIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2HttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpLambdaIntegrationProps",
		reflect.TypeOf((*HttpLambdaIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpNlbIntegration",
		reflect.TypeOf((*HttpNlbIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "completeBind", GoMethod: "CompleteBind"},
			_jsii_.MemberProperty{JsiiProperty: "connectionType", GoGetter: "ConnectionType"},
			_jsii_.MemberProperty{JsiiProperty: "httpMethod", GoGetter: "HttpMethod"},
			_jsii_.MemberProperty{JsiiProperty: "integrationType", GoGetter: "IntegrationType"},
			_jsii_.MemberProperty{JsiiProperty: "payloadFormatVersion", GoGetter: "PayloadFormatVersion"},
		},
		func() interface{} {
			j := jsiiProxy_HttpNlbIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2HttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpNlbIntegrationProps",
		reflect.TypeOf((*HttpNlbIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpPrivateIntegrationOptions",
		reflect.TypeOf((*HttpPrivateIntegrationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpServiceDiscoveryIntegration",
		reflect.TypeOf((*HttpServiceDiscoveryIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "completeBind", GoMethod: "CompleteBind"},
			_jsii_.MemberProperty{JsiiProperty: "connectionType", GoGetter: "ConnectionType"},
			_jsii_.MemberProperty{JsiiProperty: "httpMethod", GoGetter: "HttpMethod"},
			_jsii_.MemberProperty{JsiiProperty: "integrationType", GoGetter: "IntegrationType"},
			_jsii_.MemberProperty{JsiiProperty: "payloadFormatVersion", GoGetter: "PayloadFormatVersion"},
		},
		func() interface{} {
			j := jsiiProxy_HttpServiceDiscoveryIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2HttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpServiceDiscoveryIntegrationProps",
		reflect.TypeOf((*HttpServiceDiscoveryIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpSqsIntegration",
		reflect.TypeOf((*HttpSqsIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "completeBind", GoMethod: "CompleteBind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpSqsIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2HttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpSqsIntegrationProps",
		reflect.TypeOf((*HttpSqsIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpStepFunctionsIntegration",
		reflect.TypeOf((*HttpStepFunctionsIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "completeBind", GoMethod: "CompleteBind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpStepFunctionsIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2HttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpStepFunctionsIntegrationProps",
		reflect.TypeOf((*HttpStepFunctionsIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpUrlIntegration",
		reflect.TypeOf((*HttpUrlIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "completeBind", GoMethod: "CompleteBind"},
		},
		func() interface{} {
			j := jsiiProxy_HttpUrlIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2HttpRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpUrlIntegrationProps",
		reflect.TypeOf((*HttpUrlIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_integrations.WebSocketAwsIntegration",
		reflect.TypeOf((*WebSocketAwsIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_WebSocketAwsIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2WebSocketRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_integrations.WebSocketAwsIntegrationProps",
		reflect.TypeOf((*WebSocketAwsIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_integrations.WebSocketLambdaIntegration",
		reflect.TypeOf((*WebSocketLambdaIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_WebSocketLambdaIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2WebSocketRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_integrations.WebSocketLambdaIntegrationProps",
		reflect.TypeOf((*WebSocketLambdaIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apigatewayv2_integrations.WebSocketMockIntegration",
		reflect.TypeOf((*WebSocketMockIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_WebSocketMockIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awsapigatewayv2WebSocketRouteIntegration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apigatewayv2_integrations.WebSocketMockIntegrationProps",
		reflect.TypeOf((*WebSocketMockIntegrationProps)(nil)).Elem(),
	)
}
