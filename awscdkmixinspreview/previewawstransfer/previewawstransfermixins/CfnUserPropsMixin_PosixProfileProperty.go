package previewawstransfermixins


// The full POSIX identity, including user ID ( `Uid` ), group ID ( `Gid` ), and any secondary groups IDs ( `SecondaryGids` ), that controls your users' access to your Amazon EFS file systems.
//
// The POSIX permissions that are set on files and directories in your file system determine the level of access your users get when transferring files into and out of your Amazon EFS file systems.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   posixProfileProperty := &PosixProfileProperty{
//   	Gid: jsii.Number(123),
//   	SecondaryGids: []interface{}{
//   		jsii.Number(123),
//   	},
//   	Uid: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-posixprofile.html
//
type CfnUserPropsMixin_PosixProfileProperty struct {
	// The POSIX group ID used for all EFS operations by this user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-posixprofile.html#cfn-transfer-user-posixprofile-gid
	//
	Gid *float64 `field:"optional" json:"gid" yaml:"gid"`
	// The secondary POSIX group IDs used for all EFS operations by this user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-posixprofile.html#cfn-transfer-user-posixprofile-secondarygids
	//
	SecondaryGids interface{} `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
	// The POSIX user ID used for all EFS operations by this user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-posixprofile.html#cfn-transfer-user-posixprofile-uid
	//
	Uid *float64 `field:"optional" json:"uid" yaml:"uid"`
}

