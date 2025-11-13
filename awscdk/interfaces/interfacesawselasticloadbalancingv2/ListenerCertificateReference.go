package interfacesawselasticloadbalancingv2


// A reference to a ListenerCertificate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerCertificateReference := &ListenerCertificateReference{
//   	ListenerCertificateId: jsii.String("listenerCertificateId"),
//   }
//
type ListenerCertificateReference struct {
	// The Id of the ListenerCertificate resource.
	ListenerCertificateId *string `field:"required" json:"listenerCertificateId" yaml:"listenerCertificateId"`
}

