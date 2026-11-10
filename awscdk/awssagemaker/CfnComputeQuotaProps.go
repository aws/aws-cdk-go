package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnComputeQuota`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnComputeQuotaProps := &CfnComputeQuotaProps{
//   	ClusterArn: jsii.String("clusterArn"),
//   	ComputeQuotaConfig: &ComputeQuotaConfigProperty{
//   		ComputeQuotaResources: []interface{}{
//   			&ComputeQuotaResourceConfigProperty{
//   				InstanceType: jsii.String("instanceType"),
//
//   				// the properties below are optional
//   				AcceleratorPartition: &AcceleratorPartitionConfigProperty{
//   					Count: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   				Accelerators: jsii.Number(123),
//   				Count: jsii.Number(123),
//   				MemoryInGiB: jsii.Number(123),
//   				VCpu: jsii.Number(123),
//   			},
//   		},
//   		PreemptTeamTasks: jsii.String("preemptTeamTasks"),
//   		ResourceSharingConfig: &ResourceSharingConfigProperty{
//   			Strategy: jsii.String("strategy"),
//
//   			// the properties below are optional
//   			AbsoluteBorrowLimits: []interface{}{
//   				&ComputeQuotaResourceConfigProperty{
//   					InstanceType: jsii.String("instanceType"),
//
//   					// the properties below are optional
//   					AcceleratorPartition: &AcceleratorPartitionConfigProperty{
//   						Count: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   					Accelerators: jsii.Number(123),
//   					Count: jsii.Number(123),
//   					MemoryInGiB: jsii.Number(123),
//   					VCpu: jsii.Number(123),
//   				},
//   			},
//   			BorrowLimit: jsii.Number(123),
//   		},
//   	},
//   	ComputeQuotaTarget: &ComputeQuotaTargetProperty{
//   		TeamName: jsii.String("teamName"),
//
//   		// the properties below are optional
//   		FairShareWeight: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ActivationState: jsii.String("activationState"),
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html
//
type CfnComputeQuotaProps struct {
	// The ARN of the EKS-orchestrated SageMaker HyperPod cluster the compute allocation definition is bound to.
	//
	// Slurm-orchestrated clusters are rejected by the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html#cfn-sagemaker-computequota-clusterarn
	//
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// Compute allocation rules: per-resource quotas, idle-borrowing strategy, and team-task preemption policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html#cfn-sagemaker-computequota-computequotaconfig
	//
	ComputeQuotaConfig interface{} `field:"required" json:"computeQuotaConfig" yaml:"computeQuotaConfig"`
	// The team this compute allocation rule applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html#cfn-sagemaker-computequota-computequotatarget
	//
	ComputeQuotaTarget interface{} `field:"required" json:"computeQuotaTarget" yaml:"computeQuotaTarget"`
	// The name of the compute allocation definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html#cfn-sagemaker-computequota-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether the compute allocation rule is active.
	//
	// Defaults service-side to 'Enabled' when omitted on Create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html#cfn-sagemaker-computequota-activationstate
	//
	ActivationState *string `field:"optional" json:"activationState" yaml:"activationState"`
	// A free-form description of the compute allocation definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html#cfn-sagemaker-computequota-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags to associate with the compute allocation definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html#cfn-sagemaker-computequota-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

