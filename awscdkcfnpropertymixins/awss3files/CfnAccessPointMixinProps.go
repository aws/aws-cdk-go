package awss3files


// Properties for CfnAccessPointPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAccessPointMixinProps := &CfnAccessPointMixinProps{
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
//   		CreationPermissions: &CreationPermissionsProperty{
//   			OwnerGid: jsii.String("ownerGid"),
//   			OwnerUid: jsii.String("ownerUid"),
//   			Permissions: jsii.String("permissions"),
//   		},
//   		Path: jsii.String("path"),
//   	},
//   	Tags: []AccessPointTagProperty{
//   		&AccessPointTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-accesspoint.html
//
type CfnAccessPointMixinProps struct {
	// (optional) A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-accesspoint.html#cfn-s3files-accesspoint-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// The ID of the S3 Files file system that the access point provides access to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-accesspoint.html#cfn-s3files-accesspoint-filesystemid
	//
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-accesspoint.html#cfn-s3files-accesspoint-posixuser
	//
	PosixUser interface{} `field:"optional" json:"posixUser" yaml:"posixUser"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-accesspoint.html#cfn-s3files-accesspoint-rootdirectory
	//
	RootDirectory interface{} `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-accesspoint.html#cfn-s3files-accesspoint-tags
	//
	Tags *[]*CfnAccessPointPropsMixin_AccessPointTagProperty `field:"optional" json:"tags" yaml:"tags"`
}

