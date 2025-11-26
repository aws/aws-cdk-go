package previewawsfsxmixins


// The GID of the file system user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fileSystemGIDProperty := &FileSystemGIDProperty{
//   	Gid: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-filesystemgid.html
//
type CfnS3AccessPointAttachmentPropsMixin_FileSystemGIDProperty struct {
	// The GID of the file system user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-filesystemgid.html#cfn-fsx-s3accesspointattachment-filesystemgid-gid
	//
	Gid *float64 `field:"optional" json:"gid" yaml:"gid"`
}

