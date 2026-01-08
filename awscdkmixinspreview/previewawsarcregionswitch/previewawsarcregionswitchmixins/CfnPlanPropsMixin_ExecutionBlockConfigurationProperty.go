package previewawsarcregionswitchmixins


// Execution block configurations for a workflow in a Region switch plan.
//
// An execution block represents a specific type of action to perform during a Region switch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var executionBlockConfigurationProperty_ ExecutionBlockConfigurationProperty
//
//   executionBlockConfigurationProperty := &ExecutionBlockConfigurationProperty{
//   	ArcRoutingControlConfig: &ArcRoutingControlConfigurationProperty{
//   		CrossAccountRole: jsii.String("crossAccountRole"),
//   		ExternalId: jsii.String("externalId"),
//   		RegionAndRoutingControls: map[string]interface{}{
//   			"regionAndRoutingControlsKey": []interface{}{
//   				&ArcRoutingControlStateProperty{
//   					"routingControlArn": jsii.String("routingControlArn"),
//   					"state": jsii.String("state"),
//   				},
//   			},
//   		},
//   		TimeoutMinutes: jsii.Number(123),
//   	},
//   	CustomActionLambdaConfig: &CustomActionLambdaConfigurationProperty{
//   		Lambdas: []interface{}{
//   			&LambdasProperty{
//   				Arn: jsii.String("arn"),
//   				CrossAccountRole: jsii.String("crossAccountRole"),
//   				ExternalId: jsii.String("externalId"),
//   			},
//   		},
//   		RegionToRun: jsii.String("regionToRun"),
//   		RetryIntervalMinutes: jsii.Number(123),
//   		TimeoutMinutes: jsii.Number(123),
//   		Ungraceful: &LambdaUngracefulProperty{
//   			Behavior: jsii.String("behavior"),
//   		},
//   	},
//   	DocumentDbConfig: &DocumentDbConfigurationProperty{
//   		Behavior: jsii.String("behavior"),
//   		CrossAccountRole: jsii.String("crossAccountRole"),
//   		DatabaseClusterArns: []*string{
//   			jsii.String("databaseClusterArns"),
//   		},
//   		ExternalId: jsii.String("externalId"),
//   		GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   		TimeoutMinutes: jsii.Number(123),
//   		Ungraceful: &DocumentDbUngracefulProperty{
//   			Ungraceful: jsii.String("ungraceful"),
//   		},
//   	},
//   	Ec2AsgCapacityIncreaseConfig: &Ec2AsgCapacityIncreaseConfigurationProperty{
//   		Asgs: []interface{}{
//   			&AsgProperty{
//   				Arn: jsii.String("arn"),
//   				CrossAccountRole: jsii.String("crossAccountRole"),
//   				ExternalId: jsii.String("externalId"),
//   			},
//   		},
//   		CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   		TargetPercent: jsii.Number(123),
//   		TimeoutMinutes: jsii.Number(123),
//   		Ungraceful: &Ec2UngracefulProperty{
//   			MinimumSuccessPercentage: jsii.Number(123),
//   		},
//   	},
//   	EcsCapacityIncreaseConfig: &EcsCapacityIncreaseConfigurationProperty{
//   		CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   		Services: []interface{}{
//   			&ServiceProperty{
//   				ClusterArn: jsii.String("clusterArn"),
//   				CrossAccountRole: jsii.String("crossAccountRole"),
//   				ExternalId: jsii.String("externalId"),
//   				ServiceArn: jsii.String("serviceArn"),
//   			},
//   		},
//   		TargetPercent: jsii.Number(123),
//   		TimeoutMinutes: jsii.Number(123),
//   		Ungraceful: &EcsUngracefulProperty{
//   			MinimumSuccessPercentage: jsii.Number(123),
//   		},
//   	},
//   	EksResourceScalingConfig: &EksResourceScalingConfigurationProperty{
//   		CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   		EksClusters: []interface{}{
//   			&EksClusterProperty{
//   				ClusterArn: jsii.String("clusterArn"),
//   				CrossAccountRole: jsii.String("crossAccountRole"),
//   				ExternalId: jsii.String("externalId"),
//   			},
//   		},
//   		KubernetesResourceType: &KubernetesResourceTypeProperty{
//   			ApiVersion: jsii.String("apiVersion"),
//   			Kind: jsii.String("kind"),
//   		},
//   		ScalingResources: []interface{}{
//   			map[string]interface{}{
//   				"scalingResourcesKey": map[string]interface{}{
//   					"scalingResourcesKey": &KubernetesScalingResourceProperty{
//   						"hpaName": jsii.String("hpaName"),
//   						"name": jsii.String("name"),
//   						"namespace": jsii.String("namespace"),
//   					},
//   				},
//   			},
//   		},
//   		TargetPercent: jsii.Number(123),
//   		TimeoutMinutes: jsii.Number(123),
//   		Ungraceful: &EksResourceScalingUngracefulProperty{
//   			MinimumSuccessPercentage: jsii.Number(123),
//   		},
//   	},
//   	ExecutionApprovalConfig: &ExecutionApprovalConfigurationProperty{
//   		ApprovalRole: jsii.String("approvalRole"),
//   		TimeoutMinutes: jsii.Number(123),
//   	},
//   	GlobalAuroraConfig: &GlobalAuroraConfigurationProperty{
//   		Behavior: jsii.String("behavior"),
//   		CrossAccountRole: jsii.String("crossAccountRole"),
//   		DatabaseClusterArns: []*string{
//   			jsii.String("databaseClusterArns"),
//   		},
//   		ExternalId: jsii.String("externalId"),
//   		GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   		TimeoutMinutes: jsii.Number(123),
//   		Ungraceful: &GlobalAuroraUngracefulProperty{
//   			Ungraceful: jsii.String("ungraceful"),
//   		},
//   	},
//   	ParallelConfig: &ParallelExecutionBlockConfigurationProperty{
//   		Steps: []interface{}{
//   			&StepProperty{
//   				Description: jsii.String("description"),
//   				ExecutionBlockConfiguration: executionBlockConfigurationProperty_,
//   				ExecutionBlockType: jsii.String("executionBlockType"),
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	RegionSwitchPlanConfig: &RegionSwitchPlanConfigurationProperty{
//   		Arn: jsii.String("arn"),
//   		CrossAccountRole: jsii.String("crossAccountRole"),
//   		ExternalId: jsii.String("externalId"),
//   	},
//   	Route53HealthCheckConfig: &Route53HealthCheckConfigurationProperty{
//   		CrossAccountRole: jsii.String("crossAccountRole"),
//   		ExternalId: jsii.String("externalId"),
//   		HostedZoneId: jsii.String("hostedZoneId"),
//   		RecordName: jsii.String("recordName"),
//   		RecordSets: []interface{}{
//   			&Route53ResourceRecordSetProperty{
//   				RecordSetIdentifier: jsii.String("recordSetIdentifier"),
//   				Region: jsii.String("region"),
//   			},
//   		},
//   		TimeoutMinutes: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html
//
type CfnPlanPropsMixin_ExecutionBlockConfigurationProperty struct {
	// An ARC routing control execution block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-arcroutingcontrolconfig
	//
	ArcRoutingControlConfig interface{} `field:"optional" json:"arcRoutingControlConfig" yaml:"arcRoutingControlConfig"`
	// An AWS Lambda execution block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-customactionlambdaconfig
	//
	CustomActionLambdaConfig interface{} `field:"optional" json:"customActionLambdaConfig" yaml:"customActionLambdaConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-documentdbconfig
	//
	DocumentDbConfig interface{} `field:"optional" json:"documentDbConfig" yaml:"documentDbConfig"`
	// An EC2 Auto Scaling group execution block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-ec2asgcapacityincreaseconfig
	//
	Ec2AsgCapacityIncreaseConfig interface{} `field:"optional" json:"ec2AsgCapacityIncreaseConfig" yaml:"ec2AsgCapacityIncreaseConfig"`
	// The capacity increase specified for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-ecscapacityincreaseconfig
	//
	EcsCapacityIncreaseConfig interface{} `field:"optional" json:"ecsCapacityIncreaseConfig" yaml:"ecsCapacityIncreaseConfig"`
	// An AWS EKS resource scaling execution block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-eksresourcescalingconfig
	//
	EksResourceScalingConfig interface{} `field:"optional" json:"eksResourceScalingConfig" yaml:"eksResourceScalingConfig"`
	// A manual approval execution block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-executionapprovalconfig
	//
	ExecutionApprovalConfig interface{} `field:"optional" json:"executionApprovalConfig" yaml:"executionApprovalConfig"`
	// An Aurora Global Database execution block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-globalauroraconfig
	//
	GlobalAuroraConfig interface{} `field:"optional" json:"globalAuroraConfig" yaml:"globalAuroraConfig"`
	// A parallel configuration execution block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-parallelconfig
	//
	ParallelConfig interface{} `field:"optional" json:"parallelConfig" yaml:"parallelConfig"`
	// A Region switch plan execution block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-regionswitchplanconfig
	//
	RegionSwitchPlanConfig interface{} `field:"optional" json:"regionSwitchPlanConfig" yaml:"regionSwitchPlanConfig"`
	// The Amazon Route 53 health check configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-route53healthcheckconfig
	//
	Route53HealthCheckConfig interface{} `field:"optional" json:"route53HealthCheckConfig" yaml:"route53HealthCheckConfig"`
}

