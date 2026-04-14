package awss3files


// Properties for CfnFileSystemPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var policy interface{}
//
//   cfnFileSystemPolicyMixinProps := &CfnFileSystemPolicyMixinProps{
//   	FileSystemId: jsii.String("fileSystemId"),
//   	Policy: policy,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-filesystempolicy.html
//
type CfnFileSystemPolicyMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-filesystempolicy.html#cfn-s3files-filesystempolicy-filesystemid
	//
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-filesystempolicy.html#cfn-s3files-filesystempolicy-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
}

