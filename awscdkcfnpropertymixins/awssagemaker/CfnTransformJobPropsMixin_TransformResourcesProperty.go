package awssagemaker


// Describes the resources, including ML instance types and ML instance count, to use for the transform job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   transformResourcesProperty := &TransformResourcesProperty{
//   	InstanceCount: jsii.Number(123),
//   	InstanceType: jsii.String("instanceType"),
//   	VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-transformresources.html
//
type CfnTransformJobPropsMixin_TransformResourcesProperty struct {
	// The number of ML compute instances to use in the transform job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-transformresources.html#cfn-sagemaker-transformjob-transformresources-instancecount
	//
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// The ML compute instance type for the transform job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-transformresources.html#cfn-sagemaker-transformjob-transformresources-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The AWS KMS key that Amazon SageMaker uses to encrypt model data on the storage volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-transformresources.html#cfn-sagemaker-transformjob-transformresources-volumekmskeyid
	//
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
}

