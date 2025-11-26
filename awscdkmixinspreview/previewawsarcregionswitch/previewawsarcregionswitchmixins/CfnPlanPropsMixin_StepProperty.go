package previewawsarcregionswitchmixins


// Represents a step in a Region switch plan workflow.
//
// Each step performs a specific action during the Region switch process.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var stepProperty_ StepProperty
//
//   stepProperty := &StepProperty{
//   	Description: jsii.String("description"),
//   	ExecutionBlockConfiguration: &ExecutionBlockConfigurationProperty{
//   		ArcRoutingControlConfig: &ArcRoutingControlConfigurationProperty{
//   			CrossAccountRole: jsii.String("crossAccountRole"),
//   			ExternalId: jsii.String("externalId"),
//   			RegionAndRoutingControls: map[string]interface{}{
//   				"regionAndRoutingControlsKey": []interface{}{
//   					&ArcRoutingControlStateProperty{
//   						"routingControlArn": jsii.String("routingControlArn"),
//   						"state": jsii.String("state"),
//   					},
//   				},
//   			},
//   			TimeoutMinutes: jsii.Number(123),
//   		},
//   		CustomActionLambdaConfig: &CustomActionLambdaConfigurationProperty{
//   			Lambdas: []interface{}{
//   				&LambdasProperty{
//   					Arn: jsii.String("arn"),
//   					CrossAccountRole: jsii.String("crossAccountRole"),
//   					ExternalId: jsii.String("externalId"),
//   				},
//   			},
//   			RegionToRun: jsii.String("regionToRun"),
//   			RetryIntervalMinutes: jsii.Number(123),
//   			TimeoutMinutes: jsii.Number(123),
//   			Ungraceful: &LambdaUngracefulProperty{
//   				Behavior: jsii.String("behavior"),
//   			},
//   		},
//   		Ec2AsgCapacityIncreaseConfig: &Ec2AsgCapacityIncreaseConfigurationProperty{
//   			Asgs: []interface{}{
//   				&AsgProperty{
//   					Arn: jsii.String("arn"),
//   					CrossAccountRole: jsii.String("crossAccountRole"),
//   					ExternalId: jsii.String("externalId"),
//   				},
//   			},
//   			CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   			TargetPercent: jsii.Number(123),
//   			TimeoutMinutes: jsii.Number(123),
//   			Ungraceful: &Ec2UngracefulProperty{
//   				MinimumSuccessPercentage: jsii.Number(123),
//   			},
//   		},
//   		EcsCapacityIncreaseConfig: &EcsCapacityIncreaseConfigurationProperty{
//   			CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   			Services: []interface{}{
//   				&ServiceProperty{
//   					ClusterArn: jsii.String("clusterArn"),
//   					CrossAccountRole: jsii.String("crossAccountRole"),
//   					ExternalId: jsii.String("externalId"),
//   					ServiceArn: jsii.String("serviceArn"),
//   				},
//   			},
//   			TargetPercent: jsii.Number(123),
//   			TimeoutMinutes: jsii.Number(123),
//   			Ungraceful: &EcsUngracefulProperty{
//   				MinimumSuccessPercentage: jsii.Number(123),
//   			},
//   		},
//   		EksResourceScalingConfig: &EksResourceScalingConfigurationProperty{
//   			CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   			EksClusters: []interface{}{
//   				&EksClusterProperty{
//   					ClusterArn: jsii.String("clusterArn"),
//   					CrossAccountRole: jsii.String("crossAccountRole"),
//   					ExternalId: jsii.String("externalId"),
//   				},
//   			},
//   			KubernetesResourceType: &KubernetesResourceTypeProperty{
//   				ApiVersion: jsii.String("apiVersion"),
//   				Kind: jsii.String("kind"),
//   			},
//   			ScalingResources: []interface{}{
//   				map[string]interface{}{
//   					"scalingResourcesKey": map[string]interface{}{
//   						"scalingResourcesKey": &KubernetesScalingResourceProperty{
//   							"hpaName": jsii.String("hpaName"),
//   							"name": jsii.String("name"),
//   							"namespace": jsii.String("namespace"),
//   						},
//   					},
//   				},
//   			},
//   			TargetPercent: jsii.Number(123),
//   			TimeoutMinutes: jsii.Number(123),
//   			Ungraceful: &EksResourceScalingUngracefulProperty{
//   				MinimumSuccessPercentage: jsii.Number(123),
//   			},
//   		},
//   		ExecutionApprovalConfig: &ExecutionApprovalConfigurationProperty{
//   			ApprovalRole: jsii.String("approvalRole"),
//   			TimeoutMinutes: jsii.Number(123),
//   		},
//   		GlobalAuroraConfig: &GlobalAuroraConfigurationProperty{
//   			Behavior: jsii.String("behavior"),
//   			CrossAccountRole: jsii.String("crossAccountRole"),
//   			DatabaseClusterArns: []*string{
//   				jsii.String("databaseClusterArns"),
//   			},
//   			ExternalId: jsii.String("externalId"),
//   			GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   			TimeoutMinutes: jsii.Number(123),
//   			Ungraceful: &GlobalAuroraUngracefulProperty{
//   				Ungraceful: jsii.String("ungraceful"),
//   			},
//   		},
//   		ParallelConfig: &ParallelExecutionBlockConfigurationProperty{
//   			Steps: []interface{}{
//   				stepProperty_,
//   			},
//   		},
//   		RegionSwitchPlanConfig: &RegionSwitchPlanConfigurationProperty{
//   			Arn: jsii.String("arn"),
//   			CrossAccountRole: jsii.String("crossAccountRole"),
//   			ExternalId: jsii.String("externalId"),
//   		},
//   		Route53HealthCheckConfig: &Route53HealthCheckConfigurationProperty{
//   			CrossAccountRole: jsii.String("crossAccountRole"),
//   			ExternalId: jsii.String("externalId"),
//   			HostedZoneId: jsii.String("hostedZoneId"),
//   			RecordName: jsii.String("recordName"),
//   			RecordSets: []interface{}{
//   				&Route53ResourceRecordSetProperty{
//   					RecordSetIdentifier: jsii.String("recordSetIdentifier"),
//   					Region: jsii.String("region"),
//   				},
//   			},
//   			TimeoutMinutes: jsii.Number(123),
//   		},
//   	},
//   	ExecutionBlockType: jsii.String("executionBlockType"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-step.html
//
type CfnPlanPropsMixin_StepProperty struct {
	// The description of a step in a workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-step.html#cfn-arcregionswitch-plan-step-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration for an execution block in a workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-step.html#cfn-arcregionswitch-plan-step-executionblockconfiguration
	//
	ExecutionBlockConfiguration interface{} `field:"optional" json:"executionBlockConfiguration" yaml:"executionBlockConfiguration"`
	// The type of an execution block in a workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-step.html#cfn-arcregionswitch-plan-step-executionblocktype
	//
	ExecutionBlockType *string `field:"optional" json:"executionBlockType" yaml:"executionBlockType"`
	// The name of a step in a workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-step.html#cfn-arcregionswitch-plan-step-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

