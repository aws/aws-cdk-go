package awsappmesh


// An object that represents the listener's Secret Discovery Service certificate.
//
// The proxy must be configured with a local SDS provider via a Unix Domain Socket. See App Mesh [TLS documentation](https://docs.aws.amazon.com/app-mesh/latest/userguide/tls.html) for more info.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerTlsSdsCertificateProperty := &listenerTlsSdsCertificateProperty{
//   	secretName: jsii.String("secretName"),
//   }
//
type CfnVirtualNode_ListenerTlsSdsCertificateProperty struct {
	// A reference to an object that represents the name of the secret requested from the Secret Discovery Service provider representing Transport Layer Security (TLS) materials like a certificate or certificate chain.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

