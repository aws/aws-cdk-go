package awsbackup

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupPlanMixinProps",
		reflect.TypeOf((*CfnBackupPlanMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupPlanPropsMixin",
		reflect.TypeOf((*CfnBackupPlanPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBackupPlanPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupPlanPropsMixin.AdvancedBackupSettingResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_AdvancedBackupSettingResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupPlanPropsMixin.BackupPlanResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_BackupPlanResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupPlanPropsMixin.BackupRuleResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_BackupRuleResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupPlanPropsMixin.CopyActionResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_CopyActionResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupPlanPropsMixin.IndexActionsResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_IndexActionsResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupPlanPropsMixin.LifecycleResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_LifecycleResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupPlanPropsMixin.ScanActionResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_ScanActionResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupPlanPropsMixin.ScanSettingResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_ScanSettingResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupSelectionMixinProps",
		reflect.TypeOf((*CfnBackupSelectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupSelectionPropsMixin",
		reflect.TypeOf((*CfnBackupSelectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBackupSelectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupSelectionPropsMixin.BackupSelectionResourceTypeProperty",
		reflect.TypeOf((*CfnBackupSelectionPropsMixin_BackupSelectionResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupSelectionPropsMixin.ConditionResourceTypeProperty",
		reflect.TypeOf((*CfnBackupSelectionPropsMixin_ConditionResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupVaultMixinProps",
		reflect.TypeOf((*CfnBackupVaultMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupVaultPropsMixin",
		reflect.TypeOf((*CfnBackupVaultPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBackupVaultPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupVaultPropsMixin.LockConfigurationTypeProperty",
		reflect.TypeOf((*CfnBackupVaultPropsMixin_LockConfigurationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnBackupVaultPropsMixin.NotificationObjectTypeProperty",
		reflect.TypeOf((*CfnBackupVaultPropsMixin_NotificationObjectTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnFrameworkMixinProps",
		reflect.TypeOf((*CfnFrameworkMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnFrameworkPropsMixin",
		reflect.TypeOf((*CfnFrameworkPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFrameworkPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnFrameworkPropsMixin.ControlInputParameterProperty",
		reflect.TypeOf((*CfnFrameworkPropsMixin_ControlInputParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnFrameworkPropsMixin.FrameworkControlProperty",
		reflect.TypeOf((*CfnFrameworkPropsMixin_FrameworkControlProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnLogicallyAirGappedBackupVaultMixinProps",
		reflect.TypeOf((*CfnLogicallyAirGappedBackupVaultMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnLogicallyAirGappedBackupVaultPropsMixin",
		reflect.TypeOf((*CfnLogicallyAirGappedBackupVaultPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLogicallyAirGappedBackupVaultPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnLogicallyAirGappedBackupVaultPropsMixin.NotificationObjectTypeProperty",
		reflect.TypeOf((*CfnLogicallyAirGappedBackupVaultPropsMixin_NotificationObjectTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnReportPlanMixinProps",
		reflect.TypeOf((*CfnReportPlanMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnReportPlanPropsMixin",
		reflect.TypeOf((*CfnReportPlanPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReportPlanPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnReportPlanPropsMixin.ReportDeliveryChannelProperty",
		reflect.TypeOf((*CfnReportPlanPropsMixin_ReportDeliveryChannelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnReportPlanPropsMixin.ReportSettingProperty",
		reflect.TypeOf((*CfnReportPlanPropsMixin_ReportSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnRestoreTestingPlanMixinProps",
		reflect.TypeOf((*CfnRestoreTestingPlanMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnRestoreTestingPlanPropsMixin",
		reflect.TypeOf((*CfnRestoreTestingPlanPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRestoreTestingPlanPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnRestoreTestingPlanPropsMixin.RestoreTestingRecoveryPointSelectionProperty",
		reflect.TypeOf((*CfnRestoreTestingPlanPropsMixin_RestoreTestingRecoveryPointSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnRestoreTestingSelectionMixinProps",
		reflect.TypeOf((*CfnRestoreTestingSelectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnRestoreTestingSelectionPropsMixin",
		reflect.TypeOf((*CfnRestoreTestingSelectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRestoreTestingSelectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnRestoreTestingSelectionPropsMixin.KeyValueProperty",
		reflect.TypeOf((*CfnRestoreTestingSelectionPropsMixin_KeyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnRestoreTestingSelectionPropsMixin.ProtectedResourceConditionsProperty",
		reflect.TypeOf((*CfnRestoreTestingSelectionPropsMixin_ProtectedResourceConditionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnTieringConfigurationMixinProps",
		reflect.TypeOf((*CfnTieringConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnTieringConfigurationPropsMixin",
		reflect.TypeOf((*CfnTieringConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTieringConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnTieringConfigurationPropsMixin.ResourceSelectionProperty",
		reflect.TypeOf((*CfnTieringConfigurationPropsMixin_ResourceSelectionProperty)(nil)).Elem(),
	)
}
