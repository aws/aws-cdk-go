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
//   applicationListenerCertificateProps := &applicationListenerCertificateProps{
//   	listener: applicationListener,
//
//   	// the properties below are optional
//   	certificateArns: []*string{
//   		jsii.String("certificateArns"),
//   	},
//   	certificates: []iListenerCertificate{
//   		listenerCertificate,
//   	},
//   }
//
// Experimental.
type ApplicationListenerCertificateProps struct {
	// The listener to attach the rule to.
	// Experimental.
	Listener IApplicationListener `field:"required" json:"listener" yaml:"listener"`
	// ARNs of certificates to attach.
	//
	// Duplicates are not allowed.
	// Deprecated: Use `certificates` instead.
	CertificateArns *[]*string `field:"optional" json:"certificateArns" yaml:"certificateArns"`
	// Certificates to attach.
	//
	// Duplicates are not allowed.
	// Experimental.
	Certificates *[]IListenerCertificate `field:"optional" json:"certificates" yaml:"certificates"`
}

