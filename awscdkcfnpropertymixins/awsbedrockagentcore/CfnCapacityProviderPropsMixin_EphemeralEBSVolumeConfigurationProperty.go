package awsbedrockagentcore


// Parameters used to automatically set up EBS volumes when the instance is launched.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ephemeralEBSVolumeConfigurationProperty := &EphemeralEBSVolumeConfigurationProperty{
//   	EbsCardIndex: jsii.Number(123),
//   	Encrypted: jsii.Boolean(false),
//   	Iops: jsii.Number(123),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	SnapshotId: jsii.String("snapshotId"),
//   	Throughput: jsii.Number(123),
//   	VolumeInitializationRate: jsii.Number(123),
//   	VolumeSize: jsii.Number(123),
//   	VolumeType: jsii.String("volumeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration.html
//
type CfnCapacityProviderPropsMixin_EphemeralEBSVolumeConfigurationProperty struct {
	// The index of the EBS card.
	//
	// Applies to instances with multiple EBS cards.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration-ebscardindex
	//
	EbsCardIndex *float64 `field:"optional" json:"ebsCardIndex" yaml:"ebsCardIndex"`
	// Indicates whether the EBS volume is encrypted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration-encrypted
	//
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The number of I/O operations per second (IOPS).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration-iops
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Identifier of the customer managed KMS key to use for EBS encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The ID of the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration-snapshotid
	//
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The throughput to provision for a gp3 volume, in MiB/s.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration-throughput
	//
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The rate at which the volume is initialized after creation, in MiB/s.
	//
	// Supported only for volumes created from snapshots. If the snapshot is enabled for fast snapshot restore and a volume initialization rate is also specified, the volume is initialized at the specified rate instead of by fast snapshot restore. Valid range: 100-300 MiB/s.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration-volumeinitializationrate
	//
	VolumeInitializationRate *float64 `field:"optional" json:"volumeInitializationRate" yaml:"volumeInitializationRate"`
	// The size of the volume, in GiBs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration-volumesize
	//
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// The volume type.
	//
	// Defaults to gp3 if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-ephemeralebsvolumeconfiguration-volumetype
	//
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

