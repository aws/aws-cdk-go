package awscdkredshiftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Construction properties for a DatabaseSecret.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import redshift_alpha "github.com/aws/aws-cdk-go/awscdkredshiftalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//
//   databaseSecretProps := &DatabaseSecretProps{
//   	Username: jsii.String("username"),
//
//   	// the properties below are optional
//   	EncryptionKey: key,
//   	ExcludeCharacters: jsii.String("excludeCharacters"),
//   }
//
// Experimental.
type DatabaseSecretProps struct {
	// The username.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
	// The KMS key to use to encrypt the secret.
	// Default: default master key.
	//
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Characters to not include in the generated password.
	// Default: '"@/\\\ \''.
	//
	// Experimental.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
}

