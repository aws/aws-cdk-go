package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnComputeQuotaPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnComputeQuotaMixinProps := &CfnComputeQuotaMixinProps{
//   	ActivationState: jsii.String("activationState"),
//   	ClusterArn: jsii.String("clusterArn"),
//   	ComputeQuotaConfig: &ComputeQuotaConfigProperty{
//   		ComputeQuotaResources: []interface{}{
//   			&ComputeQuotaResourceConfigProperty{
//   				AcceleratorPartition: &AcceleratorPartitionConfigProperty{
//   					Count: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   				Accelerators: jsii.Number(123),
//   				Count: jsii.Number(123),
//   				InstanceType: jsii.String("instanceType"),
//   				MemoryInGiB: jsii.Number(123),
//   				VCpu: jsii.Number(123),
//   			},
//   		},
//   		PreemptTeamTasks: jsii.String("preemptTeamTasks"),
//   		ResourceSharingConfig: &ResourceSharingConfigProperty{
//   			AbsoluteBorrowLimits: []interface{}{
//   				&ComputeQuotaResourceConfigProperty{
//   					AcceleratorPartition: &AcceleratorPartitionConfigProperty{
//   						Count: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   					Accelerators: jsii.Number(123),
//   					Count: jsii.Number(123),
//   					InstanceType: jsii.String("instanceType"),
//   					MemoryInGiB: jsii.Number(123),
//   					VCpu: jsii.Number(123),
//   				},
//   			},
//   			BorrowLimit: jsii.Number(123),
//   			Strategy: jsii.String("strategy"),
//   		},
//   	},
//   	ComputeQuotaTarget: &ComputeQuotaTargetProperty{
//   		FairShareWeight: jsii.Number(123),
//   		TeamName: jsii.String("teamName"),
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
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
type CfnComputeQuotaMixinProps struct {
	// Whether the compute allocation rule is active.
	//
	// Defaults service-side to 'Enabled' when omitted on Create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html#cfn-sagemaker-computequota-activationstate
	//
	ActivationState *string `field:"optional" json:"activationState" yaml:"activationState"`
	// The ARN of the EKS-orchestrated SageMaker HyperPod cluster the compute allocation definition is bound to.
	//
	// Slurm-orchestrated clusters are rejected by the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html#cfn-sagemaker-computequota-clusterarn
	//
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// Compute allocation rules: per-resource quotas, idle-borrowing strategy, and team-task preemption policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html#cfn-sagemaker-computequota-computequotaconfig
	//
	ComputeQuotaConfig interface{} `field:"optional" json:"computeQuotaConfig" yaml:"computeQuotaConfig"`
	// The team this compute allocation rule applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html#cfn-sagemaker-computequota-computequotatarget
	//
	ComputeQuotaTarget interface{} `field:"optional" json:"computeQuotaTarget" yaml:"computeQuotaTarget"`
	// A free-form description of the compute allocation definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html#cfn-sagemaker-computequota-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the compute allocation definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html#cfn-sagemaker-computequota-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags to associate with the compute allocation definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-computequota.html#cfn-sagemaker-computequota-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

