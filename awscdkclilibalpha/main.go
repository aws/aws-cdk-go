// AWS CDK Programmatic CLI library
package awscdkclilibalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/cli-lib-alpha.AwsCdkCli",
		reflect.TypeOf((*AwsCdkCli)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bootstrap", GoMethod: "Bootstrap"},
			_jsii_.MemberMethod{JsiiMethod: "deploy", GoMethod: "Deploy"},
			_jsii_.MemberMethod{JsiiMethod: "destroy", GoMethod: "Destroy"},
			_jsii_.MemberMethod{JsiiMethod: "list", GoMethod: "List"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
		},
		func() interface{} {
			j := jsiiProxy_AwsCdkCli{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAwsCdkCli)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cli-lib-alpha.BootstrapOptions",
		reflect.TypeOf((*BootstrapOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cli-lib-alpha.CdkAppDirectoryProps",
		reflect.TypeOf((*CdkAppDirectoryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cli-lib-alpha.DeployOptions",
		reflect.TypeOf((*DeployOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cli-lib-alpha.DestroyOptions",
		reflect.TypeOf((*DestroyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/cli-lib-alpha.IAwsCdkCli",
		reflect.TypeOf((*IAwsCdkCli)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bootstrap", GoMethod: "Bootstrap"},
			_jsii_.MemberMethod{JsiiMethod: "deploy", GoMethod: "Deploy"},
			_jsii_.MemberMethod{JsiiMethod: "destroy", GoMethod: "Destroy"},
			_jsii_.MemberMethod{JsiiMethod: "list", GoMethod: "List"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
		},
		func() interface{} {
			return &jsiiProxy_IAwsCdkCli{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/cli-lib-alpha.ICloudAssemblyDirectoryProducer",
		reflect.TypeOf((*ICloudAssemblyDirectoryProducer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "produce", GoMethod: "Produce"},
			_jsii_.MemberProperty{JsiiProperty: "workingDirectory", GoGetter: "WorkingDirectory"},
		},
		func() interface{} {
			return &jsiiProxy_ICloudAssemblyDirectoryProducer{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cli-lib-alpha.ListOptions",
		reflect.TypeOf((*ListOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/cli-lib-alpha.RequireApproval",
		reflect.TypeOf((*RequireApproval)(nil)).Elem(),
		map[string]interface{}{
			"NEVER": RequireApproval_NEVER,
			"ANYCHANGE": RequireApproval_ANYCHANGE,
			"BROADENING": RequireApproval_BROADENING,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cli-lib-alpha.SharedOptions",
		reflect.TypeOf((*SharedOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/cli-lib-alpha.StackActivityProgress",
		reflect.TypeOf((*StackActivityProgress)(nil)).Elem(),
		map[string]interface{}{
			"BAR": StackActivityProgress_BAR,
			"EVENTS": StackActivityProgress_EVENTS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cli-lib-alpha.SynthOptions",
		reflect.TypeOf((*SynthOptions)(nil)).Elem(),
	)
}
