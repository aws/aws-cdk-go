package awsarcregionswitch


// Properties for defining a `CfnPlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stepProperty_ StepProperty
//
//   cfnPlanProps := &CfnPlanProps{
//   	ExecutionRole: jsii.String("executionRole"),
//   	Name: jsii.String("name"),
//   	RecoveryApproach: jsii.String("recoveryApproach"),
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	Workflows: []interface{}{
//   		&WorkflowProperty{
//   			WorkflowTargetAction: jsii.String("workflowTargetAction"),
//
//   			// the properties below are optional
//   			Steps: []interface{}{
//   				&StepProperty{
//   					ExecutionBlockConfiguration: &ExecutionBlockConfigurationProperty{
//   						ArcRoutingControlConfig: &ArcRoutingControlConfigurationProperty{
//   							RegionAndRoutingControls: map[string]interface{}{
//   								"regionAndRoutingControlsKey": []interface{}{
//   									&ArcRoutingControlStateProperty{
//   										"routingControlArn": jsii.String("routingControlArn"),
//   										"state": jsii.String("state"),
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							ExternalId: jsii.String("externalId"),
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
//
//   							// the properties below are optional
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
//
//   							// the properties below are optional
//   							CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   							TargetPercent: jsii.Number(123),
//   							TimeoutMinutes: jsii.Number(123),
//   							Ungraceful: &Ec2UngracefulProperty{
//   								MinimumSuccessPercentage: jsii.Number(123),
//   							},
//   						},
//   						EcsCapacityIncreaseConfig: &EcsCapacityIncreaseConfigurationProperty{
//   							Services: []interface{}{
//   								&ServiceProperty{
//   									ClusterArn: jsii.String("clusterArn"),
//   									CrossAccountRole: jsii.String("crossAccountRole"),
//   									ExternalId: jsii.String("externalId"),
//   									ServiceArn: jsii.String("serviceArn"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   							TargetPercent: jsii.Number(123),
//   							TimeoutMinutes: jsii.Number(123),
//   							Ungraceful: &EcsUngracefulProperty{
//   								MinimumSuccessPercentage: jsii.Number(123),
//   							},
//   						},
//   						EksResourceScalingConfig: &EksResourceScalingConfigurationProperty{
//   							KubernetesResourceType: &KubernetesResourceTypeProperty{
//   								ApiVersion: jsii.String("apiVersion"),
//   								Kind: jsii.String("kind"),
//   							},
//
//   							// the properties below are optional
//   							CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   							EksClusters: []interface{}{
//   								&EksClusterProperty{
//   									ClusterArn: jsii.String("clusterArn"),
//
//   									// the properties below are optional
//   									CrossAccountRole: jsii.String("crossAccountRole"),
//   									ExternalId: jsii.String("externalId"),
//   								},
//   							},
//   							ScalingResources: []interface{}{
//   								map[string]interface{}{
//   									"scalingResourcesKey": map[string]interface{}{
//   										"scalingResourcesKey": &KubernetesScalingResourceProperty{
//   											"name": jsii.String("name"),
//   											"namespace": jsii.String("namespace"),
//
//   											// the properties below are optional
//   											"hpaName": jsii.String("hpaName"),
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
//
//   							// the properties below are optional
//   							TimeoutMinutes: jsii.Number(123),
//   						},
//   						GlobalAuroraConfig: &GlobalAuroraConfigurationProperty{
//   							Behavior: jsii.String("behavior"),
//   							DatabaseClusterArns: []*string{
//   								jsii.String("databaseClusterArns"),
//   							},
//   							GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//
//   							// the properties below are optional
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							ExternalId: jsii.String("externalId"),
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
//
//   							// the properties below are optional
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							ExternalId: jsii.String("externalId"),
//   						},
//   						Route53HealthCheckConfig: &Route53HealthCheckConfigurationProperty{
//   							HostedZoneId: jsii.String("hostedZoneId"),
//   							RecordName: jsii.String("recordName"),
//
//   							// the properties below are optional
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							ExternalId: jsii.String("externalId"),
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
//
//   					// the properties below are optional
//   					Description: jsii.String("description"),
//   				},
//   			},
//   			WorkflowDescription: jsii.String("workflowDescription"),
//   			WorkflowTargetRegion: jsii.String("workflowTargetRegion"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AssociatedAlarms: map[string]interface{}{
//   		"associatedAlarmsKey": &AssociatedAlarmProperty{
//   			"alarmType": jsii.String("alarmType"),
//   			"resourceIdentifier": jsii.String("resourceIdentifier"),
//
//   			// the properties below are optional
//   			"crossAccountRole": jsii.String("crossAccountRole"),
//   			"externalId": jsii.String("externalId"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	PrimaryRegion: jsii.String("primaryRegion"),
//   	RecoveryTimeObjectiveMinutes: jsii.Number(123),
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
//   			MinDelayMinutesBetweenExecutions: jsii.Number(123),
//   			TargetRegion: jsii.String("targetRegion"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html
//
type CfnPlanProps struct {
	// The execution role for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-executionrole
	//
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// The name for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The recovery approach for a Region switch plan, which can be active/active (activeActive) or active/passive (activePassive).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-recoveryapproach
	//
	RecoveryApproach *string `field:"required" json:"recoveryApproach" yaml:"recoveryApproach"`
	// The AWS Regions for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-regions
	//
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// The workflows for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-workflows
	//
	Workflows interface{} `field:"required" json:"workflows" yaml:"workflows"`
	// The associated application health alarms for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-associatedalarms
	//
	AssociatedAlarms interface{} `field:"optional" json:"associatedAlarms" yaml:"associatedAlarms"`
	// The description for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The primary Region for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-primaryregion
	//
	PrimaryRegion *string `field:"optional" json:"primaryRegion" yaml:"primaryRegion"`
	// The recovery time objective for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-recoverytimeobjectiveminutes
	//
	RecoveryTimeObjectiveMinutes *float64 `field:"optional" json:"recoveryTimeObjectiveMinutes" yaml:"recoveryTimeObjectiveMinutes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The triggers for a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html#cfn-arcregionswitch-plan-triggers
	//
	Triggers interface{} `field:"optional" json:"triggers" yaml:"triggers"`
}

