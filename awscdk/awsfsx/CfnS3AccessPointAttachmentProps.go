package awsfsx


// Properties for defining a `CfnS3AccessPointAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnS3AccessPointAttachmentProps := &CfnS3AccessPointAttachmentProps{
//   	Name: jsii.String("name"),
//   	OpenZfsConfiguration: &S3AccessPointOpenZFSConfigurationProperty{
//   		FileSystemIdentity: &OpenZFSFileSystemIdentityProperty{
//   			PosixUser: &OpenZFSPosixFileSystemUserProperty{
//   				Gid: jsii.Number(123),
//   				Uid: jsii.Number(123),
//
//   				// the properties below are optional
//   				SecondaryGids: []interface{}{
//   					&FileSystemGIDProperty{
//   						Gid: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   		VolumeId: jsii.String("volumeId"),
//   	},
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	S3AccessPoint: &S3AccessPointProperty{
//   		Alias: jsii.String("alias"),
//   		Policy: policy,
//   		ResourceArn: jsii.String("resourceArn"),
//   		VpcConfiguration: &S3AccessPointVpcConfigurationProperty{
//   			VpcId: jsii.String("vpcId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-s3accesspointattachment.html
//
type CfnS3AccessPointAttachmentProps struct {
	// The name of the S3 access point attachment;
	//
	// also used for the name of the S3 access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-s3accesspointattachment.html#cfn-fsx-s3accesspointattachment-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The OpenZFSConfiguration of the S3 access point attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-s3accesspointattachment.html#cfn-fsx-s3accesspointattachment-openzfsconfiguration
	//
	OpenZfsConfiguration interface{} `field:"required" json:"openZfsConfiguration" yaml:"openZfsConfiguration"`
	// The type of Amazon FSx volume that the S3 access point is attached to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-s3accesspointattachment.html#cfn-fsx-s3accesspointattachment-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The S3 access point configuration of the S3 access point attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-s3accesspointattachment.html#cfn-fsx-s3accesspointattachment-s3accesspoint
	//
	S3AccessPoint interface{} `field:"optional" json:"s3AccessPoint" yaml:"s3AccessPoint"`
}

