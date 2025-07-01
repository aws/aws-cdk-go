package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3AccessPointOpenZFSConfigurationProperty := &S3AccessPointOpenZFSConfigurationProperty{
//   	FileSystemIdentity: &OpenZFSFileSystemIdentityProperty{
//   		PosixUser: &OpenZFSPosixFileSystemUserProperty{
//   			Gid: jsii.Number(123),
//   			Uid: jsii.Number(123),
//
//   			// the properties below are optional
//   			SecondaryGids: []interface{}{
//   				&FileSystemGIDProperty{
//   					Gid: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	VolumeId: jsii.String("volumeId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspointopenzfsconfiguration.html
//
type CfnS3AccessPointAttachment_S3AccessPointOpenZFSConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspointopenzfsconfiguration.html#cfn-fsx-s3accesspointattachment-s3accesspointopenzfsconfiguration-filesystemidentity
	//
	FileSystemIdentity interface{} `field:"required" json:"fileSystemIdentity" yaml:"fileSystemIdentity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspointopenzfsconfiguration.html#cfn-fsx-s3accesspointattachment-s3accesspointopenzfsconfiguration-volumeid
	//
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
}

