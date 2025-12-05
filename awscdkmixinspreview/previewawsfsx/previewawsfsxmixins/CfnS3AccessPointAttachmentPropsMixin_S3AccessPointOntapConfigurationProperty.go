package previewawsfsxmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3AccessPointOntapConfigurationProperty := &S3AccessPointOntapConfigurationProperty{
//   	FileSystemIdentity: &OntapFileSystemIdentityProperty{
//   		Type: jsii.String("type"),
//   		UnixUser: &OntapUnixFileSystemUserProperty{
//   			Name: jsii.String("name"),
//   		},
//   		WindowsUser: &OntapWindowsFileSystemUserProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	VolumeId: jsii.String("volumeId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspointontapconfiguration.html
//
type CfnS3AccessPointAttachmentPropsMixin_S3AccessPointOntapConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspointontapconfiguration.html#cfn-fsx-s3accesspointattachment-s3accesspointontapconfiguration-filesystemidentity
	//
	FileSystemIdentity interface{} `field:"optional" json:"fileSystemIdentity" yaml:"fileSystemIdentity"`
	// The ID of the FSx for ONTAP volume that the S3 access point is attached to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspointontapconfiguration.html#cfn-fsx-s3accesspointattachment-s3accesspointontapconfiguration-volumeid
	//
	VolumeId *string `field:"optional" json:"volumeId" yaml:"volumeId"`
}

