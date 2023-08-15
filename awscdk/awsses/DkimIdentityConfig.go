package awsses


// Configuration for DKIM identity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dkimIdentityConfig := &DkimIdentityConfig{
//   	DomainSigningPrivateKey: jsii.String("domainSigningPrivateKey"),
//   	DomainSigningSelector: jsii.String("domainSigningSelector"),
//   	NextSigningKeyLength: awscdk.Aws_ses.EasyDkimSigningKeyLength_RSA_1024_BIT,
//   }
//
type DkimIdentityConfig struct {
	// A private key that's used to generate a DKIM signature.
	// Default: - use Easy DKIM.
	//
	DomainSigningPrivateKey *string `field:"optional" json:"domainSigningPrivateKey" yaml:"domainSigningPrivateKey"`
	// A string that's used to identify a public key in the DNS configuration for a domain.
	// Default: - use Easy DKIM.
	//
	DomainSigningSelector *string `field:"optional" json:"domainSigningSelector" yaml:"domainSigningSelector"`
	// The key length of the future DKIM key pair to be generated.
	//
	// This can be changed
	// at most once per day.
	// Default: EasyDkimSigningKeyLength.RSA_2048_BIT
	//
	NextSigningKeyLength EasyDkimSigningKeyLength `field:"optional" json:"nextSigningKeyLength" yaml:"nextSigningKeyLength"`
}

