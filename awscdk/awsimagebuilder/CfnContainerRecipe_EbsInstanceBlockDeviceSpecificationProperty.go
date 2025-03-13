package awsimagebuilder


// Amazon EBS-specific block device mapping specifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsInstanceBlockDeviceSpecificationProperty := &EbsInstanceBlockDeviceSpecificationProperty{
//   	DeleteOnTermination: jsii.Boolean(false),
//   	Encrypted: jsii.Boolean(false),
//   	Iops: jsii.Number(123),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	SnapshotId: jsii.String("snapshotId"),
//   	Throughput: jsii.Number(123),
//   	VolumeSize: jsii.Number(123),
//   	VolumeType: jsii.String("volumeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html
//
type CfnContainerRecipe_EbsInstanceBlockDeviceSpecificationProperty struct {
	// Use to configure delete on termination of the associated device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-deleteontermination
	//
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Use to configure device encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-encrypted
	//
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Use to configure device IOPS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-iops
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Use to configure the KMS key to use when encrypting the device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The snapshot that defines the device contents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-snapshotid
	//
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// *For GP3 volumes only* – The throughput in MiB/s that the volume supports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-throughput
	//
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// Use to override the device's volume size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-volumesize
	//
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Use to override the device's volume type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-volumetype
	//
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

