package mixinsawsfsx


// The FSx for OpenZFS file system user that is used for authorizing all file access requests that are made using the S3 access point.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   openZFSPosixFileSystemUserProperty := &OpenZFSPosixFileSystemUserProperty{
//   	Gid: jsii.Number(123),
//   	SecondaryGids: []interface{}{
//   		&FileSystemGIDProperty{
//   			Gid: jsii.Number(123),
//   		},
//   	},
//   	Uid: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-openzfsposixfilesystemuser.html
//
type CfnS3AccessPointAttachmentPropsMixin_OpenZFSPosixFileSystemUserProperty struct {
	// The GID of the file system user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-openzfsposixfilesystemuser.html#cfn-fsx-s3accesspointattachment-openzfsposixfilesystemuser-gid
	//
	Gid *float64 `field:"optional" json:"gid" yaml:"gid"`
	// The list of secondary GIDs for the file system user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-openzfsposixfilesystemuser.html#cfn-fsx-s3accesspointattachment-openzfsposixfilesystemuser-secondarygids
	//
	SecondaryGids interface{} `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
	// The UID of the file system user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-openzfsposixfilesystemuser.html#cfn-fsx-s3accesspointattachment-openzfsposixfilesystemuser-uid
	//
	Uid *float64 `field:"optional" json:"uid" yaml:"uid"`
}

