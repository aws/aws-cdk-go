package awsiam


// Properties for defining a `CfnAccessKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessKeyProps := &CfnAccessKeyProps{
//   	UserName: jsii.String("userName"),
//
//   	// the properties below are optional
//   	Serial: jsii.Number(123),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-accesskey.html
//
type CfnAccessKeyProps struct {
	// The name of the IAM user that the new key will belong to.
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-accesskey.html#cfn-iam-accesskey-username
	//
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// This value is specific to CloudFormation and can only be *incremented* .
	//
	// Incrementing this value notifies CloudFormation that you want to rotate your access key. When you update your stack, CloudFormation will replace the existing access key with a new key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-accesskey.html#cfn-iam-accesskey-serial
	//
	Serial *float64 `field:"optional" json:"serial" yaml:"serial"`
	// The status of the access key.
	//
	// `Active` means that the key is valid for API calls, while `Inactive` means it is not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-accesskey.html#cfn-iam-accesskey-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

