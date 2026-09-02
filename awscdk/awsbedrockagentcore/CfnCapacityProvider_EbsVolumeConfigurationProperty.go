package awsbedrockagentcore


// Configuration for an EBS-backed persistent volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsVolumeConfigurationProperty := &EbsVolumeConfigurationProperty{
//   	Name: jsii.String("name"),
//   	SizeGiB: jsii.Number(123),
//
//   	// the properties below are optional
//   	Encrypted: jsii.Boolean(false),
//   	Iops: jsii.Number(123),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	SnapshotId: jsii.String("snapshotId"),
//   	Throughput: jsii.Number(123),
//   	VolumeType: jsii.String("volumeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ebsvolumeconfiguration.html
//
type CfnCapacityProvider_EbsVolumeConfigurationProperty struct {
	// The logical name of the volume, used to reference it when mounting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ebsvolumeconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The size of the volume in GiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ebsvolumeconfiguration-sizegib
	//
	SizeGiB *float64 `field:"required" json:"sizeGiB" yaml:"sizeGiB"`
	// Whether to encrypt the volume.
	//
	// Defaults to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ebsvolumeconfiguration-encrypted
	//
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The number of IOPS to provision.
	//
	// Only valid for gp3, io1, and io2 volumes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ebsvolumeconfiguration-iops
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Identifier of the KMS key to use for encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ebsvolumeconfiguration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Optional EBS snapshot ID to initialize the volume from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ebsvolumeconfiguration-snapshotid
	//
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The throughput in MiB/s.
	//
	// Only valid for gp3 volumes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ebsvolumeconfiguration-throughput
	//
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The EBS volume type.
	//
	// Defaults to gp3 if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ebsvolumeconfiguration-volumetype
	//
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

