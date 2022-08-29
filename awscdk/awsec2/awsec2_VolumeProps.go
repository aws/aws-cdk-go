package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties of an EBS Volume.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var instance instance
//   var role role
//
//
//   volume := ec2.NewVolume(this, jsii.String("Volume"), &volumeProps{
//   	availabilityZone: jsii.String("us-west-2a"),
//   	size: awscdk.Size.gibibytes(jsii.Number(500)),
//   	encrypted: jsii.Boolean(true),
//   })
//
//   volume.grantAttachVolume(role, []iInstance{
//   	instance,
//   })
//
type VolumeProps struct {
	// The Availability Zone in which to create the volume.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Indicates whether the volume is auto-enabled for I/O operations.
	//
	// By default, Amazon EBS disables I/O to the volume from attached EC2
	// instances when it determines that a volume's data is potentially inconsistent. If the consistency of the volume is not a concern, and
	// you prefer that the volume be made available immediately if it's impaired, you can configure the volume to automatically enable I/O.
	AutoEnableIo *bool `field:"optional" json:"autoEnableIo" yaml:"autoEnableIo"`
	// Indicates whether Amazon EBS Multi-Attach is enabled.
	//
	// See {@link https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-volumes-multi.html#considerations|Considerations and limitations}
	// for the constraints of multi-attach.
	EnableMultiAttach *bool `field:"optional" json:"enableMultiAttach" yaml:"enableMultiAttach"`
	// Specifies whether the volume should be encrypted.
	//
	// The effect of setting the encryption state to true depends on the volume origin
	// (new or from a snapshot), starting encryption state, ownership, and whether encryption by default is enabled. For more information,
	// see {@link https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-by-default|Encryption by Default}
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// Encrypted Amazon EBS volumes must be attached to instances that support Amazon EBS encryption. For more information, see
	// {@link https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances|Supported Instance Types.}
	Encrypted *bool `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The customer-managed encryption key that is used to encrypt the Volume.
	//
	// The encrypted property must
	// be true if this is provided.
	//
	// Note: If using an {@link aws-kms.IKey} created from a {@link aws-kms.Key.fromKeyArn()} here,
	// then the KMS key **must** have the following in its Key policy; otherwise, the Volume
	// will fail to create.
	//
	//      {
	//        "Effect": "Allow",
	//        "Principal": { "AWS": "<arn for your account-user> ex: arn:aws:iam::00000000000:root" },
	//        "Resource": "*",
	//        "Action": [
	//          "kms:DescribeKey",
	//          "kms:GenerateDataKeyWithoutPlainText",
	//        ],
	//        "Condition": {
	//          "StringEquals": {
	//            "kms:ViaService": "ec2.<Region>.amazonaws.com", (eg: ec2.us-east-1.amazonaws.com)
	//            "kms:CallerAccount": "0000000000" (your account ID)
	//          }
	//        }
	// }.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The number of I/O operations per second (IOPS) to provision for the volume.
	//
	// The maximum ratio is 50 IOPS/GiB for PROVISIONED_IOPS_SSD,
	// and 500 IOPS/GiB for both PROVISIONED_IOPS_SSD_IO2 and GENERAL_PURPOSE_SSD_GP3.
	// See {@link https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html}
	// for more information.
	//
	// This parameter is valid only for PROVISIONED_IOPS_SSD, PROVISIONED_IOPS_SSD_IO2 and GENERAL_PURPOSE_SSD_GP3 volumes.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Policy to apply when the volume is removed from the stack.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The size of the volume, in GiBs.
	//
	// You must specify either a snapshot ID or a volume size.
	// See {@link https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html}
	// for details on the allowable size for each type of volume.
	Size awscdk.Size `field:"optional" json:"size" yaml:"size"`
	// The snapshot from which to create the volume.
	//
	// You must specify either a snapshot ID or a volume size.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The value of the physicalName property of this resource.
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
	// The type of the volume;
	//
	// what type of storage to use to form the EBS Volume.
	VolumeType EbsDeviceVolumeType `field:"optional" json:"volumeType" yaml:"volumeType"`
}

