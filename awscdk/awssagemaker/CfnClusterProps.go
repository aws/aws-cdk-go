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
//   cfnClusterProps := &CfnClusterProps{
//   	ClusterName: jsii.String("clusterName"),
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
//   			CurrentCount: jsii.Number(123),
//   			InstanceStorageConfigs: []interface{}{
//   				&ClusterInstanceStorageConfigProperty{
//   					EbsVolumeConfig: &ClusterEbsVolumeConfigProperty{
//   						VolumeSizeInGb: jsii.Number(123),
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
//   	NodeRecovery: jsii.String("nodeRecovery"),
//   	Orchestrator: &OrchestratorProperty{
//   		Eks: &ClusterOrchestratorEksConfigProperty{
//   			ClusterArn: jsii.String("clusterArn"),
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
//   						VolumeSizeInGb: jsii.Number(123),
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
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
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
	// The name of the SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The instance groups of the SageMaker HyperPod cluster.
	//
	// To delete an instance group, remove it from the array.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-instancegroups
	//
	InstanceGroups interface{} `field:"optional" json:"instanceGroups" yaml:"instanceGroups"`
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
	// Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to.
	//
	// You can control access to and from your resources by configuring a VPC. For more information, see [Give SageMaker Access to Resources in your Amazon VPC](https://docs.aws.amazon.com/sagemaker/latest/dg/infrastructure-give-access.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

