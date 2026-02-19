// The CDK Construct Library for AWS Lambda in Golang
package awscdklambdagoalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-lambda-go-alpha.BundlingOptions",
		reflect.TypeOf((*BundlingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		reflect.TypeOf((*GoFunction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAlias", GoMethod: "AddAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addEnvironment", GoMethod: "AddEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "addEventSource", GoMethod: "AddEventSource"},
			_jsii_.MemberMethod{JsiiMethod: "addEventSourceMapping", GoMethod: "AddEventSourceMapping"},
			_jsii_.MemberMethod{JsiiMethod: "addFunctionUrl", GoMethod: "AddFunctionUrl"},
			_jsii_.MemberMethod{JsiiMethod: "addLayers", GoMethod: "AddLayers"},
			_jsii_.MemberMethod{JsiiMethod: "addPermission", GoMethod: "AddPermission"},
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "architecture", GoGetter: "Architecture"},
			_jsii_.MemberProperty{JsiiProperty: "canCreatePermissions", GoGetter: "CanCreatePermissions"},
			_jsii_.MemberMethod{JsiiMethod: "configureAsyncInvoke", GoMethod: "ConfigureAsyncInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberMethod{JsiiMethod: "considerWarningOnInvokeFunctionPermissions", GoMethod: "ConsiderWarningOnInvokeFunctionPermissions"},
			_jsii_.MemberProperty{JsiiProperty: "currentVersion", GoGetter: "CurrentVersion"},
			_jsii_.MemberProperty{JsiiProperty: "deadLetterQueue", GoGetter: "DeadLetterQueue"},
			_jsii_.MemberProperty{JsiiProperty: "deadLetterTopic", GoGetter: "DeadLetterTopic"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "functionArn", GoGetter: "FunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
			_jsii_.MemberProperty{JsiiProperty: "functionRef", GoGetter: "FunctionRef"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeCompositePrincipal", GoMethod: "GrantInvokeCompositePrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeLatestVersion", GoMethod: "GrantInvokeLatestVersion"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeUrl", GoMethod: "GrantInvokeUrl"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeVersion", GoMethod: "GrantInvokeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "invalidateVersionBasedOn", GoMethod: "InvalidateVersionBasedOn"},
			_jsii_.MemberProperty{JsiiProperty: "isBoundToVpc", GoGetter: "IsBoundToVpc"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersion", GoGetter: "LatestVersion"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricDuration", GoMethod: "MetricDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrors", GoMethod: "MetricErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottles", GoMethod: "MetricThrottles"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissionsNode", GoGetter: "PermissionsNode"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceArnsForGrantInvoke", GoGetter: "ResourceArnsForGrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "runtime", GoGetter: "Runtime"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tenancyConfig", GoGetter: "TenancyConfig"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "warnInvokeFunctionPermissions", GoMethod: "WarnInvokeFunctionPermissions"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_GoFunction{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaFunction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-lambda-go-alpha.GoFunctionProps",
		reflect.TypeOf((*GoFunctionProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-lambda-go-alpha.ICommandHooks",
		reflect.TypeOf((*ICommandHooks)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "afterBundling", GoMethod: "AfterBundling"},
			_jsii_.MemberMethod{JsiiMethod: "beforeBundling", GoMethod: "BeforeBundling"},
		},
		func() interface{} {
			return &jsiiProxy_ICommandHooks{}
		},
	)
}
