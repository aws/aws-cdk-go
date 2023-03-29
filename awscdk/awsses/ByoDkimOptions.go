package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for BYO DKIM.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secretValue secretValue
//
//   byoDkimOptions := &ByoDkimOptions{
//   	PrivateKey: secretValue,
//   	Selector: jsii.String("selector"),
//
//   	// the properties below are optional
//   	PublicKey: jsii.String("publicKey"),
//   }
//
type ByoDkimOptions struct {
	// The private key that's used to generate a DKIM signature.
	PrivateKey awscdk.SecretValue `field:"required" json:"privateKey" yaml:"privateKey"`
	// A string that's used to identify a public key in the DNS configuration for a domain.
	Selector *string `field:"required" json:"selector" yaml:"selector"`
	// The public key.
	//
	// If specified, a TXT record with the public key is created.
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
}

