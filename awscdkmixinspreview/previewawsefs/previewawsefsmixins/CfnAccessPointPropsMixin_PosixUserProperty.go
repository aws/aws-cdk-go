package previewawsefsmixins


// The full POSIX identity, including the user ID, group ID, and any secondary group IDs, on the access point that is used for all file system operations performed by NFS clients using the access point.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   posixUserProperty := &PosixUserProperty{
//   	Gid: jsii.String("gid"),
//   	SecondaryGids: []*string{
//   		jsii.String("secondaryGids"),
//   	},
//   	Uid: jsii.String("uid"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-posixuser.html
//
type CfnAccessPointPropsMixin_PosixUserProperty struct {
	// The POSIX group ID used for all file system operations using this access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-posixuser.html#cfn-efs-accesspoint-posixuser-gid
	//
	Gid *string `field:"optional" json:"gid" yaml:"gid"`
	// Secondary POSIX group IDs used for all file system operations using this access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-posixuser.html#cfn-efs-accesspoint-posixuser-secondarygids
	//
	SecondaryGids *[]*string `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
	// The POSIX user ID used for all file system operations using this access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-posixuser.html#cfn-efs-accesspoint-posixuser-uid
	//
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
}

