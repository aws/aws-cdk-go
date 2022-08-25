package awssecretsmanager


// Options for a SecretRotationApplication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretRotationApplicationOptions := &secretRotationApplicationOptions{
//   	isMultiUser: jsii.Boolean(false),
//   }
//
type SecretRotationApplicationOptions struct {
	// Whether the rotation application uses the mutli user scheme.
	IsMultiUser *bool `field:"optional" json:"isMultiUser" yaml:"isMultiUser"`
}

