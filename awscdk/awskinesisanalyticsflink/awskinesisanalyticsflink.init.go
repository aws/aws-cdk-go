package awskinesisanalyticsflink

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_kinesisanalytics_flink.Application",
		reflect.TypeOf((*Application)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_Application{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApplication)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_kinesisanalytics_flink.ApplicationCode",
		reflect.TypeOf((*ApplicationCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ApplicationCode{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisanalytics_flink.ApplicationCodeConfig",
		reflect.TypeOf((*ApplicationCodeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisanalytics_flink.ApplicationProps",
		reflect.TypeOf((*ApplicationProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_kinesisanalytics_flink.IApplication",
		reflect.TypeOf((*IApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IApplication{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_kinesisanalytics_flink.LogLevel",
		reflect.TypeOf((*LogLevel)(nil)).Elem(),
		map[string]interface{}{
			"DEBUG": LogLevel_DEBUG,
			"INFO": LogLevel_INFO,
			"WARN": LogLevel_WARN,
			"ERROR": LogLevel_ERROR,
		},
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_kinesisanalytics_flink.MetricsLevel",
		reflect.TypeOf((*MetricsLevel)(nil)).Elem(),
		map[string]interface{}{
			"APPLICATION": MetricsLevel_APPLICATION,
			"TASK": MetricsLevel_TASK,
			"OPERATOR": MetricsLevel_OPERATOR,
			"PARALLELISM": MetricsLevel_PARALLELISM,
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisanalytics_flink.PropertyGroups",
		reflect.TypeOf((*PropertyGroups)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_kinesisanalytics_flink.Runtime",
		reflect.TypeOf((*Runtime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_Runtime{}
		},
	)
}
