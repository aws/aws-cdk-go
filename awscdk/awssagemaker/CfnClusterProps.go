package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var onDemand interface{}
//   var spot interface{}
//
//   cfnClusterProps := &CfnClusterProps{
//   	AutoScaling: &ClusterAutoScalingConfigProperty{
//   		Mode: jsii.String("mode"),
//
//   		// the properties below are optional
//   		AutoScalerType: jsii.String("autoScalerType"),
//   	},
//   	ClusterName: jsii.String("clusterName"),
//   	ClusterRole: jsii.String("clusterRole"),
//   	InstanceGroups: []interface{}{
//   		&ClusterInstanceGroupProperty{
//   			ExecutionRole: jsii.String("executionRole"),
//   			InstanceCount: jsii.Number(123),
//   			InstanceGroupName: jsii.String("instanceGroupName"),
//   			InstanceType: jsii.String("instanceType"),
//   			LifeCycleConfig: &ClusterLifeCycleConfigProperty{
//   				OnCreate: jsii.String("onCreate"),
//   				SourceS3Uri: jsii.String("sourceS3Uri"),
//   			},
//
//   			// the properties below are optional
//   			CapacityRequirements: &ClusterCapacityRequirementsProperty{
//   				OnDemand: onDemand,
//   				Spot: spot,
//   			},
//   			CurrentCount: jsii.Number(123),
//   			ImageId: jsii.String("imageId"),
//   			InstanceStorageConfigs: []interface{}{
//   				&ClusterInstanceStorageConfigProperty{
//   					EbsVolumeConfig: &ClusterEbsVolumeConfigProperty{
//   						RootVolume: jsii.Boolean(false),
//   						VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   						VolumeSizeInGb: jsii.Number(123),
//   					},
//   					FsxLustreConfig: &ClusterFsxLustreConfigProperty{
//   						DnsName: jsii.String("dnsName"),
//   						MountName: jsii.String("mountName"),
//
//   						// the properties below are optional
//   						MountPath: jsii.String("mountPath"),
//   					},
//   					FsxOpenZfsConfig: &ClusterFsxOpenZfsConfigProperty{
//   						DnsName: jsii.String("dnsName"),
//
//   						// the properties below are optional
//   						MountPath: jsii.String("mountPath"),
//   					},
//   				},
//   			},
//   			KubernetesConfig: &ClusterKubernetesConfigProperty{
//   				Labels: map[string]*string{
//   					"labelsKey": jsii.String("labels"),
//   				},
//   				Taints: []interface{}{
//   					&ClusterKubernetesTaintProperty{
//   						Effect: jsii.String("effect"),
//   						Key: jsii.String("key"),
//
//   						// the properties below are optional
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			MinInstanceCount: jsii.Number(123),
//   			OnStartDeepHealthChecks: []*string{
//   				jsii.String("onStartDeepHealthChecks"),
//   			},
//   			OverrideVpcConfig: &VpcConfigProperty{
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   			ScheduledUpdateConfig: &ScheduledUpdateConfigProperty{
//   				ScheduleExpression: jsii.String("scheduleExpression"),
//
//   				// the properties below are optional
//   				DeploymentConfig: &DeploymentConfigProperty{
//   					AutoRollbackConfiguration: []interface{}{
//   						&AlarmDetailsProperty{
//   							AlarmName: jsii.String("alarmName"),
//   						},
//   					},
//   					RollingUpdatePolicy: &RollingUpdatePolicyProperty{
//   						MaximumBatchSize: &CapacitySizeConfigProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.Number(123),
//   						},
//
//   						// the properties below are optional
//   						RollbackMaximumBatchSize: &CapacitySizeConfigProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.Number(123),
//   						},
//   					},
//   					WaitIntervalInSeconds: jsii.Number(123),
//   				},
//   			},
//   			SlurmConfig: &ClusterSlurmConfigProperty{
//   				NodeType: jsii.String("nodeType"),
//
//   				// the properties below are optional
//   				PartitionNames: []*string{
//   					jsii.String("partitionNames"),
//   				},
//   			},
//   			ThreadsPerCore: jsii.Number(123),
//   			TrainingPlanArn: jsii.String("trainingPlanArn"),
//   		},
//   	},
//   	NodeProvisioningMode: jsii.String("nodeProvisioningMode"),
//   	NodeRecovery: jsii.String("nodeRecovery"),
//   	Orchestrator: &OrchestratorProperty{
//   		Eks: &ClusterOrchestratorEksConfigProperty{
//   			ClusterArn: jsii.String("clusterArn"),
//   		},
//   		Slurm: &ClusterOrchestratorSlurmConfigProperty{
//   			SlurmConfigStrategy: jsii.String("slurmConfigStrategy"),
//   		},
//   	},
//   	RestrictedInstanceGroups: []interface{}{
//   		&ClusterRestrictedInstanceGroupProperty{
//   			EnvironmentConfig: &EnvironmentConfigProperty{
//   				FSxLustreConfig: &FSxLustreConfigProperty{
//   					PerUnitStorageThroughput: jsii.Number(123),
//   					SizeInGiB: jsii.Number(123),
//   				},
//   			},
//   			ExecutionRole: jsii.String("executionRole"),
//   			InstanceCount: jsii.Number(123),
//   			InstanceGroupName: jsii.String("instanceGroupName"),
//   			InstanceType: jsii.String("instanceType"),
//
//   			// the properties below are optional
//   			CurrentCount: jsii.Number(123),
//   			InstanceStorageConfigs: []interface{}{
//   				&ClusterInstanceStorageConfigProperty{
//   					EbsVolumeConfig: &ClusterEbsVolumeConfigProperty{
//   						RootVolume: jsii.Boolean(false),
//   						VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   						VolumeSizeInGb: jsii.Number(123),
//   					},
//   					FsxLustreConfig: &ClusterFsxLustreConfigProperty{
//   						DnsName: jsii.String("dnsName"),
//   						MountName: jsii.String("mountName"),
//
//   						// the properties below are optional
//   						MountPath: jsii.String("mountPath"),
//   					},
//   					FsxOpenZfsConfig: &ClusterFsxOpenZfsConfigProperty{
//   						DnsName: jsii.String("dnsName"),
//
//   						// the properties below are optional
//   						MountPath: jsii.String("mountPath"),
//   					},
//   				},
//   			},
//   			OnStartDeepHealthChecks: []*string{
//   				jsii.String("onStartDeepHealthChecks"),
//   			},
//   			OverrideVpcConfig: &VpcConfigProperty{
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   			ThreadsPerCore: jsii.Number(123),
//   			TrainingPlanArn: jsii.String("trainingPlanArn"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TieredStorageConfig: &TieredStorageConfigProperty{
//   		Mode: jsii.String("mode"),
//
//   		// the properties below are optional
//   		InstanceMemoryAllocationPercentage: jsii.Number(123),
//   	},
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html
//
type CfnClusterProps struct {
	// Configuration for cluster auto-scaling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-autoscaling
	//
	AutoScaling interface{} `field:"optional" json:"autoScaling" yaml:"autoScaling"`
	// The name of the SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The cluster role for the autoscaler to assume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-clusterrole
	//
	ClusterRole *string `field:"optional" json:"clusterRole" yaml:"clusterRole"`
	// The instance groups of the SageMaker HyperPod cluster.
	//
	// To delete an instance group, remove it from the array.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-instancegroups
	//
	InstanceGroups interface{} `field:"optional" json:"instanceGroups" yaml:"instanceGroups"`
	// Determines the scaling strategy for the SageMaker HyperPod cluster.
	//
	// When set to 'Continuous', enables continuous scaling which dynamically manages node provisioning. If the parameter is omitted, uses the standard scaling approach in previous release.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-nodeprovisioningmode
	//
	NodeProvisioningMode *string `field:"optional" json:"nodeProvisioningMode" yaml:"nodeProvisioningMode"`
	// Specifies whether to enable or disable the automatic node recovery feature of SageMaker HyperPod.
	//
	// Available values are `Automatic` for enabling and `None` for disabling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-noderecovery
	//
	NodeRecovery *string `field:"optional" json:"nodeRecovery" yaml:"nodeRecovery"`
	// The orchestrator type for the SageMaker HyperPod cluster.
	//
	// Currently, `'eks'` is the only available option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-orchestrator
	//
	Orchestrator interface{} `field:"optional" json:"orchestrator" yaml:"orchestrator"`
	// The restricted instance groups of the SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-restrictedinstancegroups
	//
	RestrictedInstanceGroups interface{} `field:"optional" json:"restrictedInstanceGroups" yaml:"restrictedInstanceGroups"`
	// A tag object that consists of a key and an optional value, used to manage metadata for SageMaker AWS resources.
	//
	// You can add tags to notebook instances, training jobs, hyperparameter tuning jobs, batch transform jobs, models, labeling jobs, work teams, endpoint configurations, and endpoints. For more information on adding tags to SageMaker resources, see [AddTags](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AddTags.html) .
	//
	// For more information on adding metadata to your AWS resources with tagging, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) . For advice on best practices for managing AWS resources with tagging, see [Tagging Best Practices: Implement an Effective AWS Resource Tagging Strategy](https://docs.aws.amazon.com/https://d1.awsstatic.com/whitepapers/aws-tagging-best-practices.pdf) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Configuration for tiered storage in the SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-tieredstorageconfig
	//
	TieredStorageConfig interface{} `field:"optional" json:"tieredStorageConfig" yaml:"tieredStorageConfig"`
	// Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to.
	//
	// You can control access to and from your resources by configuring a VPC. For more information, see [Give SageMaker Access to Resources in your Amazon VPC](https://docs.aws.amazon.com/sagemaker/latest/dg/infrastructure-give-access.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

