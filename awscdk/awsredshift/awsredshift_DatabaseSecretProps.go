package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Construction properties for a DatabaseSecret.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//
//   databaseSecretProps := &DatabaseSecretProps{
//   	Username: jsii.String("username"),
//
//   	// the properties below are optional
//   	EncryptionKey: key,
//   }
//
// Experimental.
type DatabaseSecretProps struct {
	// The username.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
	// The KMS key to use to encrypt the secret.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

