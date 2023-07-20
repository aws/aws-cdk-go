package awsopsworks


// Describes an app's SSL configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sslConfigurationProperty := &SslConfigurationProperty{
//   	Certificate: jsii.String("certificate"),
//   	Chain: jsii.String("chain"),
//   	PrivateKey: jsii.String("privateKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html
//
type CfnApp_SslConfigurationProperty struct {
	// The contents of the certificate's domain.crt file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html#cfn-opsworks-app-sslconfiguration-certificate
	//
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Optional.
	//
	// Can be used to specify an intermediate certificate authority key or client authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html#cfn-opsworks-app-sslconfiguration-chain
	//
	Chain *string `field:"optional" json:"chain" yaml:"chain"`
	// The private key;
	//
	// the contents of the certificate's domain.kex file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html#cfn-opsworks-app-sslconfiguration-privatekey
	//
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
}

