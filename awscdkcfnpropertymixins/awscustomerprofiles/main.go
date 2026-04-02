package awscustomerprofiles

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnCalculatedAttributeDefinitionMixinProps",
		reflect.TypeOf((*CfnCalculatedAttributeDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnCalculatedAttributeDefinitionPropsMixin",
		reflect.TypeOf((*CfnCalculatedAttributeDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCalculatedAttributeDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnCalculatedAttributeDefinitionPropsMixin.AttributeDetailsProperty",
		reflect.TypeOf((*CfnCalculatedAttributeDefinitionPropsMixin_AttributeDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnCalculatedAttributeDefinitionPropsMixin.AttributeItemProperty",
		reflect.TypeOf((*CfnCalculatedAttributeDefinitionPropsMixin_AttributeItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnCalculatedAttributeDefinitionPropsMixin.ConditionsProperty",
		reflect.TypeOf((*CfnCalculatedAttributeDefinitionPropsMixin_ConditionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnCalculatedAttributeDefinitionPropsMixin.RangeProperty",
		reflect.TypeOf((*CfnCalculatedAttributeDefinitionPropsMixin_RangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnCalculatedAttributeDefinitionPropsMixin.ReadinessProperty",
		reflect.TypeOf((*CfnCalculatedAttributeDefinitionPropsMixin_ReadinessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnCalculatedAttributeDefinitionPropsMixin.ThresholdProperty",
		reflect.TypeOf((*CfnCalculatedAttributeDefinitionPropsMixin_ThresholdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnCalculatedAttributeDefinitionPropsMixin.ValueRangeProperty",
		reflect.TypeOf((*CfnCalculatedAttributeDefinitionPropsMixin_ValueRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnDomainMixinProps",
		reflect.TypeOf((*CfnDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnDomainPropsMixin",
		reflect.TypeOf((*CfnDomainPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnDomainPropsMixin.AttributeTypesSelectorProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_AttributeTypesSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnDomainPropsMixin.AutoMergingProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_AutoMergingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnDomainPropsMixin.ConflictResolutionProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ConflictResolutionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnDomainPropsMixin.ConsolidationProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ConsolidationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnDomainPropsMixin.DataStoreProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DataStoreProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnDomainPropsMixin.DomainStatsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DomainStatsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnDomainPropsMixin.ExportingConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ExportingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnDomainPropsMixin.JobScheduleProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_JobScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnDomainPropsMixin.MatchingProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_MatchingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnDomainPropsMixin.MatchingRuleProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_MatchingRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnDomainPropsMixin.ReadinessProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ReadinessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnDomainPropsMixin.RuleBasedMatchingProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_RuleBasedMatchingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnDomainPropsMixin.S3ExportingConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_S3ExportingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnEventStreamMixinProps",
		reflect.TypeOf((*CfnEventStreamMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnEventStreamPropsMixin",
		reflect.TypeOf((*CfnEventStreamPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventStreamPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnEventStreamPropsMixin.DestinationDetailsProperty",
		reflect.TypeOf((*CfnEventStreamPropsMixin_DestinationDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnEventTriggerMixinProps",
		reflect.TypeOf((*CfnEventTriggerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnEventTriggerPropsMixin",
		reflect.TypeOf((*CfnEventTriggerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventTriggerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnEventTriggerPropsMixin.EventTriggerConditionProperty",
		reflect.TypeOf((*CfnEventTriggerPropsMixin_EventTriggerConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnEventTriggerPropsMixin.EventTriggerDimensionProperty",
		reflect.TypeOf((*CfnEventTriggerPropsMixin_EventTriggerDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnEventTriggerPropsMixin.EventTriggerLimitsProperty",
		reflect.TypeOf((*CfnEventTriggerPropsMixin_EventTriggerLimitsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnEventTriggerPropsMixin.ObjectAttributeProperty",
		reflect.TypeOf((*CfnEventTriggerPropsMixin_ObjectAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnEventTriggerPropsMixin.PeriodProperty",
		reflect.TypeOf((*CfnEventTriggerPropsMixin_PeriodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationMixinProps",
		reflect.TypeOf((*CfnIntegrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin",
		reflect.TypeOf((*CfnIntegrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIntegrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.ConnectorOperatorProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_ConnectorOperatorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.FlowDefinitionProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_FlowDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.IncrementalPullConfigProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_IncrementalPullConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.MarketoSourcePropertiesProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_MarketoSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.ObjectTypeMappingProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_ObjectTypeMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.S3SourcePropertiesProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_S3SourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.SalesforceSourcePropertiesProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_SalesforceSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.ScheduledTriggerPropertiesProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_ScheduledTriggerPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.ServiceNowSourcePropertiesProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_ServiceNowSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.SourceConnectorPropertiesProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_SourceConnectorPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.SourceFlowConfigProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_SourceFlowConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.TaskPropertiesMapProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_TaskPropertiesMapProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.TaskProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_TaskProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.TriggerConfigProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_TriggerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.TriggerPropertiesProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_TriggerPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnIntegrationPropsMixin.ZendeskSourcePropertiesProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_ZendeskSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnObjectTypeMixinProps",
		reflect.TypeOf((*CfnObjectTypeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnObjectTypePropsMixin",
		reflect.TypeOf((*CfnObjectTypePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnObjectTypePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnObjectTypePropsMixin.FieldMapProperty",
		reflect.TypeOf((*CfnObjectTypePropsMixin_FieldMapProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnObjectTypePropsMixin.KeyMapProperty",
		reflect.TypeOf((*CfnObjectTypePropsMixin_KeyMapProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnObjectTypePropsMixin.ObjectTypeFieldProperty",
		reflect.TypeOf((*CfnObjectTypePropsMixin_ObjectTypeFieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnObjectTypePropsMixin.ObjectTypeKeyProperty",
		reflect.TypeOf((*CfnObjectTypePropsMixin_ObjectTypeKeyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnRecommenderMixinProps",
		reflect.TypeOf((*CfnRecommenderMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnRecommenderPropsMixin",
		reflect.TypeOf((*CfnRecommenderPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRecommenderPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnRecommenderPropsMixin.EventParametersProperty",
		reflect.TypeOf((*CfnRecommenderPropsMixin_EventParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnRecommenderPropsMixin.EventsConfigProperty",
		reflect.TypeOf((*CfnRecommenderPropsMixin_EventsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnRecommenderPropsMixin.MetricsProperty",
		reflect.TypeOf((*CfnRecommenderPropsMixin_MetricsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnRecommenderPropsMixin.RecommenderConfigProperty",
		reflect.TypeOf((*CfnRecommenderPropsMixin_RecommenderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnRecommenderPropsMixin.RecommenderUpdateProperty",
		reflect.TypeOf((*CfnRecommenderPropsMixin_RecommenderUpdateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnRecommenderPropsMixin.TrainingMetricsProperty",
		reflect.TypeOf((*CfnRecommenderPropsMixin_TrainingMetricsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionMixinProps",
		reflect.TypeOf((*CfnSegmentDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionPropsMixin",
		reflect.TypeOf((*CfnSegmentDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSegmentDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionPropsMixin.AddressDimensionProperty",
		reflect.TypeOf((*CfnSegmentDefinitionPropsMixin_AddressDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionPropsMixin.AttributeDimensionProperty",
		reflect.TypeOf((*CfnSegmentDefinitionPropsMixin_AttributeDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionPropsMixin.CalculatedAttributeDimensionProperty",
		reflect.TypeOf((*CfnSegmentDefinitionPropsMixin_CalculatedAttributeDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionPropsMixin.ConditionOverridesProperty",
		reflect.TypeOf((*CfnSegmentDefinitionPropsMixin_ConditionOverridesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionPropsMixin.DateDimensionProperty",
		reflect.TypeOf((*CfnSegmentDefinitionPropsMixin_DateDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionPropsMixin.DimensionProperty",
		reflect.TypeOf((*CfnSegmentDefinitionPropsMixin_DimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionPropsMixin.ExtraLengthValueProfileDimensionProperty",
		reflect.TypeOf((*CfnSegmentDefinitionPropsMixin_ExtraLengthValueProfileDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionPropsMixin.GroupProperty",
		reflect.TypeOf((*CfnSegmentDefinitionPropsMixin_GroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionPropsMixin.ProfileAttributesProperty",
		reflect.TypeOf((*CfnSegmentDefinitionPropsMixin_ProfileAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionPropsMixin.ProfileDimensionProperty",
		reflect.TypeOf((*CfnSegmentDefinitionPropsMixin_ProfileDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionPropsMixin.ProfileTypeDimensionProperty",
		reflect.TypeOf((*CfnSegmentDefinitionPropsMixin_ProfileTypeDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionPropsMixin.RangeOverrideProperty",
		reflect.TypeOf((*CfnSegmentDefinitionPropsMixin_RangeOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionPropsMixin.SegmentGroupProperty",
		reflect.TypeOf((*CfnSegmentDefinitionPropsMixin_SegmentGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnSegmentDefinitionPropsMixin.SourceSegmentProperty",
		reflect.TypeOf((*CfnSegmentDefinitionPropsMixin_SourceSegmentProperty)(nil)).Elem(),
	)
}
