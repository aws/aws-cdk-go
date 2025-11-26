package previewawssagemakermixins


// Configuration for tiered storage in the SageMaker HyperPod cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tieredStorageConfigProperty := &TieredStorageConfigProperty{
//   	InstanceMemoryAllocationPercentage: jsii.Number(123),
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-tieredstorageconfig.html
//
type CfnClusterPropsMixin_TieredStorageConfigProperty struct {
	// The percentage of instance memory to allocate for tiered storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-tieredstorageconfig.html#cfn-sagemaker-cluster-tieredstorageconfig-instancememoryallocationpercentage
	//
	InstanceMemoryAllocationPercentage *float64 `field:"optional" json:"instanceMemoryAllocationPercentage" yaml:"instanceMemoryAllocationPercentage"`
	// The mode of tiered storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-tieredstorageconfig.html#cfn-sagemaker-cluster-tieredstorageconfig-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

