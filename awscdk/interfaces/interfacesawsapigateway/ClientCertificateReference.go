package interfacesawsapigateway


// A reference to a ClientCertificate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientCertificateReference := &ClientCertificateReference{
//   	ClientCertificateId: jsii.String("clientCertificateId"),
//   }
//
type ClientCertificateReference struct {
	// The ClientCertificateId of the ClientCertificate resource.
	ClientCertificateId *string `field:"required" json:"clientCertificateId" yaml:"clientCertificateId"`
}

