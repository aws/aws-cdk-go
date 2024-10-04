// The CDK Construct Library for Amazon EventBridge Pipes
package awscdkpipesalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-pipes-alpha.DesiredState",
		reflect.TypeOf((*DesiredState)(nil)).Elem(),
		map[string]interface{}{
			"RUNNING": DesiredState_RUNNING,
			"STOPPED": DesiredState_STOPPED,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-alpha.DynamicInput",
		reflect.TypeOf((*DynamicInput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "displayHint", GoGetter: "DisplayHint"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "toJSON", GoMethod: "ToJSON"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DynamicInput{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResolvable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-alpha.EnrichmentParametersConfig",
		reflect.TypeOf((*EnrichmentParametersConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-alpha.Filter",
		reflect.TypeOf((*Filter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "filters", GoGetter: "Filters"},
		},
		func() interface{} {
			j := jsiiProxy_Filter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFilter)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-alpha.FilterPattern",
		reflect.TypeOf((*FilterPattern)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_FilterPattern{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-pipes-alpha.IEnrichment",
		reflect.TypeOf((*IEnrichment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "enrichmentArn", GoGetter: "EnrichmentArn"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
		},
		func() interface{} {
			return &jsiiProxy_IEnrichment{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-pipes-alpha.IFilter",
		reflect.TypeOf((*IFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "filters", GoGetter: "Filters"},
		},
		func() interface{} {
			return &jsiiProxy_IFilter{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-pipes-alpha.IFilterPattern",
		reflect.TypeOf((*IFilterPattern)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "pattern", GoGetter: "Pattern"},
		},
		func() interface{} {
			return &jsiiProxy_IFilterPattern{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-pipes-alpha.IInputTransformation",
		reflect.TypeOf((*IInputTransformation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IInputTransformation{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-pipes-alpha.ILogDestination",
		reflect.TypeOf((*ILogDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "grantPush", GoMethod: "GrantPush"},
		},
		func() interface{} {
			return &jsiiProxy_ILogDestination{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-pipes-alpha.IPipe",
		reflect.TypeOf((*IPipe)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "pipeArn", GoGetter: "PipeArn"},
			_jsii_.MemberProperty{JsiiProperty: "pipeName", GoGetter: "PipeName"},
			_jsii_.MemberProperty{JsiiProperty: "pipeRole", GoGetter: "PipeRole"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IPipe{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-pipes-alpha.ISource",
		reflect.TypeOf((*ISource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "sourceArn", GoGetter: "SourceArn"},
		},
		func() interface{} {
			return &jsiiProxy_ISource{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-pipes-alpha.ITarget",
		reflect.TypeOf((*ITarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "grantPush", GoMethod: "GrantPush"},
			_jsii_.MemberProperty{JsiiProperty: "targetArn", GoGetter: "TargetArn"},
		},
		func() interface{} {
			return &jsiiProxy_ITarget{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-pipes-alpha.IncludeExecutionData",
		reflect.TypeOf((*IncludeExecutionData)(nil)).Elem(),
		map[string]interface{}{
			"ALL": IncludeExecutionData_ALL,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-alpha.InputTransformation",
		reflect.TypeOf((*InputTransformation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_InputTransformation{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInputTransformation)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-alpha.InputTransformationConfig",
		reflect.TypeOf((*InputTransformationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-alpha.LogDestinationConfig",
		reflect.TypeOf((*LogDestinationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-alpha.LogDestinationParameters",
		reflect.TypeOf((*LogDestinationParameters)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-pipes-alpha.LogLevel",
		reflect.TypeOf((*LogLevel)(nil)).Elem(),
		map[string]interface{}{
			"OFF": LogLevel_OFF,
			"ERROR": LogLevel_ERROR,
			"INFO": LogLevel_INFO,
			"TRACE": LogLevel_TRACE,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-alpha.Pipe",
		reflect.TypeOf((*Pipe)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "pipeArn", GoGetter: "PipeArn"},
			_jsii_.MemberProperty{JsiiProperty: "pipeName", GoGetter: "PipeName"},
			_jsii_.MemberProperty{JsiiProperty: "pipeRole", GoGetter: "PipeRole"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Pipe{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPipe)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-alpha.PipeProps",
		reflect.TypeOf((*PipeProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-pipes-alpha.PipeVariable",
		reflect.TypeOf((*PipeVariable)(nil)).Elem(),
		map[string]interface{}{
			"ARN": PipeVariable_ARN,
			"NAME": PipeVariable_NAME,
			"SOURCE_ARN": PipeVariable_SOURCE_ARN,
			"ENRICHMENT_ARN": PipeVariable_ENRICHMENT_ARN,
			"TARGET_ARN": PipeVariable_TARGET_ARN,
			"EVENT_INGESTION_TIME": PipeVariable_EVENT_INGESTION_TIME,
			"EVENT": PipeVariable_EVENT,
			"EVENT_JSON": PipeVariable_EVENT_JSON,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-alpha.SourceConfig",
		reflect.TypeOf((*SourceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-alpha.SourceParameters",
		reflect.TypeOf((*SourceParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-alpha.SourceWithDeadLetterTarget",
		reflect.TypeOf((*SourceWithDeadLetterTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "deadLetterTarget", GoGetter: "DeadLetterTarget"},
			_jsii_.MemberMethod{JsiiMethod: "getDeadLetterTargetArn", GoMethod: "GetDeadLetterTargetArn"},
			_jsii_.MemberMethod{JsiiMethod: "grantPush", GoMethod: "GrantPush"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "sourceArn", GoGetter: "SourceArn"},
		},
		func() interface{} {
			j := jsiiProxy_SourceWithDeadLetterTarget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-alpha.TargetConfig",
		reflect.TypeOf((*TargetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-alpha.TargetParameter",
		reflect.TypeOf((*TargetParameter)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_TargetParameter{}
		},
	)
}
