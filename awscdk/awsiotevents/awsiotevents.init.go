package awsiotevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_iotevents.CfnDetectorModel",
		reflect.TypeOf((*CfnDetectorModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "detectorModelDefinition", GoGetter: "DetectorModelDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "detectorModelDescription", GoGetter: "DetectorModelDescription"},
			_jsii_.MemberProperty{JsiiProperty: "detectorModelName", GoGetter: "DetectorModelName"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationMethod", GoGetter: "EvaluationMethod"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDetectorModel{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.ActionProperty",
		reflect.TypeOf((*CfnDetectorModel_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.AssetPropertyTimestampProperty",
		reflect.TypeOf((*CfnDetectorModel_AssetPropertyTimestampProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.AssetPropertyValueProperty",
		reflect.TypeOf((*CfnDetectorModel_AssetPropertyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.AssetPropertyVariantProperty",
		reflect.TypeOf((*CfnDetectorModel_AssetPropertyVariantProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.ClearTimerProperty",
		reflect.TypeOf((*CfnDetectorModel_ClearTimerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.DetectorModelDefinitionProperty",
		reflect.TypeOf((*CfnDetectorModel_DetectorModelDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.DynamoDBProperty",
		reflect.TypeOf((*CfnDetectorModel_DynamoDBProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.DynamoDBv2Property",
		reflect.TypeOf((*CfnDetectorModel_DynamoDBv2Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.EventProperty",
		reflect.TypeOf((*CfnDetectorModel_EventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.FirehoseProperty",
		reflect.TypeOf((*CfnDetectorModel_FirehoseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.IotEventsProperty",
		reflect.TypeOf((*CfnDetectorModel_IotEventsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.IotSiteWiseProperty",
		reflect.TypeOf((*CfnDetectorModel_IotSiteWiseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.IotTopicPublishProperty",
		reflect.TypeOf((*CfnDetectorModel_IotTopicPublishProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.LambdaProperty",
		reflect.TypeOf((*CfnDetectorModel_LambdaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.OnEnterProperty",
		reflect.TypeOf((*CfnDetectorModel_OnEnterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.OnExitProperty",
		reflect.TypeOf((*CfnDetectorModel_OnExitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.OnInputProperty",
		reflect.TypeOf((*CfnDetectorModel_OnInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.PayloadProperty",
		reflect.TypeOf((*CfnDetectorModel_PayloadProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.ResetTimerProperty",
		reflect.TypeOf((*CfnDetectorModel_ResetTimerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.SetTimerProperty",
		reflect.TypeOf((*CfnDetectorModel_SetTimerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.SetVariableProperty",
		reflect.TypeOf((*CfnDetectorModel_SetVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.SnsProperty",
		reflect.TypeOf((*CfnDetectorModel_SnsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.SqsProperty",
		reflect.TypeOf((*CfnDetectorModel_SqsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.StateProperty",
		reflect.TypeOf((*CfnDetectorModel_StateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModel.TransitionEventProperty",
		reflect.TypeOf((*CfnDetectorModel_TransitionEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnDetectorModelProps",
		reflect.TypeOf((*CfnDetectorModelProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_iotevents.CfnInput",
		reflect.TypeOf((*CfnInput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "inputDefinition", GoGetter: "InputDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "inputDescription", GoGetter: "InputDescription"},
			_jsii_.MemberProperty{JsiiProperty: "inputName", GoGetter: "InputName"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInput{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnInput.AttributeProperty",
		reflect.TypeOf((*CfnInput_AttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnInput.InputDefinitionProperty",
		reflect.TypeOf((*CfnInput_InputDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.CfnInputProps",
		reflect.TypeOf((*CfnInputProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_iotevents.IInput",
		reflect.TypeOf((*IInput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "inputName", GoGetter: "InputName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IInput{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_iotevents.Input",
		reflect.TypeOf((*Input)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "inputName", GoGetter: "InputName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_Input{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInput)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iotevents.InputProps",
		reflect.TypeOf((*InputProps)(nil)).Elem(),
	)
}
