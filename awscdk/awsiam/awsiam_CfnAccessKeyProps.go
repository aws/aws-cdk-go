package awsiam


// Properties for defining a `CfnAccessKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessKeyProps := &cfnAccessKeyProps{
//   	userName: jsii.String("userName"),
//
//   	// the properties below are optional
//   	serial: jsii.Number(123),
//   	status: jsii.String("status"),
//   }
//
type CfnAccessKeyProps struct {
	// The name of the IAM user that the new key will belong to.
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// This value is specific to CloudFormation and can only be *incremented* .
	//
	// Incrementing this value notifies CloudFormation that you want to rotate your access key. When you update your stack, CloudFormation will replace the existing access key with a new key.
	Serial *float64 `field:"optional" json:"serial" yaml:"serial"`
	// The status of the access key.
	//
	// `Active` means that the key is valid for API calls, while `Inactive` means it is not.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

