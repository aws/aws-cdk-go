package awsefs


// Required if the `RootDirectory` > `Path` specified does not exist.
//
// Specifies the POSIX IDs and permissions to apply to the access point's `RootDirectory` > `Path` . If the access point root directory does not exist, EFS creates it with these settings when a client connects to the access point. When specifying `CreationInfo` , you must include values for all properties.
//
// Amazon EFS creates a root directory only if you have provided the CreationInfo: OwnUid, OwnGID, and permissions for the directory. If you do not provide this information, Amazon EFS does not create the root directory. If the root directory does not exist, attempts to mount using the access point will fail.
//
// > If you do not provide `CreationInfo` and the specified `RootDirectory` does not exist, attempts to mount the file system using the access point will fail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   creationInfoProperty := &CreationInfoProperty{
//   	OwnerGid: jsii.String("ownerGid"),
//   	OwnerUid: jsii.String("ownerUid"),
//   	Permissions: jsii.String("permissions"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-creationinfo.html
//
type CfnAccessPoint_CreationInfoProperty struct {
	// Specifies the POSIX group ID to apply to the `RootDirectory` .
	//
	// Accepts values from 0 to 2^32 (4294967295).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-creationinfo.html#cfn-efs-accesspoint-creationinfo-ownergid
	//
	OwnerGid *string `field:"required" json:"ownerGid" yaml:"ownerGid"`
	// Specifies the POSIX user ID to apply to the `RootDirectory` .
	//
	// Accepts values from 0 to 2^32 (4294967295).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-creationinfo.html#cfn-efs-accesspoint-creationinfo-owneruid
	//
	OwnerUid *string `field:"required" json:"ownerUid" yaml:"ownerUid"`
	// Specifies the POSIX permissions to apply to the `RootDirectory` , in the format of an octal number representing the file's mode bits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-creationinfo.html#cfn-efs-accesspoint-creationinfo-permissions
	//
	Permissions *string `field:"required" json:"permissions" yaml:"permissions"`
}

