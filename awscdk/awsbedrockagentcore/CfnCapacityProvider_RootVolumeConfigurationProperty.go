package awsbedrockagentcore


// Customer-facing configuration for the (service-managed) root volume.
//
// The service provisions the root volume at its own AMI size estimate plus FreeSpaceGiB, and pins the visible free space to FreeSpaceGiB with a filler file, so the space you are guaranteed does not change as the underlying AMI grows. The device name and the delete-on-termination behavior are service-owned and are not configurable.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rootVolumeConfigurationProperty := &RootVolumeConfigurationProperty{
//   	Encrypted: jsii.Boolean(false),
//   	FreeSpaceGiB: jsii.Number(123),
//   	Iops: jsii.Number(123),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Throughput: jsii.Number(123),
//   	VolumeType: jsii.String("volumeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-rootvolumeconfiguration.html
//
type CfnCapacityProvider_RootVolumeConfigurationProperty struct {
	// Indicates whether the EBS volume is encrypted.
	//
	// Encrypted volumes can only be attached to instances that support Amazon EBS encryption. If you are creating a volume from a snapshot, you can't specify an encryption value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-rootvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-rootvolumeconfiguration-encrypted
	//
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The free space guaranteed on the root volume, in GiB.
	//
	// The service adds the operating system overhead on top of this value. Defaults to 8 GiB. The maximum is below the 65,536 GiB gp3 ceiling because the service adds the AMI size bucket on top of this value, and the resulting total must still be a provisionable gp3 volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-rootvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-rootvolumeconfiguration-freespacegib
	//
	FreeSpaceGiB *float64 `field:"optional" json:"freeSpaceGiB" yaml:"freeSpaceGiB"`
	// The number of IOPS to provision.
	//
	// Only valid for gp3, io1, and io2 volumes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-rootvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-rootvolumeconfiguration-iops
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Identifier of the customer managed KMS key to use for EBS encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-rootvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-rootvolumeconfiguration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The throughput to provision for a gp3 volume, in MiB/s.
	//
	// Valid range: 125-2000 MiB/s.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-rootvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-rootvolumeconfiguration-throughput
	//
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The EBS volume type.
	//
	// Defaults to gp3 if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-rootvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-rootvolumeconfiguration-volumetype
	//
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

