package awsiot


// An object that specifies the TLS configuration for a domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tlsConfigProperty := &TlsConfigProperty{
//   	SecurityPolicy: jsii.String("securityPolicy"),
//   }
//
type CfnDomainConfiguration_TlsConfigProperty struct {
	// The security policy for a domain configuration.
	//
	// For more information, see [Security policies](https://docs.aws.amazon.com/iot/latest/developerguide/transport-security.html#tls-policy-table) in the *AWS IoT Core developer guide* .
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
}

