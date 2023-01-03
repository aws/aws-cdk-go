package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   presignedUrlConfigProperty := &presignedUrlConfigProperty{
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	expiresInSec: jsii.Number(123),
//   }
//
type CfnJobTemplate_PresignedUrlConfigProperty struct {
	// `CfnJobTemplate.PresignedUrlConfigProperty.RoleArn`.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// `CfnJobTemplate.PresignedUrlConfigProperty.ExpiresInSec`.
	ExpiresInSec *float64 `field:"optional" json:"expiresInSec" yaml:"expiresInSec"`
}

