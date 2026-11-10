package awssagemaker


// Per-instance-type compute quota element.
//
// Reused under ComputeQuotaConfig.ComputeQuotaResources and ResourceSharingConfig.AbsoluteBorrowLimits.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeQuotaResourceConfigProperty := &ComputeQuotaResourceConfigProperty{
//   	InstanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	AcceleratorPartition: &AcceleratorPartitionConfigProperty{
//   		Count: jsii.Number(123),
//   		Type: jsii.String("type"),
//   	},
//   	Accelerators: jsii.Number(123),
//   	Count: jsii.Number(123),
//   	MemoryInGiB: jsii.Number(123),
//   	VCpu: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-computequotaresourceconfig.html
//
type CfnComputeQuota_ComputeQuotaResourceConfigProperty struct {
	// The cluster instance type the quota applies to (e.g., 'ml.m5.xlarge', 'ml.p5.48xlarge'). Validated by the service against the supported HyperPod instance types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-computequotaresourceconfig.html#cfn-sagemaker-computequota-computequotaresourceconfig-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Multi-instance GPU (MIG) partition profile and count for fractional accelerator allocation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-computequotaresourceconfig.html#cfn-sagemaker-computequota-computequotaresourceconfig-acceleratorpartition
	//
	AcceleratorPartition interface{} `field:"optional" json:"acceleratorPartition" yaml:"acceleratorPartition"`
	// The number of accelerators (e.g., GPUs) of this type allocated to the team. SageMaker auto-fills VCpu / MemoryInGiB ratio-proportionally if only this is set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-computequotaresourceconfig.html#cfn-sagemaker-computequota-computequotaresourceconfig-accelerators
	//
	Accelerators *float64 `field:"optional" json:"accelerators" yaml:"accelerators"`
	// The number of instances of this type allocated to the team.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-computequotaresourceconfig.html#cfn-sagemaker-computequota-computequotaresourceconfig-count
	//
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The memory in GiB allocated to the team.
	//
	// Floating-point. SageMaker auto-fills Accelerators / VCpu ratio-proportionally if only this is set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-computequotaresourceconfig.html#cfn-sagemaker-computequota-computequotaresourceconfig-memoryingib
	//
	MemoryInGiB *float64 `field:"optional" json:"memoryInGiB" yaml:"memoryInGiB"`
	// The number of vCPUs allocated to the team.
	//
	// Floating-point. SageMaker auto-fills Accelerators / MemoryInGiB ratio-proportionally if only this is set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-computequotaresourceconfig.html#cfn-sagemaker-computequota-computequotaresourceconfig-vcpu
	//
	VCpu *float64 `field:"optional" json:"vCpu" yaml:"vCpu"`
}

