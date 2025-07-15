package awsfsx


// The FSx for OpenZFS file system user that is used for authorizing all file access requests that are made using the S3 access point.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openZFSPosixFileSystemUserProperty := &OpenZFSPosixFileSystemUserProperty{
//   	Gid: jsii.Number(123),
//   	Uid: jsii.Number(123),
//
//   	// the properties below are optional
//   	SecondaryGids: []interface{}{
//   		&FileSystemGIDProperty{
//   			Gid: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-openzfsposixfilesystemuser.html
//
type CfnS3AccessPointAttachment_OpenZFSPosixFileSystemUserProperty struct {
	// The GID of the file system user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-openzfsposixfilesystemuser.html#cfn-fsx-s3accesspointattachment-openzfsposixfilesystemuser-gid
	//
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// The UID of the file system user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-openzfsposixfilesystemuser.html#cfn-fsx-s3accesspointattachment-openzfsposixfilesystemuser-uid
	//
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
	// The list of secondary GIDs for the file system user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-openzfsposixfilesystemuser.html#cfn-fsx-s3accesspointattachment-openzfsposixfilesystemuser-secondarygids
	//
	SecondaryGids interface{} `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

