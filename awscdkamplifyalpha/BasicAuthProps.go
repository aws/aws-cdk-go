// The CDK Construct Library for AWS::Amplify
package awscdkamplifyalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a BasicAuth.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import amplify_alpha "github.com/aws/aws-cdk-go/awscdkamplifyalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//   var secretValue secretValue
//
//   basicAuthProps := &BasicAuthProps{
//   	Username: jsii.String("username"),
//
//   	// the properties below are optional
//   	EncryptionKey: key,
//   	Password: secretValue,
//   }
//
// Experimental.
type BasicAuthProps struct {
	// The username.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
	// The encryption key to use to encrypt the password when it's generated in Secrets Manager.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The password.
	// Experimental.
	Password awscdk.SecretValue `field:"optional" json:"password" yaml:"password"`
}

