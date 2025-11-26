package previewawsssmmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnAssociationMixinProps",
		reflect.TypeOf((*CfnAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnAssociationPropsMixin",
		reflect.TypeOf((*CfnAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnAssociationPropsMixin.InstanceAssociationOutputLocationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_InstanceAssociationOutputLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnAssociationPropsMixin.S3OutputLocationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_S3OutputLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnAssociationPropsMixin.TargetProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_TargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnDocumentMixinProps",
		reflect.TypeOf((*CfnDocumentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnDocumentPropsMixin",
		reflect.TypeOf((*CfnDocumentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDocumentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnDocumentPropsMixin.AttachmentsSourceProperty",
		reflect.TypeOf((*CfnDocumentPropsMixin_AttachmentsSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnDocumentPropsMixin.DocumentRequiresProperty",
		reflect.TypeOf((*CfnDocumentPropsMixin_DocumentRequiresProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowMixinProps",
		reflect.TypeOf((*CfnMaintenanceWindowMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowPropsMixin",
		reflect.TypeOf((*CfnMaintenanceWindowPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMaintenanceWindowPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTargetMixinProps",
		reflect.TypeOf((*CfnMaintenanceWindowTargetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTargetPropsMixin",
		reflect.TypeOf((*CfnMaintenanceWindowTargetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMaintenanceWindowTargetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTargetPropsMixin.TargetsProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTargetPropsMixin_TargetsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTaskMixinProps",
		reflect.TypeOf((*CfnMaintenanceWindowTaskMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTaskPropsMixin",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMaintenanceWindowTaskPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTaskPropsMixin.CloudWatchOutputConfigProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_CloudWatchOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTaskPropsMixin.LoggingInfoProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_LoggingInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTaskPropsMixin.MaintenanceWindowAutomationParametersProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_MaintenanceWindowAutomationParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTaskPropsMixin.MaintenanceWindowLambdaParametersProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_MaintenanceWindowLambdaParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTaskPropsMixin.MaintenanceWindowRunCommandParametersProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_MaintenanceWindowRunCommandParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTaskPropsMixin.MaintenanceWindowStepFunctionsParametersProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_MaintenanceWindowStepFunctionsParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTaskPropsMixin.NotificationConfigProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_NotificationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTaskPropsMixin.TargetProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_TargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTaskPropsMixin.TaskInvocationParametersProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_TaskInvocationParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnParameterMixinProps",
		reflect.TypeOf((*CfnParameterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnParameterPropsMixin",
		reflect.TypeOf((*CfnParameterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnParameterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnPatchBaselineMixinProps",
		reflect.TypeOf((*CfnPatchBaselineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnPatchBaselinePropsMixin",
		reflect.TypeOf((*CfnPatchBaselinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPatchBaselinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnPatchBaselinePropsMixin.PatchFilterGroupProperty",
		reflect.TypeOf((*CfnPatchBaselinePropsMixin_PatchFilterGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnPatchBaselinePropsMixin.PatchFilterProperty",
		reflect.TypeOf((*CfnPatchBaselinePropsMixin_PatchFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnPatchBaselinePropsMixin.PatchSourceProperty",
		reflect.TypeOf((*CfnPatchBaselinePropsMixin_PatchSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnPatchBaselinePropsMixin.RuleGroupProperty",
		reflect.TypeOf((*CfnPatchBaselinePropsMixin_RuleGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnPatchBaselinePropsMixin.RuleProperty",
		reflect.TypeOf((*CfnPatchBaselinePropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnResourceDataSyncMixinProps",
		reflect.TypeOf((*CfnResourceDataSyncMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnResourceDataSyncPropsMixin",
		reflect.TypeOf((*CfnResourceDataSyncPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourceDataSyncPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnResourceDataSyncPropsMixin.AwsOrganizationsSourceProperty",
		reflect.TypeOf((*CfnResourceDataSyncPropsMixin_AwsOrganizationsSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnResourceDataSyncPropsMixin.S3DestinationProperty",
		reflect.TypeOf((*CfnResourceDataSyncPropsMixin_S3DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnResourceDataSyncPropsMixin.SyncSourceProperty",
		reflect.TypeOf((*CfnResourceDataSyncPropsMixin_SyncSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnResourcePolicyMixinProps",
		reflect.TypeOf((*CfnResourcePolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnResourcePolicyPropsMixin",
		reflect.TypeOf((*CfnResourcePolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourcePolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
