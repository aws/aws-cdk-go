package awsopsworks


// Describes an Amazon EBS volume configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeConfigurationProperty := &VolumeConfigurationProperty{
//   	Encrypted: jsii.Boolean(false),
//   	Iops: jsii.Number(123),
//   	MountPoint: jsii.String("mountPoint"),
//   	NumberOfDisks: jsii.Number(123),
//   	RaidLevel: jsii.Number(123),
//   	Size: jsii.Number(123),
//   	VolumeType: jsii.String("volumeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html
//
type CfnLayer_VolumeConfigurationProperty struct {
	// Specifies whether an Amazon EBS volume is encrypted.
	//
	// For more information, see [Amazon EBS Encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volumeconfiguration-encrypted
	//
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The number of I/O operations per second (IOPS) to provision for the volume.
	//
	// For PIOPS volumes, the IOPS per disk.
	//
	// If you specify `io1` for the volume type, you must specify this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volumeconfiguration-iops
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The volume mount point.
	//
	// For example "/dev/sdh".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volumeconfiguration-mountpoint
	//
	MountPoint *string `field:"optional" json:"mountPoint" yaml:"mountPoint"`
	// The number of disks in the volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volumeconfiguration-numberofdisks
	//
	NumberOfDisks *float64 `field:"optional" json:"numberOfDisks" yaml:"numberOfDisks"`
	// The volume [RAID level](https://docs.aws.amazon.com/http://en.wikipedia.org/wiki/Standard_RAID_levels) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volumeconfiguration-raidlevel
	//
	RaidLevel *float64 `field:"optional" json:"raidLevel" yaml:"raidLevel"`
	// The volume size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volumeconfiguration-size
	//
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The volume type. For more information, see [Amazon EBS Volume Types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) .
	//
	// - `standard` - Magnetic. Magnetic volumes must have a minimum size of 1 GiB and a maximum size of 1024 GiB.
	// - `io1` - Provisioned IOPS (SSD). PIOPS volumes must have a minimum size of 4 GiB and a maximum size of 16384 GiB.
	// - `gp2` - General Purpose (SSD). General purpose volumes must have a minimum size of 1 GiB and a maximum size of 16384 GiB.
	// - `st1` - Throughput Optimized hard disk drive (HDD). Throughput optimized HDD volumes must have a minimum size of 500 GiB and a maximum size of 16384 GiB.
	// - `sc1` - Cold HDD. Cold HDD volumes must have a minimum size of 500 GiB and a maximum size of 16384 GiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volumeconfiguration-volumetype
	//
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

