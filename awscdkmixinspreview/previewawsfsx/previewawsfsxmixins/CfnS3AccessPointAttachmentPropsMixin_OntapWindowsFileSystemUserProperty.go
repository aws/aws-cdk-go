package previewawsfsxmixins


// The FSx for ONTAP Windows file system user that is used for authorizing all file access requests that are made using the S3 access point.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ontapWindowsFileSystemUserProperty := &OntapWindowsFileSystemUserProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-ontapwindowsfilesystemuser.html
//
type CfnS3AccessPointAttachmentPropsMixin_OntapWindowsFileSystemUserProperty struct {
	// The name of the Windows user.
	//
	// The name can be up to 256 characters long and supports Active Directory users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-ontapwindowsfilesystemuser.html#cfn-fsx-s3accesspointattachment-ontapwindowsfilesystemuser-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

