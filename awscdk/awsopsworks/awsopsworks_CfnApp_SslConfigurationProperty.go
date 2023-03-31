package awsopsworks


// Describes an app's SSL configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sslConfigurationProperty := &sslConfigurationProperty{
//   	certificate: jsii.String("certificate"),
//   	chain: jsii.String("chain"),
//   	privateKey: jsii.String("privateKey"),
//   }
//
type CfnApp_SslConfigurationProperty struct {
	// The contents of the certificate's domain.crt file.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Optional.
	//
	// Can be used to specify an intermediate certificate authority key or client authentication.
	Chain *string `field:"optional" json:"chain" yaml:"chain"`
	// The private key;
	//
	// the contents of the certificate's domain.kex file.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
}

