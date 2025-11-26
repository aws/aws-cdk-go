package previewawssagemakermixins


// Configuration for the cluster used to run model monitoring jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clusterConfigProperty := &ClusterConfigProperty{
//   	InstanceCount: jsii.Number(123),
//   	InstanceType: jsii.String("instanceType"),
//   	VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   	VolumeSizeInGb: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-clusterconfig.html
//
type CfnMonitoringSchedulePropsMixin_ClusterConfigProperty struct {
	// The number of ML compute instances to use in the model monitoring job.
	//
	// For distributed processing jobs, specify a value greater than 1. The default value is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-clusterconfig.html#cfn-sagemaker-monitoringschedule-clusterconfig-instancecount
	//
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// The ML compute instance type for the processing job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-clusterconfig.html#cfn-sagemaker-monitoringschedule-clusterconfig-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the model monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-clusterconfig.html#cfn-sagemaker-monitoringschedule-clusterconfig-volumekmskeyid
	//
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
	// The size of the ML storage volume, in gigabytes, that you want to provision.
	//
	// You must specify sufficient ML storage for your scenario.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-clusterconfig.html#cfn-sagemaker-monitoringschedule-clusterconfig-volumesizeingb
	//
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

