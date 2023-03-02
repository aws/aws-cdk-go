package awslightsail


// Properties for defining a `CfnLoadBalancerTlsCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoadBalancerTlsCertificateProps := &cfnLoadBalancerTlsCertificateProps{
//   	certificateDomainName: jsii.String("certificateDomainName"),
//   	certificateName: jsii.String("certificateName"),
//   	loadBalancerName: jsii.String("loadBalancerName"),
//
//   	// the properties below are optional
//   	certificateAlternativeNames: []*string{
//   		jsii.String("certificateAlternativeNames"),
//   	},
//   	httpsRedirectionEnabled: jsii.Boolean(false),
//   	isAttached: jsii.Boolean(false),
//   }
//
type CfnLoadBalancerTlsCertificateProps struct {
	// The domain name for the SSL/TLS certificate.
	//
	// For example, `example.com` or `www.example.com` .
	CertificateDomainName *string `field:"required" json:"certificateDomainName" yaml:"certificateDomainName"`
	// The name of the SSL/TLS certificate.
	CertificateName *string `field:"required" json:"certificateName" yaml:"certificateName"`
	// The name of the load balancer that the SSL/TLS certificate is attached to.
	LoadBalancerName *string `field:"required" json:"loadBalancerName" yaml:"loadBalancerName"`
	// An array of alternative domain names and subdomain names for your SSL/TLS certificate.
	//
	// In addition to the primary domain name, you can have up to nine alternative domain names. Wildcards (such as `*.example.com` ) are not supported.
	CertificateAlternativeNames *[]*string `field:"optional" json:"certificateAlternativeNames" yaml:"certificateAlternativeNames"`
	// A Boolean value indicating whether HTTPS redirection is enabled for the load balancer that the TLS certificate is attached to.
	HttpsRedirectionEnabled interface{} `field:"optional" json:"httpsRedirectionEnabled" yaml:"httpsRedirectionEnabled"`
	// A Boolean value indicating whether the SSL/TLS certificate is attached to a Lightsail load balancer.
	IsAttached interface{} `field:"optional" json:"isAttached" yaml:"isAttached"`
}

