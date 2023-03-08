package awsimagebuilder


// Amazon EBS-specific block device mapping specifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsInstanceBlockDeviceSpecificationProperty := &ebsInstanceBlockDeviceSpecificationProperty{
//   	deleteOnTermination: jsii.Boolean(false),
//   	encrypted: jsii.Boolean(false),
//   	iops: jsii.Number(123),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	snapshotId: jsii.String("snapshotId"),
//   	throughput: jsii.Number(123),
//   	volumeSize: jsii.Number(123),
//   	volumeType: jsii.String("volumeType"),
//   }
//
type CfnContainerRecipe_EbsInstanceBlockDeviceSpecificationProperty struct {
	// Use to configure delete on termination of the associated device.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Use to configure device encryption.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Use to configure device IOPS.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Use to configure the KMS key to use when encrypting the device.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The snapshot that defines the device contents.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// *For GP3 volumes only* – The throughput in MiB/s that the volume supports.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// Use to override the device's volume size.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Use to override the device's volume type.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

