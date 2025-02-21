package awselasticloadbalancingv2


// Properties for adding a set of certificates to a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationListener applicationListener
//   var listenerCertificate listenerCertificate
//
//   applicationListenerCertificateProps := &ApplicationListenerCertificateProps{
//   	Listener: applicationListener,
//
//   	// the properties below are optional
//   	Certificates: []iListenerCertificate{
//   		listenerCertificate,
//   	},
//   }
//
type ApplicationListenerCertificateProps struct {
	// The listener to attach the rule to.
	Listener IApplicationListener `field:"required" json:"listener" yaml:"listener"`
	// Certificates to attach.
	//
	// Duplicates are not allowed.
	// Default: - One of 'certificates' and 'certificateArns' is required.
	//
	Certificates *[]IListenerCertificate `field:"optional" json:"certificates" yaml:"certificates"`
}

