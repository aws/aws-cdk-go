package awss3files


// Properties for defining a `CfnFileSystemPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnFileSystemPolicyProps := &CfnFileSystemPolicyProps{
//   	FileSystemId: jsii.String("fileSystemId"),
//   	Policy: policy,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-filesystempolicy.html
//
type CfnFileSystemPolicyProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-filesystempolicy.html#cfn-s3files-filesystempolicy-filesystemid
	//
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-filesystempolicy.html#cfn-s3files-filesystempolicy-policy
	//
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
}

