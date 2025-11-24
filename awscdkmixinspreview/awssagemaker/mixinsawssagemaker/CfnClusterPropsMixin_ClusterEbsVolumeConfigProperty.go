package mixinsawssagemaker


// Defines the configuration for attaching an additional Amazon Elastic Block Store (EBS) volume to each instance of the SageMaker HyperPod cluster instance group.
//
// To learn more, see [SageMaker HyperPod release notes: June 20, 2024](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-hyperpod-release-notes.html#sagemaker-hyperpod-release-notes-20240620) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clusterEbsVolumeConfigProperty := &ClusterEbsVolumeConfigProperty{
//   	RootVolume: jsii.Boolean(false),
//   	VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   	VolumeSizeInGb: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterebsvolumeconfig.html
//
type CfnClusterPropsMixin_ClusterEbsVolumeConfigProperty struct {
	// Specifies whether the configuration is for the cluster's root or secondary Amazon EBS volume.
	//
	// You can specify two `ClusterEbsVolumeConfig` fields to configure both the root and secondary volumes. Set the value to `True` if you'd like to provide your own customer managed AWS KMS key to encrypt the root volume. When `True` :
	//
	// - The configuration is applied to the root volume.
	// - You can't specify the `VolumeSizeInGB` field. The size of the root volume is determined for you.
	// - You must specify a KMS key ID for `VolumeKmsKeyId` to encrypt the root volume with your own KMS key instead of an AWS owned KMS key.
	//
	// Otherwise, by default, the value is `False` , and the following applies:
	//
	// - The configuration is applied to the secondary volume, while the root volume is encrypted with an AWS owned key.
	// - You must specify the `VolumeSizeInGB` field.
	// - You can optionally specify the `VolumeKmsKeyId` to encrypt the secondary volume with your own KMS key instead of an AWS owned KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterebsvolumeconfig.html#cfn-sagemaker-cluster-clusterebsvolumeconfig-rootvolume
	//
	RootVolume interface{} `field:"optional" json:"rootVolume" yaml:"rootVolume"`
	// The ID of a KMS key to encrypt the Amazon EBS volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterebsvolumeconfig.html#cfn-sagemaker-cluster-clusterebsvolumeconfig-volumekmskeyid
	//
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
	// The size in gigabytes (GB) of the additional EBS volume to be attached to the instances in the SageMaker HyperPod cluster instance group.
	//
	// The additional EBS volume is attached to each instance within the SageMaker HyperPod cluster instance group and mounted to `/opt/sagemaker` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterebsvolumeconfig.html#cfn-sagemaker-cluster-clusterebsvolumeconfig-volumesizeingb
	//
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

