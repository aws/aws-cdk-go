package interfacesawscertificatemanager


// A reference to a AcmeEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acmeEndpointReference := &AcmeEndpointReference{
//   	AcmeEndpointArn: jsii.String("acmeEndpointArn"),
//   }
//
type AcmeEndpointReference struct {
	// The AcmeEndpointArn of the AcmeEndpoint resource.
	AcmeEndpointArn *string `field:"required" json:"acmeEndpointArn" yaml:"acmeEndpointArn"`
}

