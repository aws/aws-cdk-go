package awscdksyntheticsalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-synthetics-alpha.ArtifactsBucketLocation",
		reflect.TypeOf((*ArtifactsBucketLocation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-synthetics-alpha.AssetCode",
		reflect.TypeOf((*AssetCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_AssetCode{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Code)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-synthetics-alpha.Canary",
		reflect.TypeOf((*Canary)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsBucket", GoGetter: "ArtifactsBucket"},
			_jsii_.MemberProperty{JsiiProperty: "canaryId", GoGetter: "CanaryId"},
			_jsii_.MemberProperty{JsiiProperty: "canaryName", GoGetter: "CanaryName"},
			_jsii_.MemberProperty{JsiiProperty: "canaryState", GoGetter: "CanaryState"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "metricDuration", GoMethod: "MetricDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricSuccessPercent", GoMethod: "MetricSuccessPercent"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Canary{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-synthetics-alpha.CanaryProps",
		reflect.TypeOf((*CanaryProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-synthetics-alpha.Code",
		reflect.TypeOf((*Code)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_Code{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-synthetics-alpha.CodeConfig",
		reflect.TypeOf((*CodeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-synthetics-alpha.CronOptions",
		reflect.TypeOf((*CronOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-synthetics-alpha.CustomTestOptions",
		reflect.TypeOf((*CustomTestOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-synthetics-alpha.InlineCode",
		reflect.TypeOf((*InlineCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_InlineCode{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Code)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-synthetics-alpha.Runtime",
		reflect.TypeOf((*Runtime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "family", GoGetter: "Family"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_Runtime{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-synthetics-alpha.RuntimeFamily",
		reflect.TypeOf((*RuntimeFamily)(nil)).Elem(),
		map[string]interface{}{
			"NODEJS": RuntimeFamily_NODEJS,
			"PYTHON": RuntimeFamily_PYTHON,
			"OTHER": RuntimeFamily_OTHER,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-synthetics-alpha.S3Code",
		reflect.TypeOf((*S3Code)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_S3Code{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Code)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-synthetics-alpha.Schedule",
		reflect.TypeOf((*Schedule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "expressionString", GoGetter: "ExpressionString"},
		},
		func() interface{} {
			return &jsiiProxy_Schedule{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-synthetics-alpha.Test",
		reflect.TypeOf((*Test)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "code", GoGetter: "Code"},
			_jsii_.MemberProperty{JsiiProperty: "handler", GoGetter: "Handler"},
		},
		func() interface{} {
			return &jsiiProxy_Test{}
		},
	)
}
