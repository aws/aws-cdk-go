package previewawsefsmixins


// Properties for CfnAccessPointPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAccessPointMixinProps := &CfnAccessPointMixinProps{
//   	AccessPointTags: []AccessPointTagProperty{
//   		&AccessPointTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ClientToken: jsii.String("clientToken"),
//   	FileSystemId: jsii.String("fileSystemId"),
//   	PosixUser: &PosixUserProperty{
//   		Gid: jsii.String("gid"),
//   		SecondaryGids: []*string{
//   			jsii.String("secondaryGids"),
//   		},
//   		Uid: jsii.String("uid"),
//   	},
//   	RootDirectory: &RootDirectoryProperty{
//   		CreationInfo: &CreationInfoProperty{
//   			OwnerGid: jsii.String("ownerGid"),
//   			OwnerUid: jsii.String("ownerUid"),
//   			Permissions: jsii.String("permissions"),
//   		},
//   		Path: jsii.String("path"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-accesspoint.html
//
type CfnAccessPointMixinProps struct {
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-accesspoint.html#cfn-efs-accesspoint-accesspointtags
	//
	AccessPointTags *[]*CfnAccessPointPropsMixin_AccessPointTagProperty `field:"optional" json:"accessPointTags" yaml:"accessPointTags"`
	// The opaque string specified in the request to ensure idempotent creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-accesspoint.html#cfn-efs-accesspoint-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// The ID of the EFS file system that the access point applies to.
	//
	// Accepts only the ID format for input when specifying a file system, for example `fs-0123456789abcedf2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-accesspoint.html#cfn-efs-accesspoint-filesystemid
	//
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// The full POSIX identity, including the user ID, group ID, and secondary group IDs on the access point that is used for all file operations by NFS clients using the access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-accesspoint.html#cfn-efs-accesspoint-posixuser
	//
	PosixUser interface{} `field:"optional" json:"posixUser" yaml:"posixUser"`
	// The directory on the EFS file system that the access point exposes as the root directory to NFS clients using the access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-accesspoint.html#cfn-efs-accesspoint-rootdirectory
	//
	RootDirectory interface{} `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
}

