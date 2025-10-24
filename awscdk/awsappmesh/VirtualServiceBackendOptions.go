package awsappmesh


// Represents the properties needed to define a Virtual Service backend.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var mutualTlsCertificate MutualTlsCertificate
//   var subjectAlternativeNames SubjectAlternativeNames
//   var tlsValidationTrust TlsValidationTrust
//
//   virtualServiceBackendOptions := &VirtualServiceBackendOptions{
//   	TlsClientPolicy: &TlsClientPolicy{
//   		Validation: &TlsValidation{
//   			Trust: tlsValidationTrust,
//
//   			// the properties below are optional
//   			SubjectAlternativeNames: subjectAlternativeNames,
//   		},
//
//   		// the properties below are optional
//   		Enforce: jsii.Boolean(false),
//   		MutualTlsCertificate: mutualTlsCertificate,
//   		Ports: []*f64{
//   			jsii.Number(123),
//   		},
//   	},
//   }
//
type VirtualServiceBackendOptions struct {
	// TLS properties for  Client policy for the backend.
	// Default: - none.
	//
	TlsClientPolicy *TlsClientPolicy `field:"optional" json:"tlsClientPolicy" yaml:"tlsClientPolicy"`
}

