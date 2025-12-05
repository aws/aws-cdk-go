package previewawsfsxmixins


// Properties for CfnS3AccessPointAttachmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policy interface{}
//
//   cfnS3AccessPointAttachmentMixinProps := &CfnS3AccessPointAttachmentMixinProps{
//   	Name: jsii.String("name"),
//   	OntapConfiguration: &S3AccessPointOntapConfigurationProperty{
//   		FileSystemIdentity: &OntapFileSystemIdentityProperty{
//   			Type: jsii.String("type"),
//   			UnixUser: &OntapUnixFileSystemUserProperty{
//   				Name: jsii.String("name"),
//   			},
//   			WindowsUser: &OntapWindowsFileSystemUserProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		VolumeId: jsii.String("volumeId"),
//   	},
//   	OpenZfsConfiguration: &S3AccessPointOpenZFSConfigurationProperty{
//   		FileSystemIdentity: &OpenZFSFileSystemIdentityProperty{
//   			PosixUser: &OpenZFSPosixFileSystemUserProperty{
//   				Gid: jsii.Number(123),
//   				SecondaryGids: []interface{}{
//   					&FileSystemGIDProperty{
//   						Gid: jsii.Number(123),
//   					},
//   				},
//   				Uid: jsii.Number(123),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   		VolumeId: jsii.String("volumeId"),
//   	},
//   	S3AccessPoint: &S3AccessPointProperty{
//   		Alias: jsii.String("alias"),
//   		Policy: policy,
//   		ResourceArn: jsii.String("resourceArn"),
//   		VpcConfiguration: &S3AccessPointVpcConfigurationProperty{
//   			VpcId: jsii.String("vpcId"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-s3accesspointattachment.html
//
type CfnS3AccessPointAttachmentMixinProps struct {
	// The name of the S3 access point attachment;
	//
	// also used for the name of the S3 access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-s3accesspointattachment.html#cfn-fsx-s3accesspointattachment-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-s3accesspointattachment.html#cfn-fsx-s3accesspointattachment-ontapconfiguration
	//
	OntapConfiguration interface{} `field:"optional" json:"ontapConfiguration" yaml:"ontapConfiguration"`
	// The OpenZFSConfiguration of the S3 access point attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-s3accesspointattachment.html#cfn-fsx-s3accesspointattachment-openzfsconfiguration
	//
	OpenZfsConfiguration interface{} `field:"optional" json:"openZfsConfiguration" yaml:"openZfsConfiguration"`
	// The S3 access point configuration of the S3 access point attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-s3accesspointattachment.html#cfn-fsx-s3accesspointattachment-s3accesspoint
	//
	S3AccessPoint interface{} `field:"optional" json:"s3AccessPoint" yaml:"s3AccessPoint"`
	// The type of Amazon FSx volume that the S3 access point is attached to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-s3accesspointattachment.html#cfn-fsx-s3accesspointattachment-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

