package awsappmesh


// An object that represents a virtual gateway's listener's Transport Layer Security (TLS) Secret Discovery Service validation context trust.
//
// The proxy must be configured with a local SDS provider via a Unix Domain Socket. See App Mesh [TLS documentation](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html) for more info.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayTlsValidationContextSdsTrustProperty := &virtualGatewayTlsValidationContextSdsTrustProperty{
//   	secretName: jsii.String("secretName"),
//   }
//
type CfnVirtualGateway_VirtualGatewayTlsValidationContextSdsTrustProperty struct {
	// A reference to an object that represents the name of the secret for a virtual gateway's Transport Layer Security (TLS) Secret Discovery Service validation context trust.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

