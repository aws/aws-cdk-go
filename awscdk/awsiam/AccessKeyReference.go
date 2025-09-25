package awsiam


// A reference to a AccessKey resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessKeyReference := &AccessKeyReference{
//   	AccessKeyId: jsii.String("accessKeyId"),
//   }
//
type AccessKeyReference struct {
	// The Id of the AccessKey resource.
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
}

