package previewawsarcregionswitchmixins


// Properties for CfnPlanPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var stepProperty_ StepProperty
//
//   cfnPlanMixinProps := &CfnPlanMixinProps{
//   	AssociatedAlarms: map[string]interface{}{
//   		"associatedAlarmsKey": &AssociatedAlarmProperty{
//   			"alarmType": jsii.String("alarmType"),
//   			"crossAccountRole": jsii.String("crossAccountRole"),
//   			"externalId": jsii.String("externalId"),
//   			"resourceIdentifier": jsii.String("resourceIdentifier"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ExecutionRole: jsii.String("executionRole"),
//   	Name: jsii.String("name"),
//   	PrimaryRegion: jsii.String("primaryRegion"),
//   	RecoveryApproach: jsii.String("recoveryApproach"),
//   	RecoveryTimeObjectiveMinutes: jsii.Number(123),
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Triggers: []interface{}{
//   		&TriggerProperty{
//   			Action: jsii.String("action"),
//   			Conditions: []interface{}{
//   				&TriggerConditionProperty{
//   					AssociatedAlarmName: jsii.String("associatedAlarmName"),
//   					Condition: jsii.String("condition"),
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			MinDelayMinutesBetweenExecutions: jsii.Number(123),
//   			TargetRegion: jsii.String("targetRegion"),
//   		},
//   	},
//   	Workflows: []interface{}{
//   		&WorkflowProperty{
//   			Steps: []interface{}{
//   				&StepProperty{
//   					Description: jsii.String("description"),
//   					ExecutionBlockConfiguration: &ExecutionBlockConfigurationProperty{
//   						ArcRoutingControlConfig: &ArcRoutingControlConfigurationProperty{
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							ExternalId: jsii.String("externalId"),
//   							RegionAndRoutingControls: map[string]interface{}{
//   								"regionAndRoutingControlsKey": []interface{}{
//   									&ArcRoutingControlStateProperty{
//   										"routingControlArn": jsii.String("routingControlArn"),
//   										"state": jsii.String("state"),
//   									},
//   								},
//   							},
//   							TimeoutMinutes: jsii.Number(123),
//   						},
//   						CustomActionLambdaConfig: &CustomActionLambdaConfigurationProperty{
//   							Lambdas: []interface{}{
//   								&LambdasProperty{
//   									Arn: jsii.String("arn"),
//   									CrossAccountRole: jsii.String("crossAccountRole"),
//   									ExternalId: jsii.String("externalId"),
//   								},
//   							},
//   							RegionToRun: jsii.String("regionToRun"),
//   							RetryIntervalMinutes: jsii.Number(123),
//   							TimeoutMinutes: jsii.Number(123),
//   							Ungraceful: &LambdaUngracefulProperty{
//   								Behavior: jsii.String("behavior"),
//   							},
//   						},
//   						Ec2AsgCapacityIncreaseConfig: &Ec2AsgCapacityIncreaseConfigurationProperty{
//   							Asgs: []interface{}{
//   								&AsgProperty{
//   									Arn: jsii.String("arn"),
//   									CrossAccountRole: jsii.String("crossAccountRole"),
//   									ExternalId: jsii.String("externalId"),
//   								},
//   							},
//   							CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   							TargetPercent: jsii.Number(123),
//   							TimeoutMinutes: jsii.Number(123),
//   							Ungraceful: &Ec2UngracefulProperty{
//   								MinimumSuccessPercentage: jsii.Number(123),
//   							},
//   						},
//   						EcsCapacityIncreaseConfig: &EcsCapacityIncreaseConfigurationProperty{
//   							CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   							Services: []interface{}{
//   								&ServiceProperty{
//   									ClusterArn: jsii.String("clusterArn"),
//   									CrossAccountRole: jsii.String("crossAccountRole"),
//   									ExternalId: jsii.String("externalId"),
//   									ServiceArn: jsii.String("serviceArn"),
//   								},
//   							},
//   							TargetPercent: jsii.Number(123),
//   							TimeoutMinutes: jsii.Number(123),
//   							Ungraceful: &EcsUngracefulProperty{
//   								MinimumSuccessPercentage: jsii.Number(123),
//   							},
//   						},
//   						EksResourceScalingConfig: &EksResourceScalingConfigurationProperty{
//   							CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   							EksClusters: []interface{}{
//   								&EksClusterProperty{
//   									ClusterArn: jsii.String("clusterArn"),
//   									CrossAccountRole: jsii.String("crossAccountRole"),
//   									ExternalId: jsii.String("externalId"),
//   								},
//   							},
//   							KubernetesResourceType: &KubernetesResourceTypeProperty{
//   								ApiVersion: jsii.String("apiVersion"),
//   								Kind: jsii.String("kind"),
//   							},
//   							ScalingResources: []interface{}{
//   								map[string]interface{}{
//   									"scalingResourcesKey": map[string]interface{}{
//   										"scalingResourcesKey": &KubernetesScalingResourceProperty{
//   											"hpaName": jsii.String("hpaName"),
//   											"name": jsii.String("name"),
//   											"namespace": jsii.String("namespace"),
//   										},
//   									},
//   								},
//   							},
//   							TargetPercent: jsii.Number(123),
//   							TimeoutMinutes: jsii.Number(123),
//   							Ungraceful: &EksResourceScalingUngracefulProperty{
//   								MinimumSuccessPercentage: jsii.Number(123),
//   							},
//   						},
//   						ExecutionApprovalConfig: &ExecutionApprovalConfigurationProperty{
//   							ApprovalRole: jsii.String("approvalRole"),
//   							TimeoutMinutes: jsii.Number(123),
//   						},
//   						GlobalAuroraConfig: &GlobalAuroraConfigurationProperty{
//   							Behavior: jsii.String("behavior"),
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							DatabaseClusterArns: []*string{
//   								jsii.String("databaseClusterArns"),
//   							},
//   							ExternalId: jsii.String("externalId"),
//   							GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   							TimeoutMinutes: jsii.Number(123),
//   							Ungraceful: &GlobalAuroraUngracefulProperty{
//   								Ungraceful: jsii.String("ungraceful"),
//   							},
//   						},
//   						ParallelConfig: &ParallelExecutionBlockConfigurationProperty{
//   							Steps: []interface{}{
//   								stepProperty_,
//   							},
//   						},
//   						RegionSwitchPlanConfig: &RegionSwitchPlanConfigurationProperty{
//   							Arn: jsii.String("arn"),
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							ExternalId: jsii.String("externalId"),
//   						},
//   						Route53HealthCheckConfig: &Route53HealthCheckConfigurationProperty{
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							ExternalId: jsii.String("externalId"),
//   							HostedZoneId: jsii.String("hostedZoneId"),
//   							RecordName: jsii.String("recordName"),
//   							RecordSets: []interface{}{
//   								&Route53ResourceRecordSetProperty{
//   									RecordSetIdentifier: jsii.String("recordSetIdentifier"),
//   									Region: jsii.String("region"),
//   								},
//   							},
//   							TimeoutMinutes: jsii.Number(123),
//   						},
//   					},
//   					ExecutionBlockType: jsii.String("executionBlockType"),
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			WorkflowDescription: jsii.String("workflowDescription"),
//   			WorkflowTargetAction: jsii.String("workflowTargetAction"),
//   			WorkflowTargetRegion: jsii.String("workflowTargetRegion"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html
//
type CfnPlanMixinProps struct {
	// The associated application health alarms for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-associatedalarms
	//
	AssociatedAlarms interface{} `field:"optional" json:"associatedAlarms" yaml:"associatedAlarms"`
	// The description for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The execution role for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-executionrole
	//
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The primary Region for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-primaryregion
	//
	PrimaryRegion *string `field:"optional" json:"primaryRegion" yaml:"primaryRegion"`
	// The recovery approach for a Region switch plan, which can be active/active (activeActive) or active/passive (activePassive).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-recoveryapproach
	//
	RecoveryApproach *string `field:"optional" json:"recoveryApproach" yaml:"recoveryApproach"`
	// The recovery time objective for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-recoverytimeobjectiveminutes
	//
	RecoveryTimeObjectiveMinutes *float64 `field:"optional" json:"recoveryTimeObjectiveMinutes" yaml:"recoveryTimeObjectiveMinutes"`
	// The AWS Regions for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The triggers for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-triggers
	//
	Triggers interface{} `field:"optional" json:"triggers" yaml:"triggers"`
	// The workflows for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-workflows
	//
	Workflows interface{} `field:"optional" json:"workflows" yaml:"workflows"`
}

