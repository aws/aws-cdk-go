package awsarcregionswitch


// Represents a workflow in a Region switch plan.
//
// A workflow defines a sequence of steps to execute during a Region switch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stepProperty_ stepProperty
//
//   workflowProperty := &WorkflowProperty{
//   	WorkflowTargetAction: jsii.String("workflowTargetAction"),
//
//   	// the properties below are optional
//   	Steps: []interface{}{
//   		&stepProperty{
//   			ExecutionBlockConfiguration: &ExecutionBlockConfigurationProperty{
//   				ArcRoutingControlConfig: &ArcRoutingControlConfigurationProperty{
//   					RegionAndRoutingControls: map[string]interface{}{
//   						"regionAndRoutingControlsKey": []interface{}{
//   							&ArcRoutingControlStateProperty{
//   								"routingControlArn": jsii.String("routingControlArn"),
//   								"state": jsii.String("state"),
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					CrossAccountRole: jsii.String("crossAccountRole"),
//   					ExternalId: jsii.String("externalId"),
//   					TimeoutMinutes: jsii.Number(123),
//   				},
//   				CustomActionLambdaConfig: &CustomActionLambdaConfigurationProperty{
//   					Lambdas: []interface{}{
//   						&LambdasProperty{
//   							Arn: jsii.String("arn"),
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							ExternalId: jsii.String("externalId"),
//   						},
//   					},
//   					RegionToRun: jsii.String("regionToRun"),
//   					RetryIntervalMinutes: jsii.Number(123),
//
//   					// the properties below are optional
//   					TimeoutMinutes: jsii.Number(123),
//   					Ungraceful: &LambdaUngracefulProperty{
//   						Behavior: jsii.String("behavior"),
//   					},
//   				},
//   				Ec2AsgCapacityIncreaseConfig: &Ec2AsgCapacityIncreaseConfigurationProperty{
//   					Asgs: []interface{}{
//   						&AsgProperty{
//   							Arn: jsii.String("arn"),
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							ExternalId: jsii.String("externalId"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   					TargetPercent: jsii.Number(123),
//   					TimeoutMinutes: jsii.Number(123),
//   					Ungraceful: &Ec2UngracefulProperty{
//   						MinimumSuccessPercentage: jsii.Number(123),
//   					},
//   				},
//   				EcsCapacityIncreaseConfig: &EcsCapacityIncreaseConfigurationProperty{
//   					Services: []interface{}{
//   						&ServiceProperty{
//   							ClusterArn: jsii.String("clusterArn"),
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							ExternalId: jsii.String("externalId"),
//   							ServiceArn: jsii.String("serviceArn"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   					TargetPercent: jsii.Number(123),
//   					TimeoutMinutes: jsii.Number(123),
//   					Ungraceful: &EcsUngracefulProperty{
//   						MinimumSuccessPercentage: jsii.Number(123),
//   					},
//   				},
//   				EksResourceScalingConfig: &EksResourceScalingConfigurationProperty{
//   					KubernetesResourceType: &KubernetesResourceTypeProperty{
//   						ApiVersion: jsii.String("apiVersion"),
//   						Kind: jsii.String("kind"),
//   					},
//
//   					// the properties below are optional
//   					CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   					EksClusters: []interface{}{
//   						&EksClusterProperty{
//   							ClusterArn: jsii.String("clusterArn"),
//
//   							// the properties below are optional
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							ExternalId: jsii.String("externalId"),
//   						},
//   					},
//   					ScalingResources: []interface{}{
//   						map[string]interface{}{
//   							"scalingResourcesKey": map[string]interface{}{
//   								"scalingResourcesKey": &KubernetesScalingResourceProperty{
//   									"name": jsii.String("name"),
//   									"namespace": jsii.String("namespace"),
//
//   									// the properties below are optional
//   									"hpaName": jsii.String("hpaName"),
//   								},
//   							},
//   						},
//   					},
//   					TargetPercent: jsii.Number(123),
//   					TimeoutMinutes: jsii.Number(123),
//   					Ungraceful: &EksResourceScalingUngracefulProperty{
//   						MinimumSuccessPercentage: jsii.Number(123),
//   					},
//   				},
//   				ExecutionApprovalConfig: &ExecutionApprovalConfigurationProperty{
//   					ApprovalRole: jsii.String("approvalRole"),
//
//   					// the properties below are optional
//   					TimeoutMinutes: jsii.Number(123),
//   				},
//   				GlobalAuroraConfig: &GlobalAuroraConfigurationProperty{
//   					Behavior: jsii.String("behavior"),
//   					DatabaseClusterArns: []*string{
//   						jsii.String("databaseClusterArns"),
//   					},
//   					GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//
//   					// the properties below are optional
//   					CrossAccountRole: jsii.String("crossAccountRole"),
//   					ExternalId: jsii.String("externalId"),
//   					TimeoutMinutes: jsii.Number(123),
//   					Ungraceful: &GlobalAuroraUngracefulProperty{
//   						Ungraceful: jsii.String("ungraceful"),
//   					},
//   				},
//   				ParallelConfig: &ParallelExecutionBlockConfigurationProperty{
//   					Steps: []interface{}{
//   						stepProperty_,
//   					},
//   				},
//   				RegionSwitchPlanConfig: &RegionSwitchPlanConfigurationProperty{
//   					Arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					CrossAccountRole: jsii.String("crossAccountRole"),
//   					ExternalId: jsii.String("externalId"),
//   				},
//   				Route53HealthCheckConfig: &Route53HealthCheckConfigurationProperty{
//   					HostedZoneId: jsii.String("hostedZoneId"),
//   					RecordName: jsii.String("recordName"),
//
//   					// the properties below are optional
//   					CrossAccountRole: jsii.String("crossAccountRole"),
//   					ExternalId: jsii.String("externalId"),
//   					RecordSets: []interface{}{
//   						&Route53ResourceRecordSetProperty{
//   							RecordSetIdentifier: jsii.String("recordSetIdentifier"),
//   							Region: jsii.String("region"),
//   						},
//   					},
//   					TimeoutMinutes: jsii.Number(123),
//   				},
//   			},
//   			ExecutionBlockType: jsii.String("executionBlockType"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   	},
//   	WorkflowDescription: jsii.String("workflowDescription"),
//   	WorkflowTargetRegion: jsii.String("workflowTargetRegion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-workflow.html
//
type CfnPlan_WorkflowProperty struct {
	// The action that the workflow performs.
	//
	// Valid values include ACTIVATE and DEACTIVATE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-workflow.html#cfn-arcregionswitch-plan-workflow-workflowtargetaction
	//
	WorkflowTargetAction *string `field:"required" json:"workflowTargetAction" yaml:"workflowTargetAction"`
	// The steps that make up the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-workflow.html#cfn-arcregionswitch-plan-workflow-steps
	//
	Steps interface{} `field:"optional" json:"steps" yaml:"steps"`
	// The description of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-workflow.html#cfn-arcregionswitch-plan-workflow-workflowdescription
	//
	WorkflowDescription *string `field:"optional" json:"workflowDescription" yaml:"workflowDescription"`
	// The AWS Region that the workflow targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-workflow.html#cfn-arcregionswitch-plan-workflow-workflowtargetregion
	//
	WorkflowTargetRegion *string `field:"optional" json:"workflowTargetRegion" yaml:"workflowTargetRegion"`
}

