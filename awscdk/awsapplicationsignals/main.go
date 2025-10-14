package awsapplicationsignals

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_applicationsignals.CfnDiscovery",
		reflect.TypeOf((*CfnDiscovery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAccountId", GoGetter: "AttrAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "discoveryRef", GoGetter: "DiscoveryRef"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDiscovery{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDiscoveryRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnDiscoveryProps",
		reflect.TypeOf((*CfnDiscoveryProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_applicationsignals.CfnGroupingConfiguration",
		reflect.TypeOf((*CfnGroupingConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAccountId", GoGetter: "AttrAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "groupingAttributeDefinitions", GoGetter: "GroupingAttributeDefinitions"},
			_jsii_.MemberProperty{JsiiProperty: "groupingConfigurationRef", GoGetter: "GroupingConfigurationRef"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGroupingConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGroupingConfigurationRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnGroupingConfiguration.GroupingAttributeDefinitionProperty",
		reflect.TypeOf((*CfnGroupingConfiguration_GroupingAttributeDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnGroupingConfigurationProps",
		reflect.TypeOf((*CfnGroupingConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective",
		reflect.TypeOf((*CfnServiceLevelObjective)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedTime", GoGetter: "AttrCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrEvaluationType", GoGetter: "AttrEvaluationType"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedTime", GoGetter: "AttrLastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "burnRateConfigurations", GoGetter: "BurnRateConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "exclusionWindows", GoGetter: "ExclusionWindows"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "goal", GoGetter: "Goal"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "requestBasedSli", GoGetter: "RequestBasedSli"},
			_jsii_.MemberProperty{JsiiProperty: "serviceLevelObjectiveRef", GoGetter: "ServiceLevelObjectiveRef"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sli", GoGetter: "Sli"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServiceLevelObjective{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IServiceLevelObjectiveRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.BurnRateConfigurationProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_BurnRateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.CalendarIntervalProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_CalendarIntervalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.DependencyConfigProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_DependencyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.DimensionProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_DimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.ExclusionWindowProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_ExclusionWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.GoalProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_GoalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.IntervalProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_IntervalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.MetricDataQueryProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_MetricDataQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.MetricProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_MetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.MetricStatProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_MetricStatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.MonitoredRequestCountMetricProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_MonitoredRequestCountMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.RecurrenceRuleProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_RecurrenceRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.RequestBasedSliMetricProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_RequestBasedSliMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.RequestBasedSliProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_RequestBasedSliProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.RollingIntervalProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_RollingIntervalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.SliMetricProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_SliMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.SliProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_SliProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective.WindowProperty",
		reflect.TypeOf((*CfnServiceLevelObjective_WindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjectiveProps",
		reflect.TypeOf((*CfnServiceLevelObjectiveProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.DiscoveryReference",
		reflect.TypeOf((*DiscoveryReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.GroupingConfigurationReference",
		reflect.TypeOf((*GroupingConfigurationReference)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_applicationsignals.IDiscoveryRef",
		reflect.TypeOf((*IDiscoveryRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "discoveryRef", GoGetter: "DiscoveryRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IDiscoveryRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_applicationsignals.IGroupingConfigurationRef",
		reflect.TypeOf((*IGroupingConfigurationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "groupingConfigurationRef", GoGetter: "GroupingConfigurationRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IGroupingConfigurationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_applicationsignals.IServiceLevelObjectiveRef",
		reflect.TypeOf((*IServiceLevelObjectiveRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceLevelObjectiveRef", GoGetter: "ServiceLevelObjectiveRef"},
		},
		func() interface{} {
			j := jsiiProxy_IServiceLevelObjectiveRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_applicationsignals.ServiceLevelObjectiveReference",
		reflect.TypeOf((*ServiceLevelObjectiveReference)(nil)).Elem(),
	)
}
