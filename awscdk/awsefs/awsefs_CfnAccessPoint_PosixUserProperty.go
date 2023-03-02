package awsefs


// The full POSIX identity, including the user ID, group ID, and any secondary group IDs, on the access point that is used for all file system operations performed by NFS clients using the access point.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   posixUserProperty := &posixUserProperty{
//   	gid: jsii.String("gid"),
//   	uid: jsii.String("uid"),
//
//   	// the properties below are optional
//   	secondaryGids: []*string{
//   		jsii.String("secondaryGids"),
//   	},
//   }
//
type CfnAccessPoint_PosixUserProperty struct {
	// The POSIX group ID used for all file system operations using this access point.
	Gid *string `field:"required" json:"gid" yaml:"gid"`
	// The POSIX user ID used for all file system operations using this access point.
	Uid *string `field:"required" json:"uid" yaml:"uid"`
	// Secondary POSIX group IDs used for all file system operations using this access point.
	SecondaryGids *[]*string `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

