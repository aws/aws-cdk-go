package awsarcregionswitch


// Execution block configurations for a workflow in a Region switch plan.
//
// An execution block represents a specific type of action to perform during a Region switch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var executionBlockConfigurationProperty_ ExecutionBlockConfigurationProperty
//
//   executionBlockConfigurationProperty := &ExecutionBlockConfigurationProperty{
//   	ArcRoutingControlConfig: &ArcRoutingControlConfigurationProperty{
//   		RegionAndRoutingControls: map[string]interface{}{
//   			"regionAndRoutingControlsKey": []interface{}{
//   				&ArcRoutingControlStateProperty{
//   					"routingControlArn": jsii.String("routingControlArn"),
//   					"state": jsii.String("state"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		CrossAccountRole: jsii.String("crossAccountRole"),
//   		ExternalId: jsii.String("externalId"),
//   		TimeoutMinutes: jsii.Number(123),
//   	},
//   	AuroraProvisionedScalingConfig: &AuroraProvisionedScalingConfigurationProperty{
//   		GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   		InstanceArns: map[string]*string{
//   			"instanceArnsKey": jsii.String("instanceArns"),
//   		},
//   		RegionDatabaseClusterArns: map[string]*string{
//   			"regionDatabaseClusterArnsKey": jsii.String("regionDatabaseClusterArns"),
//   		},
//
//   		// the properties below are optional
//   		CrossAccountRole: jsii.String("crossAccountRole"),
//   		ExternalId: jsii.String("externalId"),
//   		TimeoutMinutes: jsii.Number(123),
//   	},
//   	AuroraServerlessScalingConfig: &AuroraServerlessScalingConfigurationProperty{
//   		GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   		RegionDatabaseClusterArns: map[string]*string{
//   			"regionDatabaseClusterArnsKey": jsii.String("regionDatabaseClusterArns"),
//   		},
//
//   		// the properties below are optional
//   		CrossAccountRole: jsii.String("crossAccountRole"),
//   		ExternalId: jsii.String("externalId"),
//   		TargetPercent: jsii.Number(123),
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
//
//   		// the properties below are optional
//   		TimeoutMinutes: jsii.Number(123),
//   		Ungraceful: &LambdaUngracefulProperty{
//   			Behavior: jsii.String("behavior"),
//   		},
//   	},
//   	DocumentDbConfig: &DocumentDbConfigurationProperty{
//   		Behavior: jsii.String("behavior"),
//   		DatabaseClusterArns: []*string{
//   			jsii.String("databaseClusterArns"),
//   		},
//   		GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//
//   		// the properties below are optional
//   		CrossAccountRole: jsii.String("crossAccountRole"),
//   		ExternalId: jsii.String("externalId"),
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
//
//   		// the properties below are optional
//   		CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   		TargetPercent: jsii.Number(123),
//   		TimeoutMinutes: jsii.Number(123),
//   		Ungraceful: &Ec2UngracefulProperty{
//   			MinimumSuccessPercentage: jsii.Number(123),
//   		},
//   	},
//   	EcsCapacityIncreaseConfig: &EcsCapacityIncreaseConfigurationProperty{
//   		Services: []interface{}{
//   			&ServiceProperty{
//   				ClusterArn: jsii.String("clusterArn"),
//   				CrossAccountRole: jsii.String("crossAccountRole"),
//   				ExternalId: jsii.String("externalId"),
//   				ServiceArn: jsii.String("serviceArn"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   		TargetPercent: jsii.Number(123),
//   		TimeoutMinutes: jsii.Number(123),
//   		Ungraceful: &EcsUngracefulProperty{
//   			MinimumSuccessPercentage: jsii.Number(123),
//   		},
//   	},
//   	EksResourceScalingConfig: &EksResourceScalingConfigurationProperty{
//   		KubernetesResourceType: &KubernetesResourceTypeProperty{
//   			ApiVersion: jsii.String("apiVersion"),
//   			Kind: jsii.String("kind"),
//   		},
//
//   		// the properties below are optional
//   		CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   		EksClusters: []interface{}{
//   			&EksClusterProperty{
//   				ClusterArn: jsii.String("clusterArn"),
//
//   				// the properties below are optional
//   				CrossAccountRole: jsii.String("crossAccountRole"),
//   				ExternalId: jsii.String("externalId"),
//   			},
//   		},
//   		ScalingResources: []interface{}{
//   			map[string]interface{}{
//   				"scalingResourcesKey": map[string]interface{}{
//   					"scalingResourcesKey": &KubernetesScalingResourceProperty{
//   						"name": jsii.String("name"),
//   						"namespace": jsii.String("namespace"),
//
//   						// the properties below are optional
//   						"hpaName": jsii.String("hpaName"),
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
//
//   		// the properties below are optional
//   		TimeoutMinutes: jsii.Number(123),
//   	},
//   	GlobalAuroraConfig: &GlobalAuroraConfigurationProperty{
//   		Behavior: jsii.String("behavior"),
//   		DatabaseClusterArns: []*string{
//   			jsii.String("databaseClusterArns"),
//   		},
//   		GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//
//   		// the properties below are optional
//   		CrossAccountRole: jsii.String("crossAccountRole"),
//   		ExternalId: jsii.String("externalId"),
//   		TimeoutMinutes: jsii.Number(123),
//   		Ungraceful: &GlobalAuroraUngracefulProperty{
//   			Ungraceful: jsii.String("ungraceful"),
//   		},
//   	},
//   	LambdaEventSourceMappingConfig: &LambdaEventSourceMappingConfigurationProperty{
//   		Action: jsii.String("action"),
//   		RegionEventSourceMappings: map[string]interface{}{
//   			"regionEventSourceMappingsKey": &EventSourceMappingProperty{
//   				"arn": jsii.String("arn"),
//
//   				// the properties below are optional
//   				"crossAccountRole": jsii.String("crossAccountRole"),
//   				"externalId": jsii.String("externalId"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		TimeoutMinutes: jsii.Number(123),
//   		Ungraceful: &LambdaEventSourceMappingUngracefulProperty{
//   			Behavior: jsii.String("behavior"),
//   		},
//   	},
//   	NeptuneGlobalDatabaseConfig: &NeptuneGlobalDatabaseConfigurationProperty{
//   		Behavior: jsii.String("behavior"),
//   		GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   		RegionDatabaseClusterArns: map[string]*string{
//   			"regionDatabaseClusterArnsKey": jsii.String("regionDatabaseClusterArns"),
//   		},
//
//   		// the properties below are optional
//   		CrossAccountRole: jsii.String("crossAccountRole"),
//   		ExternalId: jsii.String("externalId"),
//   		TimeoutMinutes: jsii.Number(123),
//   		Ungraceful: &NeptuneUngracefulProperty{
//   			Ungraceful: jsii.String("ungraceful"),
//   		},
//   	},
//   	ParallelConfig: &ParallelExecutionBlockConfigurationProperty{
//   		Steps: []interface{}{
//   			&StepProperty{
//   				ExecutionBlockConfiguration: executionBlockConfigurationProperty_,
//   				ExecutionBlockType: jsii.String("executionBlockType"),
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Description: jsii.String("description"),
//   			},
//   		},
//   	},
//   	RdsCreateCrossRegionReadReplicaConfig: &RdsCreateCrossRegionReplicaConfigurationProperty{
//   		DbInstanceArnMap: map[string]*string{
//   			"dbInstanceArnMapKey": jsii.String("dbInstanceArnMap"),
//   		},
//
//   		// the properties below are optional
//   		CrossAccountRole: jsii.String("crossAccountRole"),
//   		ExternalId: jsii.String("externalId"),
//   		TimeoutMinutes: jsii.Number(123),
//   	},
//   	RdsPromoteReadReplicaConfig: &RdsPromoteReadReplicaConfigurationProperty{
//   		DbInstanceArnMap: map[string]*string{
//   			"dbInstanceArnMapKey": jsii.String("dbInstanceArnMap"),
//   		},
//
//   		// the properties below are optional
//   		CrossAccountRole: jsii.String("crossAccountRole"),
//   		ExternalId: jsii.String("externalId"),
//   		TimeoutMinutes: jsii.Number(123),
//   	},
//   	RegionSwitchPlanConfig: &RegionSwitchPlanConfigurationProperty{
//   		Arn: jsii.String("arn"),
//
//   		// the properties below are optional
//   		CrossAccountRole: jsii.String("crossAccountRole"),
//   		ExternalId: jsii.String("externalId"),
//   	},
//   	Route53HealthCheckConfig: &Route53HealthCheckConfigurationProperty{
//   		HostedZoneId: jsii.String("hostedZoneId"),
//   		RecordName: jsii.String("recordName"),
//
//   		// the properties below are optional
//   		CrossAccountRole: jsii.String("crossAccountRole"),
//   		ExternalId: jsii.String("externalId"),
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
type CfnPlan_ExecutionBlockConfigurationProperty struct {
	// An ARC routing control execution block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-arcroutingcontrolconfig
	//
	ArcRoutingControlConfig interface{} `field:"optional" json:"arcRoutingControlConfig" yaml:"arcRoutingControlConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-auroraprovisionedscalingconfig
	//
	AuroraProvisionedScalingConfig interface{} `field:"optional" json:"auroraProvisionedScalingConfig" yaml:"auroraProvisionedScalingConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-auroraserverlessscalingconfig
	//
	AuroraServerlessScalingConfig interface{} `field:"optional" json:"auroraServerlessScalingConfig" yaml:"auroraServerlessScalingConfig"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-lambdaeventsourcemappingconfig
	//
	LambdaEventSourceMappingConfig interface{} `field:"optional" json:"lambdaEventSourceMappingConfig" yaml:"lambdaEventSourceMappingConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-neptuneglobaldatabaseconfig
	//
	NeptuneGlobalDatabaseConfig interface{} `field:"optional" json:"neptuneGlobalDatabaseConfig" yaml:"neptuneGlobalDatabaseConfig"`
	// A parallel configuration execution block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-parallelconfig
	//
	ParallelConfig interface{} `field:"optional" json:"parallelConfig" yaml:"parallelConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-rdscreatecrossregionreadreplicaconfig
	//
	RdsCreateCrossRegionReadReplicaConfig interface{} `field:"optional" json:"rdsCreateCrossRegionReadReplicaConfig" yaml:"rdsCreateCrossRegionReadReplicaConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-rdspromotereadreplicaconfig
	//
	RdsPromoteReadReplicaConfig interface{} `field:"optional" json:"rdsPromoteReadReplicaConfig" yaml:"rdsPromoteReadReplicaConfig"`
	// A Region switch plan execution block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-regionswitchplanconfig
	//
	RegionSwitchPlanConfig interface{} `field:"optional" json:"regionSwitchPlanConfig" yaml:"regionSwitchPlanConfig"`
	// The Amazon Route 53 health check configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionblockconfiguration.html#cfn-arcregionswitch-plan-executionblockconfiguration-route53healthcheckconfig
	//
	Route53HealthCheckConfig interface{} `field:"optional" json:"route53HealthCheckConfig" yaml:"route53HealthCheckConfig"`
}

