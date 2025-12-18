package previewawsfsxmixins


// Specifies the file system user identity that will be used for authorizing all file access requests that are made using the S3 access point.
//
// The identity can be either a UNIX user or a Windows user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ontapFileSystemIdentityProperty := &OntapFileSystemIdentityProperty{
//   	Type: jsii.String("type"),
//   	UnixUser: &OntapUnixFileSystemUserProperty{
//   		Name: jsii.String("name"),
//   	},
//   	WindowsUser: &OntapWindowsFileSystemUserProperty{
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-ontapfilesystemidentity.html
//
type CfnS3AccessPointAttachmentPropsMixin_OntapFileSystemIdentityProperty struct {
	// Specifies the FSx for ONTAP user identity type.
	//
	// Valid values are `UNIX` and `WINDOWS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-ontapfilesystemidentity.html#cfn-fsx-s3accesspointattachment-ontapfilesystemidentity-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Specifies the UNIX user identity for file system operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-ontapfilesystemidentity.html#cfn-fsx-s3accesspointattachment-ontapfilesystemidentity-unixuser
	//
	UnixUser interface{} `field:"optional" json:"unixUser" yaml:"unixUser"`
	// Specifies the Windows user identity for file system operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-ontapfilesystemidentity.html#cfn-fsx-s3accesspointattachment-ontapfilesystemidentity-windowsuser
	//
	WindowsUser interface{} `field:"optional" json:"windowsUser" yaml:"windowsUser"`
}

