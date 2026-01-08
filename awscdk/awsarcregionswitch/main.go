package awsarcregionswitch

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan",
		reflect.TypeOf((*CfnPlan)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "associatedAlarms", GoGetter: "AssociatedAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrHealthChecksForPlan", GoGetter: "AttrHealthChecksForPlan"},
			_jsii_.MemberProperty{JsiiProperty: "attrOwner", GoGetter: "AttrOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrPlanHealthChecks", GoGetter: "AttrPlanHealthChecks"},
			_jsii_.MemberProperty{JsiiProperty: "attrRoute53HealthChecks", GoGetter: "AttrRoute53HealthChecks"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersion", GoGetter: "AttrVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "planRef", GoGetter: "PlanRef"},
			_jsii_.MemberProperty{JsiiProperty: "primaryRegion", GoGetter: "PrimaryRegion"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryApproach", GoGetter: "RecoveryApproach"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryTimeObjectiveMinutes", GoGetter: "RecoveryTimeObjectiveMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberProperty{JsiiProperty: "regions", GoGetter: "Regions"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "reportConfiguration", GoGetter: "ReportConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggers", GoGetter: "Triggers"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "workflows", GoGetter: "Workflows"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPlan{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsarcregionswitchIPlanRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.ArcRoutingControlConfigurationProperty",
		reflect.TypeOf((*CfnPlan_ArcRoutingControlConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.ArcRoutingControlStateProperty",
		reflect.TypeOf((*CfnPlan_ArcRoutingControlStateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.AsgProperty",
		reflect.TypeOf((*CfnPlan_AsgProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.AssociatedAlarmProperty",
		reflect.TypeOf((*CfnPlan_AssociatedAlarmProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.CustomActionLambdaConfigurationProperty",
		reflect.TypeOf((*CfnPlan_CustomActionLambdaConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.DocumentDbConfigurationProperty",
		reflect.TypeOf((*CfnPlan_DocumentDbConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.DocumentDbUngracefulProperty",
		reflect.TypeOf((*CfnPlan_DocumentDbUngracefulProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.Ec2AsgCapacityIncreaseConfigurationProperty",
		reflect.TypeOf((*CfnPlan_Ec2AsgCapacityIncreaseConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.Ec2UngracefulProperty",
		reflect.TypeOf((*CfnPlan_Ec2UngracefulProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.EcsCapacityIncreaseConfigurationProperty",
		reflect.TypeOf((*CfnPlan_EcsCapacityIncreaseConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.EcsUngracefulProperty",
		reflect.TypeOf((*CfnPlan_EcsUngracefulProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.EksClusterProperty",
		reflect.TypeOf((*CfnPlan_EksClusterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.EksResourceScalingConfigurationProperty",
		reflect.TypeOf((*CfnPlan_EksResourceScalingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.EksResourceScalingUngracefulProperty",
		reflect.TypeOf((*CfnPlan_EksResourceScalingUngracefulProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.ExecutionApprovalConfigurationProperty",
		reflect.TypeOf((*CfnPlan_ExecutionApprovalConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.ExecutionBlockConfigurationProperty",
		reflect.TypeOf((*CfnPlan_ExecutionBlockConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.GlobalAuroraConfigurationProperty",
		reflect.TypeOf((*CfnPlan_GlobalAuroraConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.GlobalAuroraUngracefulProperty",
		reflect.TypeOf((*CfnPlan_GlobalAuroraUngracefulProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.HealthCheckStateProperty",
		reflect.TypeOf((*CfnPlan_HealthCheckStateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.KubernetesResourceTypeProperty",
		reflect.TypeOf((*CfnPlan_KubernetesResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.KubernetesScalingResourceProperty",
		reflect.TypeOf((*CfnPlan_KubernetesScalingResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.LambdaUngracefulProperty",
		reflect.TypeOf((*CfnPlan_LambdaUngracefulProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.LambdasProperty",
		reflect.TypeOf((*CfnPlan_LambdasProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.ParallelExecutionBlockConfigurationProperty",
		reflect.TypeOf((*CfnPlan_ParallelExecutionBlockConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.RegionSwitchPlanConfigurationProperty",
		reflect.TypeOf((*CfnPlan_RegionSwitchPlanConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.ReportConfigurationProperty",
		reflect.TypeOf((*CfnPlan_ReportConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.ReportOutputConfigurationProperty",
		reflect.TypeOf((*CfnPlan_ReportOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.Route53HealthCheckConfigurationProperty",
		reflect.TypeOf((*CfnPlan_Route53HealthCheckConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.Route53HealthChecksProperty",
		reflect.TypeOf((*CfnPlan_Route53HealthChecksProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.Route53ResourceRecordSetProperty",
		reflect.TypeOf((*CfnPlan_Route53ResourceRecordSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.S3ReportOutputConfigurationProperty",
		reflect.TypeOf((*CfnPlan_S3ReportOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.ServiceProperty",
		reflect.TypeOf((*CfnPlan_ServiceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.StepProperty",
		reflect.TypeOf((*CfnPlan_StepProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.TriggerConditionProperty",
		reflect.TypeOf((*CfnPlan_TriggerConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.TriggerProperty",
		reflect.TypeOf((*CfnPlan_TriggerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlan.WorkflowProperty",
		reflect.TypeOf((*CfnPlan_WorkflowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_arcregionswitch.CfnPlanProps",
		reflect.TypeOf((*CfnPlanProps)(nil)).Elem(),
	)
}
