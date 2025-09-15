package awssecretsmanager


// A reference to a Secret resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretReference := &SecretReference{
//   	SecretId: jsii.String("secretId"),
//   }
//
type SecretReference struct {
	// The Id of the Secret resource.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
}

