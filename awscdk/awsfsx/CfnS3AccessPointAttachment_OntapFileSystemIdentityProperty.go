package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ontapFileSystemIdentityProperty := &OntapFileSystemIdentityProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
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
type CfnS3AccessPointAttachment_OntapFileSystemIdentityProperty struct {
	// Specifies the FSx for ONTAP user identity type, accepts either UNIX or WINDOWS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-ontapfilesystemidentity.html#cfn-fsx-s3accesspointattachment-ontapfilesystemidentity-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-ontapfilesystemidentity.html#cfn-fsx-s3accesspointattachment-ontapfilesystemidentity-unixuser
	//
	UnixUser interface{} `field:"optional" json:"unixUser" yaml:"unixUser"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-ontapfilesystemidentity.html#cfn-fsx-s3accesspointattachment-ontapfilesystemidentity-windowsuser
	//
	WindowsUser interface{} `field:"optional" json:"windowsUser" yaml:"windowsUser"`
}

