package awsimagebuilder


// The image recipe EBS instance block device specification includes the Amazon EBS-specific block device mapping specifications for the image.
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
type CfnImageRecipe_EbsInstanceBlockDeviceSpecificationProperty struct {
	// Configures delete on termination of the associated device.
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
	// Overrides the volume size of the device.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Overrides the volume type of the device.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

