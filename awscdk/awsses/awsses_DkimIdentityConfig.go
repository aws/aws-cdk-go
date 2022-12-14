package awsses


// Configuration for DKIM identity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dkimIdentityConfig := &dkimIdentityConfig{
//   	domainSigningPrivateKey: jsii.String("domainSigningPrivateKey"),
//   	domainSigningSelector: jsii.String("domainSigningSelector"),
//   	nextSigningKeyLength: awscdk.Aws_ses.easyDkimSigningKeyLength_RSA_1024_BIT,
//   }
//
type DkimIdentityConfig struct {
	// A private key that's used to generate a DKIM signature.
	DomainSigningPrivateKey *string `field:"optional" json:"domainSigningPrivateKey" yaml:"domainSigningPrivateKey"`
	// A string that's used to identify a public key in the DNS configuration for a domain.
	DomainSigningSelector *string `field:"optional" json:"domainSigningSelector" yaml:"domainSigningSelector"`
	// The key length of the future DKIM key pair to be generated.
	//
	// This can be changed
	// at most once per day.
	NextSigningKeyLength EasyDkimSigningKeyLength `field:"optional" json:"nextSigningKeyLength" yaml:"nextSigningKeyLength"`
}

