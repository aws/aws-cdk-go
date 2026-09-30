package awsstoragegateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVolume`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVolumeProps := &CfnVolumeProps{
//   	GatewayArn: jsii.String("gatewayArn"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	TargetName: jsii.String("targetName"),
//   	VolumeSizeInBytes: jsii.Number(123),
//
//   	// the properties below are optional
//   	KmsEncrypted: jsii.Boolean(false),
//   	KmsKey: jsii.String("kmsKey"),
//   	SnapshotId: jsii.String("snapshotId"),
//   	SourceVolumeArn: jsii.String("sourceVolumeArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-volume.html
//
type CfnVolumeProps struct {
	// The Amazon Resource Name (ARN) of the gateway on which to create the volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-volume.html#cfn-storagegateway-volume-gatewayarn
	//
	GatewayArn *string `field:"required" json:"gatewayArn" yaml:"gatewayArn"`
	// The network interface of the gateway on which to expose the iSCSI target.
	//
	// Only IPv4 addresses are accepted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-volume.html#cfn-storagegateway-volume-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The name of the iSCSI target used by an initiator to connect to a volume and used as a suffix for the target ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-volume.html#cfn-storagegateway-volume-targetname
	//
	TargetName *string `field:"required" json:"targetName" yaml:"targetName"`
	// The size of the volume in bytes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-volume.html#cfn-storagegateway-volume-volumesizeinbytes
	//
	VolumeSizeInBytes *float64 `field:"required" json:"volumeSizeInBytes" yaml:"volumeSizeInBytes"`
	// Set to true to use Amazon S3 server-side encryption with your own KMS key, or false to use a key managed by Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-volume.html#cfn-storagegateway-volume-kmsencrypted
	//
	KmsEncrypted interface{} `field:"optional" json:"kmsEncrypted" yaml:"kmsEncrypted"`
	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for Amazon S3 server-side encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-volume.html#cfn-storagegateway-volume-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The snapshot ID of the snapshot to restore as the new cached volume (e.g., snap-1122aabb).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-volume.html#cfn-storagegateway-volume-snapshotid
	//
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The ARN of an existing volume from which to create the new volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-volume.html#cfn-storagegateway-volume-sourcevolumearn
	//
	SourceVolumeArn *string `field:"optional" json:"sourceVolumeArn" yaml:"sourceVolumeArn"`
	// A list of up to 50 tags to assign to the volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-volume.html#cfn-storagegateway-volume-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

