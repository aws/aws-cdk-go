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
//   			ThreadsPerCore: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ClusterName: jsii.String("clusterName"),
//   	NodeRecovery: jsii.String("nodeRecovery"),
//   	Orchestrator: &OrchestratorProperty{
//   		Eks: &ClusterOrchestratorEksConfigProperty{
//   			ClusterArn: jsii.String("clusterArn"),
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
	// The instance groups of the SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-instancegroups
	//
	InstanceGroups interface{} `field:"required" json:"instanceGroups" yaml:"instanceGroups"`
	// The name of the HyperPod Cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// If node auto-recovery is set to true, faulty nodes will be replaced or rebooted when a failure is detected.
	//
	// If set to false, nodes will be labelled when a fault is detected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-noderecovery
	//
	NodeRecovery *string `field:"optional" json:"nodeRecovery" yaml:"nodeRecovery"`
	// Specifies parameter(s) specific to the orchestrator, e.g. specify the EKS cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-orchestrator
	//
	Orchestrator interface{} `field:"optional" json:"orchestrator" yaml:"orchestrator"`
	// Custom tags for managing the SageMaker HyperPod cluster as an AWS resource.
	//
	// You can add tags to your cluster in the same way you add them in other AWS services that support tagging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to.
	//
	// You can control access to and from your resources by configuring a VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html#cfn-sagemaker-cluster-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

