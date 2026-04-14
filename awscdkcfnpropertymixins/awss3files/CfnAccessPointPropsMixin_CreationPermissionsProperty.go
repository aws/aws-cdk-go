package awss3files


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   creationPermissionsProperty := &CreationPermissionsProperty{
//   	OwnerGid: jsii.String("ownerGid"),
//   	OwnerUid: jsii.String("ownerUid"),
//   	Permissions: jsii.String("permissions"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-accesspoint-creationpermissions.html
//
type CfnAccessPointPropsMixin_CreationPermissionsProperty struct {
	// Specifies the POSIX group ID to apply to the RootDirectory.
	//
	// Accepts values from 0 to 2^32 (4294967295).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-accesspoint-creationpermissions.html#cfn-s3files-accesspoint-creationpermissions-ownergid
	//
	OwnerGid *string `field:"optional" json:"ownerGid" yaml:"ownerGid"`
	// Specifies the POSIX user ID to apply to the RootDirectory.
	//
	// Accepts values from 0 to 2^32 (4294967295).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-accesspoint-creationpermissions.html#cfn-s3files-accesspoint-creationpermissions-owneruid
	//
	OwnerUid *string `field:"optional" json:"ownerUid" yaml:"ownerUid"`
	// Specifies the POSIX permissions to apply to the RootDirectory, in the format of an octal number representing the file's mode bits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-accesspoint-creationpermissions.html#cfn-s3files-accesspoint-creationpermissions-permissions
	//
	Permissions *string `field:"optional" json:"permissions" yaml:"permissions"`
}

