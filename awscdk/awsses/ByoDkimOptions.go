package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for BYO DKIM.
//
// Example:
//   var myHostedZone iPublicHostedZone
//
//
//   ses.NewEmailIdentity(this, jsii.String("Identity"), &EmailIdentityProps{
//   	Identity: ses.Identity_PublicHostedZone(myHostedZone),
//   	DkimIdentity: ses.DkimIdentity_ByoDkim(&ByoDkimOptions{
//   		PrivateKey: awscdk.SecretValue_SecretsManager(jsii.String("dkim-private-key")),
//   		PublicKey: jsii.String("...base64-encoded-public-key..."),
//   		Selector: jsii.String("selector"),
//   	}),
//   })
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

