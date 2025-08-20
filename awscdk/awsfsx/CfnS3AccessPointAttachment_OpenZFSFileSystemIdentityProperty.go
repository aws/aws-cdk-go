package awsfsx


// Specifies the file system user identity that will be used for authorizing all file access requests that are made using the S3 access point.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openZFSFileSystemIdentityProperty := &OpenZFSFileSystemIdentityProperty{
//   	PosixUser: &OpenZFSPosixFileSystemUserProperty{
//   		Gid: jsii.Number(123),
//   		Uid: jsii.Number(123),
//
//   		// the properties below are optional
//   		SecondaryGids: []interface{}{
//   			&FileSystemGIDProperty{
//   				Gid: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-openzfsfilesystemidentity.html
//
type CfnS3AccessPointAttachment_OpenZFSFileSystemIdentityProperty struct {
	// Specifies the UID and GIDs of the file system POSIX user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-openzfsfilesystemidentity.html#cfn-fsx-s3accesspointattachment-openzfsfilesystemidentity-posixuser
	//
	PosixUser interface{} `field:"required" json:"posixUser" yaml:"posixUser"`
	// Specifies the FSx for OpenZFS user identity type, accepts only `POSIX` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-openzfsfilesystemidentity.html#cfn-fsx-s3accesspointattachment-openzfsfilesystemidentity-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

