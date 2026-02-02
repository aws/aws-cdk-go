package previewawsbackupmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupPlanMixinProps",
		reflect.TypeOf((*CfnBackupPlanMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupPlanPropsMixin",
		reflect.TypeOf((*CfnBackupPlanPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBackupPlanPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupPlanPropsMixin.AdvancedBackupSettingResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_AdvancedBackupSettingResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupPlanPropsMixin.BackupPlanResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_BackupPlanResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupPlanPropsMixin.BackupRuleResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_BackupRuleResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupPlanPropsMixin.CopyActionResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_CopyActionResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupPlanPropsMixin.IndexActionsResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_IndexActionsResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupPlanPropsMixin.LifecycleResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_LifecycleResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupPlanPropsMixin.ScanActionResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_ScanActionResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupPlanPropsMixin.ScanSettingResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlanPropsMixin_ScanSettingResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupSelectionMixinProps",
		reflect.TypeOf((*CfnBackupSelectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupSelectionPropsMixin",
		reflect.TypeOf((*CfnBackupSelectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBackupSelectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupSelectionPropsMixin.BackupSelectionResourceTypeProperty",
		reflect.TypeOf((*CfnBackupSelectionPropsMixin_BackupSelectionResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupSelectionPropsMixin.ConditionResourceTypeProperty",
		reflect.TypeOf((*CfnBackupSelectionPropsMixin_ConditionResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupVaultMixinProps",
		reflect.TypeOf((*CfnBackupVaultMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupVaultPropsMixin",
		reflect.TypeOf((*CfnBackupVaultPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBackupVaultPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupVaultPropsMixin.LockConfigurationTypeProperty",
		reflect.TypeOf((*CfnBackupVaultPropsMixin_LockConfigurationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnBackupVaultPropsMixin.NotificationObjectTypeProperty",
		reflect.TypeOf((*CfnBackupVaultPropsMixin_NotificationObjectTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnFrameworkMixinProps",
		reflect.TypeOf((*CfnFrameworkMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnFrameworkPropsMixin",
		reflect.TypeOf((*CfnFrameworkPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFrameworkPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnFrameworkPropsMixin.ControlInputParameterProperty",
		reflect.TypeOf((*CfnFrameworkPropsMixin_ControlInputParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnFrameworkPropsMixin.FrameworkControlProperty",
		reflect.TypeOf((*CfnFrameworkPropsMixin_FrameworkControlProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnLogicallyAirGappedBackupVaultMixinProps",
		reflect.TypeOf((*CfnLogicallyAirGappedBackupVaultMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnLogicallyAirGappedBackupVaultPropsMixin",
		reflect.TypeOf((*CfnLogicallyAirGappedBackupVaultPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLogicallyAirGappedBackupVaultPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnLogicallyAirGappedBackupVaultPropsMixin.NotificationObjectTypeProperty",
		reflect.TypeOf((*CfnLogicallyAirGappedBackupVaultPropsMixin_NotificationObjectTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnReportPlanMixinProps",
		reflect.TypeOf((*CfnReportPlanMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnReportPlanPropsMixin",
		reflect.TypeOf((*CfnReportPlanPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReportPlanPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnReportPlanPropsMixin.ReportDeliveryChannelProperty",
		reflect.TypeOf((*CfnReportPlanPropsMixin_ReportDeliveryChannelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnReportPlanPropsMixin.ReportSettingProperty",
		reflect.TypeOf((*CfnReportPlanPropsMixin_ReportSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnRestoreTestingPlanMixinProps",
		reflect.TypeOf((*CfnRestoreTestingPlanMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnRestoreTestingPlanPropsMixin",
		reflect.TypeOf((*CfnRestoreTestingPlanPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRestoreTestingPlanPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnRestoreTestingPlanPropsMixin.RestoreTestingRecoveryPointSelectionProperty",
		reflect.TypeOf((*CfnRestoreTestingPlanPropsMixin_RestoreTestingRecoveryPointSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnRestoreTestingSelectionMixinProps",
		reflect.TypeOf((*CfnRestoreTestingSelectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnRestoreTestingSelectionPropsMixin",
		reflect.TypeOf((*CfnRestoreTestingSelectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRestoreTestingSelectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnRestoreTestingSelectionPropsMixin.KeyValueProperty",
		reflect.TypeOf((*CfnRestoreTestingSelectionPropsMixin_KeyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnRestoreTestingSelectionPropsMixin.ProtectedResourceConditionsProperty",
		reflect.TypeOf((*CfnRestoreTestingSelectionPropsMixin_ProtectedResourceConditionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnTieringConfigurationMixinProps",
		reflect.TypeOf((*CfnTieringConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnTieringConfigurationPropsMixin",
		reflect.TypeOf((*CfnTieringConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTieringConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnTieringConfigurationPropsMixin.ResourceSelectionProperty",
		reflect.TypeOf((*CfnTieringConfigurationPropsMixin_ResourceSelectionProperty)(nil)).Elem(),
	)
}
