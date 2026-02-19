// The CDK Construct Library for AWS::Amplify
package awscdkapplicationsignalsalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.ApplicationSignalsIntegration",
		reflect.TypeOf((*ApplicationSignalsIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationSignalsIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-applicationsignals-alpha.ApplicationSignalsIntegrationProps",
		reflect.TypeOf((*ApplicationSignalsIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.CloudWatchAgentIntegration",
		reflect.TypeOf((*CloudWatchAgentIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agentContainer", GoGetter: "AgentContainer"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_CloudWatchAgentIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-applicationsignals-alpha.CloudWatchAgentIntegrationProps",
		reflect.TypeOf((*CloudWatchAgentIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-applicationsignals-alpha.CloudWatchAgentOptions",
		reflect.TypeOf((*CloudWatchAgentOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.CloudWatchAgentVersion",
		reflect.TypeOf((*CloudWatchAgentVersion)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CloudWatchAgentVersion{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.CommonExporting",
		reflect.TypeOf((*CommonExporting)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CommonExporting{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.DotNetInjector",
		reflect.TypeOf((*DotNetInjector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberMethod{JsiiMethod: "injectAdditionalEnvironments", GoMethod: "InjectAdditionalEnvironments"},
			_jsii_.MemberMethod{JsiiMethod: "injectInitContainer", GoMethod: "InjectInitContainer"},
			_jsii_.MemberProperty{JsiiProperty: "instrumentationVersion", GoGetter: "InstrumentationVersion"},
			_jsii_.MemberMethod{JsiiMethod: "overrideAdditionalEnvironments", GoMethod: "OverrideAdditionalEnvironments"},
			_jsii_.MemberMethod{JsiiMethod: "renderDefaultContainer", GoMethod: "RenderDefaultContainer"},
			_jsii_.MemberProperty{JsiiProperty: "sharedVolumeName", GoGetter: "SharedVolumeName"},
		},
		func() interface{} {
			j := jsiiProxy_DotNetInjector{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Injector)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.DotNetLinuxInjector",
		reflect.TypeOf((*DotNetLinuxInjector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberMethod{JsiiMethod: "injectAdditionalEnvironments", GoMethod: "InjectAdditionalEnvironments"},
			_jsii_.MemberMethod{JsiiMethod: "injectInitContainer", GoMethod: "InjectInitContainer"},
			_jsii_.MemberProperty{JsiiProperty: "instrumentationVersion", GoGetter: "InstrumentationVersion"},
			_jsii_.MemberMethod{JsiiMethod: "overrideAdditionalEnvironments", GoMethod: "OverrideAdditionalEnvironments"},
			_jsii_.MemberMethod{JsiiMethod: "renderDefaultContainer", GoMethod: "RenderDefaultContainer"},
			_jsii_.MemberProperty{JsiiProperty: "sharedVolumeName", GoGetter: "SharedVolumeName"},
		},
		func() interface{} {
			j := jsiiProxy_DotNetLinuxInjector{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DotNetInjector)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.DotNetWindowsInjector",
		reflect.TypeOf((*DotNetWindowsInjector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberMethod{JsiiMethod: "injectAdditionalEnvironments", GoMethod: "InjectAdditionalEnvironments"},
			_jsii_.MemberMethod{JsiiMethod: "injectInitContainer", GoMethod: "InjectInitContainer"},
			_jsii_.MemberProperty{JsiiProperty: "instrumentationVersion", GoGetter: "InstrumentationVersion"},
			_jsii_.MemberMethod{JsiiMethod: "overrideAdditionalEnvironments", GoMethod: "OverrideAdditionalEnvironments"},
			_jsii_.MemberMethod{JsiiMethod: "renderDefaultContainer", GoMethod: "RenderDefaultContainer"},
			_jsii_.MemberProperty{JsiiProperty: "sharedVolumeName", GoGetter: "SharedVolumeName"},
		},
		func() interface{} {
			j := jsiiProxy_DotNetWindowsInjector{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DotNetInjector)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		reflect.TypeOf((*DotnetInstrumentation)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DotnetInstrumentation{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentationVersion",
		reflect.TypeOf((*DotnetInstrumentationVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "imageRepo", GoGetter: "ImageRepo"},
			_jsii_.MemberMethod{JsiiMethod: "imageURI", GoMethod: "ImageURI"},
			_jsii_.MemberProperty{JsiiProperty: "memoryLimit", GoGetter: "MemoryLimit"},
			_jsii_.MemberMethod{JsiiMethod: "memoryLimitMiB", GoMethod: "MemoryLimitMiB"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_DotnetInstrumentationVersion{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InstrumentationVersion)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-applicationsignals-alpha.EnvironmentExtension",
		reflect.TypeOf((*EnvironmentExtension)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.Injector",
		reflect.TypeOf((*Injector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberMethod{JsiiMethod: "injectAdditionalEnvironments", GoMethod: "InjectAdditionalEnvironments"},
			_jsii_.MemberMethod{JsiiMethod: "injectInitContainer", GoMethod: "InjectInitContainer"},
			_jsii_.MemberProperty{JsiiProperty: "instrumentationVersion", GoGetter: "InstrumentationVersion"},
			_jsii_.MemberMethod{JsiiMethod: "overrideAdditionalEnvironments", GoMethod: "OverrideAdditionalEnvironments"},
			_jsii_.MemberMethod{JsiiMethod: "renderDefaultContainer", GoMethod: "RenderDefaultContainer"},
			_jsii_.MemberProperty{JsiiProperty: "sharedVolumeName", GoGetter: "SharedVolumeName"},
		},
		func() interface{} {
			return &jsiiProxy_Injector{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-applicationsignals-alpha.InstrumentationProps",
		reflect.TypeOf((*InstrumentationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.InstrumentationVersion",
		reflect.TypeOf((*InstrumentationVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "imageRepo", GoGetter: "ImageRepo"},
			_jsii_.MemberMethod{JsiiMethod: "imageURI", GoMethod: "ImageURI"},
			_jsii_.MemberProperty{JsiiProperty: "memoryLimit", GoGetter: "MemoryLimit"},
			_jsii_.MemberMethod{JsiiMethod: "memoryLimitMiB", GoMethod: "MemoryLimitMiB"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_InstrumentationVersion{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInjector",
		reflect.TypeOf((*JavaInjector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberMethod{JsiiMethod: "injectAdditionalEnvironments", GoMethod: "InjectAdditionalEnvironments"},
			_jsii_.MemberMethod{JsiiMethod: "injectInitContainer", GoMethod: "InjectInitContainer"},
			_jsii_.MemberProperty{JsiiProperty: "instrumentationVersion", GoGetter: "InstrumentationVersion"},
			_jsii_.MemberMethod{JsiiMethod: "overrideAdditionalEnvironments", GoMethod: "OverrideAdditionalEnvironments"},
			_jsii_.MemberMethod{JsiiMethod: "renderDefaultContainer", GoMethod: "RenderDefaultContainer"},
			_jsii_.MemberProperty{JsiiProperty: "sharedVolumeName", GoGetter: "SharedVolumeName"},
		},
		func() interface{} {
			j := jsiiProxy_JavaInjector{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Injector)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInstrumentation",
		reflect.TypeOf((*JavaInstrumentation)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_JavaInstrumentation{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInstrumentationVersion",
		reflect.TypeOf((*JavaInstrumentationVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "imageRepo", GoGetter: "ImageRepo"},
			_jsii_.MemberMethod{JsiiMethod: "imageURI", GoMethod: "ImageURI"},
			_jsii_.MemberProperty{JsiiProperty: "memoryLimit", GoGetter: "MemoryLimit"},
			_jsii_.MemberMethod{JsiiMethod: "memoryLimitMiB", GoMethod: "MemoryLimitMiB"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_JavaInstrumentationVersion{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InstrumentationVersion)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.LogsExporting",
		reflect.TypeOf((*LogsExporting)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_LogsExporting{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.MetricsExporting",
		reflect.TypeOf((*MetricsExporting)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MetricsExporting{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.NodeInjector",
		reflect.TypeOf((*NodeInjector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberMethod{JsiiMethod: "injectAdditionalEnvironments", GoMethod: "InjectAdditionalEnvironments"},
			_jsii_.MemberMethod{JsiiMethod: "injectInitContainer", GoMethod: "InjectInitContainer"},
			_jsii_.MemberProperty{JsiiProperty: "instrumentationVersion", GoGetter: "InstrumentationVersion"},
			_jsii_.MemberMethod{JsiiMethod: "overrideAdditionalEnvironments", GoMethod: "OverrideAdditionalEnvironments"},
			_jsii_.MemberMethod{JsiiMethod: "renderDefaultContainer", GoMethod: "RenderDefaultContainer"},
			_jsii_.MemberProperty{JsiiProperty: "sharedVolumeName", GoGetter: "SharedVolumeName"},
		},
		func() interface{} {
			j := jsiiProxy_NodeInjector{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Injector)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.NodeInstrumentation",
		reflect.TypeOf((*NodeInstrumentation)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_NodeInstrumentation{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.NodeInstrumentationVersion",
		reflect.TypeOf((*NodeInstrumentationVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "imageRepo", GoGetter: "ImageRepo"},
			_jsii_.MemberMethod{JsiiMethod: "imageURI", GoMethod: "ImageURI"},
			_jsii_.MemberProperty{JsiiProperty: "memoryLimit", GoGetter: "MemoryLimit"},
			_jsii_.MemberMethod{JsiiMethod: "memoryLimitMiB", GoMethod: "MemoryLimitMiB"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_NodeInstrumentationVersion{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InstrumentationVersion)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.PythonInjector",
		reflect.TypeOf((*PythonInjector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberMethod{JsiiMethod: "injectAdditionalEnvironments", GoMethod: "InjectAdditionalEnvironments"},
			_jsii_.MemberMethod{JsiiMethod: "injectInitContainer", GoMethod: "InjectInitContainer"},
			_jsii_.MemberProperty{JsiiProperty: "instrumentationVersion", GoGetter: "InstrumentationVersion"},
			_jsii_.MemberMethod{JsiiMethod: "overrideAdditionalEnvironments", GoMethod: "OverrideAdditionalEnvironments"},
			_jsii_.MemberMethod{JsiiMethod: "renderDefaultContainer", GoMethod: "RenderDefaultContainer"},
			_jsii_.MemberProperty{JsiiProperty: "sharedVolumeName", GoGetter: "SharedVolumeName"},
		},
		func() interface{} {
			j := jsiiProxy_PythonInjector{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Injector)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.PythonInstrumentation",
		reflect.TypeOf((*PythonInstrumentation)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_PythonInstrumentation{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.PythonInstrumentationVersion",
		reflect.TypeOf((*PythonInstrumentationVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "imageRepo", GoGetter: "ImageRepo"},
			_jsii_.MemberMethod{JsiiMethod: "imageURI", GoMethod: "ImageURI"},
			_jsii_.MemberProperty{JsiiProperty: "memoryLimit", GoGetter: "MemoryLimit"},
			_jsii_.MemberMethod{JsiiMethod: "memoryLimitMiB", GoMethod: "MemoryLimitMiB"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_PythonInstrumentationVersion{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InstrumentationVersion)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		reflect.TypeOf((*TraceExporting)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_TraceExporting{}
		},
	)
}
