package awss3files


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   posixUserProperty := &PosixUserProperty{
//   	Gid: jsii.String("gid"),
//   	Uid: jsii.String("uid"),
//
//   	// the properties below are optional
//   	SecondaryGids: []*string{
//   		jsii.String("secondaryGids"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-accesspoint-posixuser.html
//
type CfnAccessPoint_PosixUserProperty struct {
	// The POSIX group ID used for all file system operations using this access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-accesspoint-posixuser.html#cfn-s3files-accesspoint-posixuser-gid
	//
	Gid *string `field:"required" json:"gid" yaml:"gid"`
	// The POSIX user ID used for all file system operations using this access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-accesspoint-posixuser.html#cfn-s3files-accesspoint-posixuser-uid
	//
	Uid *string `field:"required" json:"uid" yaml:"uid"`
	// Secondary POSIX group IDs used for all file system operations using this access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-accesspoint-posixuser.html#cfn-s3files-accesspoint-posixuser-secondarygids
	//
	SecondaryGids *[]*string `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

