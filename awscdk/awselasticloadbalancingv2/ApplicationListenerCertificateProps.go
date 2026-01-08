package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancingv2"
)

// Properties for adding a set of certificates to a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var listenerCertificate ListenerCertificate
//   var listenerRef IListenerRef
//
//   applicationListenerCertificateProps := &ApplicationListenerCertificateProps{
//   	Listener: listenerRef,
//
//   	// the properties below are optional
//   	Certificates: []IListenerCertificate{
//   		listenerCertificate,
//   	},
//   }
//
type ApplicationListenerCertificateProps struct {
	// The listener to attach the rule to.
	Listener interfacesawselasticloadbalancingv2.IListenerRef `field:"required" json:"listener" yaml:"listener"`
	// Certificates to attach.
	//
	// Duplicates are not allowed.
	// Default: - One of 'certificates' and 'certificateArns' is required.
	//
	Certificates *[]IListenerCertificate `field:"optional" json:"certificates" yaml:"certificates"`
}

