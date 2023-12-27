package awssagemaker


// Details about the POSIX identity that is used for file system operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customPosixUserConfigProperty := &CustomPosixUserConfigProperty{
//   	Gid: jsii.Number(123),
//   	Uid: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-customposixuserconfig.html
//
type CfnDomain_CustomPosixUserConfigProperty struct {
	// The POSIX group ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-customposixuserconfig.html#cfn-sagemaker-domain-customposixuserconfig-gid
	//
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// The POSIX user ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-customposixuserconfig.html#cfn-sagemaker-domain-customposixuserconfig-uid
	//
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
}

