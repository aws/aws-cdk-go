package awssagemaker


// Configuration for tiered storage in the SageMaker HyperPod cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tieredStorageConfigProperty := &TieredStorageConfigProperty{
//   	Mode: jsii.String("mode"),
//
//   	// the properties below are optional
//   	InstanceMemoryAllocationPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-tieredstorageconfig.html
//
type CfnCluster_TieredStorageConfigProperty struct {
	// The mode of tiered storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-tieredstorageconfig.html#cfn-sagemaker-cluster-tieredstorageconfig-mode
	//
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The percentage of instance memory to allocate for tiered storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-tieredstorageconfig.html#cfn-sagemaker-cluster-tieredstorageconfig-instancememoryallocationpercentage
	//
	InstanceMemoryAllocationPercentage *float64 `field:"optional" json:"instanceMemoryAllocationPercentage" yaml:"instanceMemoryAllocationPercentage"`
}

