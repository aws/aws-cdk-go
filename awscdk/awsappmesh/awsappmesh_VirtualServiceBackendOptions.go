package awsappmesh


// Represents the properties needed to define a Virtual Service backend.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var mutualTlsCertificate mutualTlsCertificate
//   var subjectAlternativeNames subjectAlternativeNames
//   var tlsValidationTrust tlsValidationTrust
//
//   virtualServiceBackendOptions := &virtualServiceBackendOptions{
//   	tlsClientPolicy: &tlsClientPolicy{
//   		validation: &tlsValidation{
//   			trust: tlsValidationTrust,
//
//   			// the properties below are optional
//   			subjectAlternativeNames: subjectAlternativeNames,
//   		},
//
//   		// the properties below are optional
//   		enforce: jsii.Boolean(false),
//   		mutualTlsCertificate: mutualTlsCertificate,
//   		ports: []*f64{
//   			jsii.Number(123),
//   		},
//   	},
//   }
//
type VirtualServiceBackendOptions struct {
	// TLS properties for  Client policy for the backend.
	TlsClientPolicy *TlsClientPolicy `field:"optional" json:"tlsClientPolicy" yaml:"tlsClientPolicy"`
}

