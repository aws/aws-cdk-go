package mixinsawscleanrooms

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplateMixinProps",
		reflect.TypeOf((*CfnAnalysisTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplatePropsMixin",
		reflect.TypeOf((*CfnAnalysisTemplatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAnalysisTemplatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplatePropsMixin.AnalysisParameterProperty",
		reflect.TypeOf((*CfnAnalysisTemplatePropsMixin_AnalysisParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplatePropsMixin.AnalysisSchemaProperty",
		reflect.TypeOf((*CfnAnalysisTemplatePropsMixin_AnalysisSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplatePropsMixin.AnalysisSourceMetadataProperty",
		reflect.TypeOf((*CfnAnalysisTemplatePropsMixin_AnalysisSourceMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplatePropsMixin.AnalysisSourceProperty",
		reflect.TypeOf((*CfnAnalysisTemplatePropsMixin_AnalysisSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplatePropsMixin.AnalysisTemplateArtifactMetadataProperty",
		reflect.TypeOf((*CfnAnalysisTemplatePropsMixin_AnalysisTemplateArtifactMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplatePropsMixin.AnalysisTemplateArtifactProperty",
		reflect.TypeOf((*CfnAnalysisTemplatePropsMixin_AnalysisTemplateArtifactProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplatePropsMixin.AnalysisTemplateArtifactsProperty",
		reflect.TypeOf((*CfnAnalysisTemplatePropsMixin_AnalysisTemplateArtifactsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplatePropsMixin.ErrorMessageConfigurationProperty",
		reflect.TypeOf((*CfnAnalysisTemplatePropsMixin_ErrorMessageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplatePropsMixin.HashProperty",
		reflect.TypeOf((*CfnAnalysisTemplatePropsMixin_HashProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnAnalysisTemplatePropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnAnalysisTemplatePropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnCollaborationMixinProps",
		reflect.TypeOf((*CfnCollaborationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnCollaborationPropsMixin",
		reflect.TypeOf((*CfnCollaborationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCollaborationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnCollaborationPropsMixin.DataEncryptionMetadataProperty",
		reflect.TypeOf((*CfnCollaborationPropsMixin_DataEncryptionMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnCollaborationPropsMixin.JobComputePaymentConfigProperty",
		reflect.TypeOf((*CfnCollaborationPropsMixin_JobComputePaymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnCollaborationPropsMixin.MLMemberAbilitiesProperty",
		reflect.TypeOf((*CfnCollaborationPropsMixin_MLMemberAbilitiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnCollaborationPropsMixin.MLPaymentConfigProperty",
		reflect.TypeOf((*CfnCollaborationPropsMixin_MLPaymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnCollaborationPropsMixin.MemberSpecificationProperty",
		reflect.TypeOf((*CfnCollaborationPropsMixin_MemberSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnCollaborationPropsMixin.ModelInferencePaymentConfigProperty",
		reflect.TypeOf((*CfnCollaborationPropsMixin_ModelInferencePaymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnCollaborationPropsMixin.ModelTrainingPaymentConfigProperty",
		reflect.TypeOf((*CfnCollaborationPropsMixin_ModelTrainingPaymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnCollaborationPropsMixin.PaymentConfigurationProperty",
		reflect.TypeOf((*CfnCollaborationPropsMixin_PaymentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnCollaborationPropsMixin.QueryComputePaymentConfigProperty",
		reflect.TypeOf((*CfnCollaborationPropsMixin_QueryComputePaymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTableAssociationMixinProps",
		reflect.TypeOf((*CfnConfiguredTableAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTableAssociationPropsMixin",
		reflect.TypeOf((*CfnConfiguredTableAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfiguredTableAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTableAssociationPropsMixin.ConfiguredTableAssociationAnalysisRuleAggregationProperty",
		reflect.TypeOf((*CfnConfiguredTableAssociationPropsMixin_ConfiguredTableAssociationAnalysisRuleAggregationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTableAssociationPropsMixin.ConfiguredTableAssociationAnalysisRuleCustomProperty",
		reflect.TypeOf((*CfnConfiguredTableAssociationPropsMixin_ConfiguredTableAssociationAnalysisRuleCustomProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTableAssociationPropsMixin.ConfiguredTableAssociationAnalysisRuleListProperty",
		reflect.TypeOf((*CfnConfiguredTableAssociationPropsMixin_ConfiguredTableAssociationAnalysisRuleListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTableAssociationPropsMixin.ConfiguredTableAssociationAnalysisRulePolicyProperty",
		reflect.TypeOf((*CfnConfiguredTableAssociationPropsMixin_ConfiguredTableAssociationAnalysisRulePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTableAssociationPropsMixin.ConfiguredTableAssociationAnalysisRulePolicyV1Property",
		reflect.TypeOf((*CfnConfiguredTableAssociationPropsMixin_ConfiguredTableAssociationAnalysisRulePolicyV1Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTableAssociationPropsMixin.ConfiguredTableAssociationAnalysisRuleProperty",
		reflect.TypeOf((*CfnConfiguredTableAssociationPropsMixin_ConfiguredTableAssociationAnalysisRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTableMixinProps",
		reflect.TypeOf((*CfnConfiguredTableMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfiguredTablePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.AggregateColumnProperty",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_AggregateColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.AggregationConstraintProperty",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_AggregationConstraintProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.AnalysisRuleAggregationProperty",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_AnalysisRuleAggregationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.AnalysisRuleCustomProperty",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_AnalysisRuleCustomProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.AnalysisRuleListProperty",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_AnalysisRuleListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.AnalysisRuleProperty",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_AnalysisRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.AthenaTableReferenceProperty",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_AthenaTableReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.ConfiguredTableAnalysisRulePolicyProperty",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_ConfiguredTableAnalysisRulePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.ConfiguredTableAnalysisRulePolicyV1Property",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_ConfiguredTableAnalysisRulePolicyV1Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.DifferentialPrivacyColumnProperty",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_DifferentialPrivacyColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.DifferentialPrivacyProperty",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_DifferentialPrivacyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.GlueTableReferenceProperty",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_GlueTableReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.SnowflakeTableReferenceProperty",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_SnowflakeTableReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.SnowflakeTableSchemaProperty",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_SnowflakeTableSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.SnowflakeTableSchemaV1Property",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_SnowflakeTableSchemaV1Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnConfiguredTablePropsMixin.TableReferenceProperty",
		reflect.TypeOf((*CfnConfiguredTablePropsMixin_TableReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnIdMappingTableMixinProps",
		reflect.TypeOf((*CfnIdMappingTableMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnIdMappingTablePropsMixin",
		reflect.TypeOf((*CfnIdMappingTablePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdMappingTablePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnIdMappingTablePropsMixin.IdMappingTableInputReferenceConfigProperty",
		reflect.TypeOf((*CfnIdMappingTablePropsMixin_IdMappingTableInputReferenceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnIdMappingTablePropsMixin.IdMappingTableInputReferencePropertiesProperty",
		reflect.TypeOf((*CfnIdMappingTablePropsMixin_IdMappingTableInputReferencePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnIdMappingTablePropsMixin.IdMappingTableInputSourceProperty",
		reflect.TypeOf((*CfnIdMappingTablePropsMixin_IdMappingTableInputSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnIdNamespaceAssociationMixinProps",
		reflect.TypeOf((*CfnIdNamespaceAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnIdNamespaceAssociationPropsMixin",
		reflect.TypeOf((*CfnIdNamespaceAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdNamespaceAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnIdNamespaceAssociationPropsMixin.IdMappingConfigProperty",
		reflect.TypeOf((*CfnIdNamespaceAssociationPropsMixin_IdMappingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnIdNamespaceAssociationPropsMixin.IdNamespaceAssociationInputReferenceConfigProperty",
		reflect.TypeOf((*CfnIdNamespaceAssociationPropsMixin_IdNamespaceAssociationInputReferenceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnIdNamespaceAssociationPropsMixin.IdNamespaceAssociationInputReferencePropertiesProperty",
		reflect.TypeOf((*CfnIdNamespaceAssociationPropsMixin_IdNamespaceAssociationInputReferencePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipMixinProps",
		reflect.TypeOf((*CfnMembershipMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin",
		reflect.TypeOf((*CfnMembershipPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMembershipPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin.MembershipJobComputePaymentConfigProperty",
		reflect.TypeOf((*CfnMembershipPropsMixin_MembershipJobComputePaymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin.MembershipMLPaymentConfigProperty",
		reflect.TypeOf((*CfnMembershipPropsMixin_MembershipMLPaymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin.MembershipModelInferencePaymentConfigProperty",
		reflect.TypeOf((*CfnMembershipPropsMixin_MembershipModelInferencePaymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin.MembershipModelTrainingPaymentConfigProperty",
		reflect.TypeOf((*CfnMembershipPropsMixin_MembershipModelTrainingPaymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin.MembershipPaymentConfigurationProperty",
		reflect.TypeOf((*CfnMembershipPropsMixin_MembershipPaymentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin.MembershipProtectedJobOutputConfigurationProperty",
		reflect.TypeOf((*CfnMembershipPropsMixin_MembershipProtectedJobOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin.MembershipProtectedJobResultConfigurationProperty",
		reflect.TypeOf((*CfnMembershipPropsMixin_MembershipProtectedJobResultConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin.MembershipProtectedQueryOutputConfigurationProperty",
		reflect.TypeOf((*CfnMembershipPropsMixin_MembershipProtectedQueryOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin.MembershipProtectedQueryResultConfigurationProperty",
		reflect.TypeOf((*CfnMembershipPropsMixin_MembershipProtectedQueryResultConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin.MembershipQueryComputePaymentConfigProperty",
		reflect.TypeOf((*CfnMembershipPropsMixin_MembershipQueryComputePaymentConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin.ProtectedJobS3OutputConfigurationInputProperty",
		reflect.TypeOf((*CfnMembershipPropsMixin_ProtectedJobS3OutputConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipPropsMixin.ProtectedQueryS3OutputConfigurationProperty",
		reflect.TypeOf((*CfnMembershipPropsMixin_ProtectedQueryS3OutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnPrivacyBudgetTemplateMixinProps",
		reflect.TypeOf((*CfnPrivacyBudgetTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnPrivacyBudgetTemplatePropsMixin",
		reflect.TypeOf((*CfnPrivacyBudgetTemplatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPrivacyBudgetTemplatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnPrivacyBudgetTemplatePropsMixin.BudgetParameterProperty",
		reflect.TypeOf((*CfnPrivacyBudgetTemplatePropsMixin_BudgetParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnPrivacyBudgetTemplatePropsMixin.ParametersProperty",
		reflect.TypeOf((*CfnPrivacyBudgetTemplatePropsMixin_ParametersProperty)(nil)).Elem(),
	)
}
