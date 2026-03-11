package awsssm

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnAssociationMixinProps",
		reflect.TypeOf((*CfnAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnAssociationPropsMixin",
		reflect.TypeOf((*CfnAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnAssociationPropsMixin.InstanceAssociationOutputLocationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_InstanceAssociationOutputLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnAssociationPropsMixin.S3OutputLocationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_S3OutputLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnAssociationPropsMixin.TargetProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_TargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnDocumentMixinProps",
		reflect.TypeOf((*CfnDocumentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnDocumentPropsMixin",
		reflect.TypeOf((*CfnDocumentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDocumentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnDocumentPropsMixin.AttachmentsSourceProperty",
		reflect.TypeOf((*CfnDocumentPropsMixin_AttachmentsSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnDocumentPropsMixin.DocumentRequiresProperty",
		reflect.TypeOf((*CfnDocumentPropsMixin_DocumentRequiresProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowMixinProps",
		reflect.TypeOf((*CfnMaintenanceWindowMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowPropsMixin",
		reflect.TypeOf((*CfnMaintenanceWindowPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMaintenanceWindowPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTargetMixinProps",
		reflect.TypeOf((*CfnMaintenanceWindowTargetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTargetPropsMixin",
		reflect.TypeOf((*CfnMaintenanceWindowTargetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMaintenanceWindowTargetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTargetPropsMixin.TargetsProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTargetPropsMixin_TargetsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTaskMixinProps",
		reflect.TypeOf((*CfnMaintenanceWindowTaskMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTaskPropsMixin",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMaintenanceWindowTaskPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTaskPropsMixin.CloudWatchOutputConfigProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_CloudWatchOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTaskPropsMixin.LoggingInfoProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_LoggingInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTaskPropsMixin.MaintenanceWindowAutomationParametersProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_MaintenanceWindowAutomationParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTaskPropsMixin.MaintenanceWindowLambdaParametersProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_MaintenanceWindowLambdaParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTaskPropsMixin.MaintenanceWindowRunCommandParametersProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_MaintenanceWindowRunCommandParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTaskPropsMixin.MaintenanceWindowStepFunctionsParametersProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_MaintenanceWindowStepFunctionsParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTaskPropsMixin.NotificationConfigProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_NotificationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTaskPropsMixin.TargetProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_TargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnMaintenanceWindowTaskPropsMixin.TaskInvocationParametersProperty",
		reflect.TypeOf((*CfnMaintenanceWindowTaskPropsMixin_TaskInvocationParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnParameterMixinProps",
		reflect.TypeOf((*CfnParameterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnParameterPropsMixin",
		reflect.TypeOf((*CfnParameterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnParameterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnPatchBaselineMixinProps",
		reflect.TypeOf((*CfnPatchBaselineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnPatchBaselinePropsMixin",
		reflect.TypeOf((*CfnPatchBaselinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPatchBaselinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnPatchBaselinePropsMixin.PatchFilterGroupProperty",
		reflect.TypeOf((*CfnPatchBaselinePropsMixin_PatchFilterGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnPatchBaselinePropsMixin.PatchFilterProperty",
		reflect.TypeOf((*CfnPatchBaselinePropsMixin_PatchFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnPatchBaselinePropsMixin.PatchSourceProperty",
		reflect.TypeOf((*CfnPatchBaselinePropsMixin_PatchSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnPatchBaselinePropsMixin.RuleGroupProperty",
		reflect.TypeOf((*CfnPatchBaselinePropsMixin_RuleGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnPatchBaselinePropsMixin.RuleProperty",
		reflect.TypeOf((*CfnPatchBaselinePropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnResourceDataSyncMixinProps",
		reflect.TypeOf((*CfnResourceDataSyncMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnResourceDataSyncPropsMixin",
		reflect.TypeOf((*CfnResourceDataSyncPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourceDataSyncPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnResourceDataSyncPropsMixin.AwsOrganizationsSourceProperty",
		reflect.TypeOf((*CfnResourceDataSyncPropsMixin_AwsOrganizationsSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnResourceDataSyncPropsMixin.S3DestinationProperty",
		reflect.TypeOf((*CfnResourceDataSyncPropsMixin_S3DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnResourceDataSyncPropsMixin.SyncSourceProperty",
		reflect.TypeOf((*CfnResourceDataSyncPropsMixin_SyncSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnResourcePolicyMixinProps",
		reflect.TypeOf((*CfnResourcePolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ssm.CfnResourcePolicyPropsMixin",
		reflect.TypeOf((*CfnResourcePolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourcePolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
