package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

// Properties to define an application listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var certificate certificate
//
//   applicationListenerProps := &ApplicationListenerProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Certificate: certificate,
//   	Port: jsii.Number(123),
//   	Protocol: awscdk.Aws_elasticloadbalancingv2.ApplicationProtocol_HTTP,
//   	SslPolicy: awscdk.*Aws_elasticloadbalancingv2.SslPolicy_RECOMMENDED,
//   }
//
// Experimental.
type ApplicationListenerProps struct {
	// Name of the listener.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Certificate Manager certificate to associate with the load balancer.
	//
	// Setting this option will set the load balancer protocol to HTTPS.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// The port on which the listener listens for requests.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol for connections from clients to the load balancer.
	//
	// The load balancer port is determined from the protocol (port 80 for
	// HTTP, port 443 for HTTPS).  A domain name and zone must be also be
	// specified if using HTTPS.
	// Experimental.
	Protocol awselasticloadbalancingv2.ApplicationProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The security policy that defines which ciphers and protocols are supported by the ALB Listener.
	// Experimental.
	SslPolicy awselasticloadbalancingv2.SslPolicy `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
}

