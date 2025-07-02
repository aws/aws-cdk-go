package awsfsx


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-openzfsfilesystemidentity.html#cfn-fsx-s3accesspointattachment-openzfsfilesystemidentity-posixuser
	//
	PosixUser interface{} `field:"required" json:"posixUser" yaml:"posixUser"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-openzfsfilesystemidentity.html#cfn-fsx-s3accesspointattachment-openzfsfilesystemidentity-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

