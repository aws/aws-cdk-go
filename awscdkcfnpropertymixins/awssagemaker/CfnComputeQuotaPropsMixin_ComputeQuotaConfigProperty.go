package awssagemaker


// Compute allocation rules: per-resource quotas, idle-borrowing strategy, and team-task preemption policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   computeQuotaConfigProperty := &ComputeQuotaConfigProperty{
//   	ComputeQuotaResources: []interface{}{
//   		&ComputeQuotaResourceConfigProperty{
//   			AcceleratorPartition: &AcceleratorPartitionConfigProperty{
//   				Count: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   			Accelerators: jsii.Number(123),
//   			Count: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			MemoryInGiB: jsii.Number(123),
//   			VCpu: jsii.Number(123),
//   		},
//   	},
//   	PreemptTeamTasks: jsii.String("preemptTeamTasks"),
//   	ResourceSharingConfig: &ResourceSharingConfigProperty{
//   		AbsoluteBorrowLimits: []interface{}{
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
//   		BorrowLimit: jsii.Number(123),
//   		Strategy: jsii.String("strategy"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-computequotaconfig.html
//
type CfnComputeQuotaPropsMixin_ComputeQuotaConfigProperty struct {
	// Per-instance-type compute quotas for the team.
	//
	// When VCpu / MemoryInGiB / Accelerators are partially specified, SageMaker auto-fills the remaining fields ratio-proportionally to the underlying instance — Read may return values the customer did not supply (drift may flag this).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-computequotaconfig.html#cfn-sagemaker-computequota-computequotaconfig-computequotaresources
	//
	ComputeQuotaResources interface{} `field:"optional" json:"computeQuotaResources" yaml:"computeQuotaResources"`
	// Team-task preemption policy.
	//
	// Defaults service-side to 'LowerPriority' when omitted on Create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-computequotaconfig.html#cfn-sagemaker-computequota-computequotaconfig-preemptteamtasks
	//
	PreemptTeamTasks *string `field:"optional" json:"preemptTeamTasks" yaml:"preemptTeamTasks"`
	// Idle-resource-sharing strategy for this team's quota.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-computequotaconfig.html#cfn-sagemaker-computequota-computequotaconfig-resourcesharingconfig
	//
	ResourceSharingConfig interface{} `field:"optional" json:"resourceSharingConfig" yaml:"resourceSharingConfig"`
}

