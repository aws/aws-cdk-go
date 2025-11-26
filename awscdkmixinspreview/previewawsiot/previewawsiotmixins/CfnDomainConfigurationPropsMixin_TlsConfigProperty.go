package previewawsiotmixins


// An object that specifies the TLS configuration for a domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tlsConfigProperty := &TlsConfigProperty{
//   	SecurityPolicy: jsii.String("securityPolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-tlsconfig.html
//
type CfnDomainConfigurationPropsMixin_TlsConfigProperty struct {
	// The security policy for a domain configuration.
	//
	// For more information, see [Security policies](https://docs.aws.amazon.com/iot/latest/developerguide/transport-security.html#tls-policy-table) in the *AWS IoT Core developer guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-tlsconfig.html#cfn-iot-domainconfiguration-tlsconfig-securitypolicy
	//
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
}

