package awssagemaker


// Multi-instance GPU (MIG) partition profile and count for fractional accelerator allocation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   acceleratorPartitionConfigProperty := &AcceleratorPartitionConfigProperty{
//   	Count: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-acceleratorpartitionconfig.html
//
type CfnComputeQuotaPropsMixin_AcceleratorPartitionConfigProperty struct {
	// The number of MIG partitions of the specified Type allocated to the team.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-acceleratorpartitionconfig.html#cfn-sagemaker-computequota-acceleratorpartitionconfig-count
	//
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// MIG profile type.
	//
	// Service rejects MIG profiles that the chosen InstanceType does not support.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-acceleratorpartitionconfig.html#cfn-sagemaker-computequota-acceleratorpartitionconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

