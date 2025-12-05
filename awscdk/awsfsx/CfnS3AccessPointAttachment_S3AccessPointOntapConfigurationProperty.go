package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3AccessPointOntapConfigurationProperty := &S3AccessPointOntapConfigurationProperty{
//   	FileSystemIdentity: &OntapFileSystemIdentityProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
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
type CfnS3AccessPointAttachment_S3AccessPointOntapConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspointontapconfiguration.html#cfn-fsx-s3accesspointattachment-s3accesspointontapconfiguration-filesystemidentity
	//
	FileSystemIdentity interface{} `field:"required" json:"fileSystemIdentity" yaml:"fileSystemIdentity"`
	// The ID of the FSx for ONTAP volume that the S3 access point is attached to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspointontapconfiguration.html#cfn-fsx-s3accesspointattachment-s3accesspointontapconfiguration-volumeid
	//
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
}

