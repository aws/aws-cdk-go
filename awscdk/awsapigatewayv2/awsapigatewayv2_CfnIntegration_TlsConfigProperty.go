package awsapigatewayv2


// The `TlsConfig` property specifies the TLS configuration for a private integration.
//
// If you specify a TLS configuration, private integration traffic uses the HTTPS protocol. Supported only for HTTP APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tlsConfigProperty := &tlsConfigProperty{
//   	serverNameToVerify: jsii.String("serverNameToVerify"),
//   }
//
type CfnIntegration_TlsConfigProperty struct {
	// If you specify a server name, API Gateway uses it to verify the hostname on the integration's certificate.
	//
	// The server name is also included in the TLS handshake to support Server Name Indication (SNI) or virtual hosting.
	ServerNameToVerify *string `field:"optional" json:"serverNameToVerify" yaml:"serverNameToVerify"`
}

