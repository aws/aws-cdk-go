package previewawsfsxmixins


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-ontapwindowsfilesystemuser.html#cfn-fsx-s3accesspointattachment-ontapwindowsfilesystemuser-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

