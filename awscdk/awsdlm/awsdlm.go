package awsdlm

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy",
		reflect.TypeOf((*CfnLifecyclePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArn", GoGetter: "ExecutionRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyDetails", GoGetter: "PolicyDetails"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLifecyclePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.ActionProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.ArchiveRetainRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_ArchiveRetainRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.ArchiveRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_ArchiveRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.CreateRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_CreateRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.CrossRegionCopyActionProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_CrossRegionCopyActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.CrossRegionCopyDeprecateRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_CrossRegionCopyDeprecateRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.CrossRegionCopyRetainRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_CrossRegionCopyRetainRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.CrossRegionCopyRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_CrossRegionCopyRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.DeprecateRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_DeprecateRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.EventParametersProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_EventParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.EventSourceProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_EventSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.FastRestoreRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_FastRestoreRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.ParametersProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_ParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.PolicyDetailsProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_PolicyDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.RetainRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_RetainRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.RetentionArchiveTierProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_RetentionArchiveTierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.ScheduleProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_ScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy.ShareRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicy_ShareRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicyProps",
		reflect.TypeOf((*CfnLifecyclePolicyProps)(nil)).Elem(),
	)
}
