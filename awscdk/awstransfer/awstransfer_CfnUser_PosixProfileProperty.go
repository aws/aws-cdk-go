package awstransfer


// The full POSIX identity, including user ID ( `Uid` ), group ID ( `Gid` ), and any secondary groups IDs ( `SecondaryGids` ), that controls your users' access to your Amazon EFS file systems.
//
// The POSIX permissions that are set on files and directories in your file system determine the level of access your users get when transferring files into and out of your Amazon EFS file systems.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   posixProfileProperty := &posixProfileProperty{
//   	gid: jsii.Number(123),
//   	uid: jsii.Number(123),
//
//   	// the properties below are optional
//   	secondaryGids: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
type CfnUser_PosixProfileProperty struct {
	// The POSIX group ID used for all EFS operations by this user.
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// The POSIX user ID used for all EFS operations by this user.
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
	// The secondary POSIX group IDs used for all EFS operations by this user.
	SecondaryGids interface{} `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

