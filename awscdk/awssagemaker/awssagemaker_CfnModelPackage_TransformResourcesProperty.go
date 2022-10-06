package awssagemaker


// Describes the resources, including ML instance types and ML instance count, to use for transform job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformResourcesProperty := &transformResourcesProperty{
//   	instanceCount: jsii.Number(123),
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   }
//
type CfnModelPackage_TransformResourcesProperty struct {
	// The number of ML compute instances to use in the transform job.
	//
	// For distributed transform jobs, specify a value greater than 1. The default value is `1` .
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// The ML compute instance type for the transform job.
	//
	// If you are using built-in algorithms to transform moderately sized datasets, we recommend using ml.m4.xlarge or `ml.m5.large` instance types.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt model data on the storage volume attached to the ML compute instance(s) that run the batch transform job.
	//
	// > Certain Nitro-based instances include local storage, dependent on the instance type. Local storage volumes are encrypted using a hardware module on the instance. You can't request a `VolumeKmsKeyId` when using an instance type with local storage.
	// >
	// > For a list of instance types that support local instance storage, see [Instance Store Volumes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#instance-store-volumes) .
	// >
	// > For more information about local instance storage encryption, see [SSD Instance Store Volumes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ssd-instance-store.html) .
	//
	// The `VolumeKmsKeyId` can be any of the following formats:
	//
	// - Key ID: `1234abcd-12ab-34cd-56ef-1234567890ab`
	// - Key ARN: `arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`
	// - Alias name: `alias/ExampleAlias`
	// - Alias name ARN: `arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias`.
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
}

