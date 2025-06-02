package awslightsail


// Properties for defining a `CfnLoadBalancerTlsCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoadBalancerTlsCertificateProps := &CfnLoadBalancerTlsCertificateProps{
//   	CertificateDomainName: jsii.String("certificateDomainName"),
//   	CertificateName: jsii.String("certificateName"),
//   	LoadBalancerName: jsii.String("loadBalancerName"),
//
//   	// the properties below are optional
//   	CertificateAlternativeNames: []*string{
//   		jsii.String("certificateAlternativeNames"),
//   	},
//   	HttpsRedirectionEnabled: jsii.Boolean(false),
//   	IsAttached: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancertlscertificate.html
//
type CfnLoadBalancerTlsCertificateProps struct {
	// The domain name for the SSL/TLS certificate.
	//
	// For example, `example.com` or `www.example.com` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancertlscertificate.html#cfn-lightsail-loadbalancertlscertificate-certificatedomainname
	//
	CertificateDomainName *string `field:"required" json:"certificateDomainName" yaml:"certificateDomainName"`
	// The name of the SSL/TLS certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancertlscertificate.html#cfn-lightsail-loadbalancertlscertificate-certificatename
	//
	CertificateName *string `field:"required" json:"certificateName" yaml:"certificateName"`
	// The name of the load balancer that the SSL/TLS certificate is attached to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancertlscertificate.html#cfn-lightsail-loadbalancertlscertificate-loadbalancername
	//
	LoadBalancerName *string `field:"required" json:"loadBalancerName" yaml:"loadBalancerName"`
	// An array of alternative domain names and subdomain names for your SSL/TLS certificate.
	//
	// In addition to the primary domain name, you can have up to nine alternative domain names. Wildcards (such as `*.example.com` ) are not supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancertlscertificate.html#cfn-lightsail-loadbalancertlscertificate-certificatealternativenames
	//
	CertificateAlternativeNames *[]*string `field:"optional" json:"certificateAlternativeNames" yaml:"certificateAlternativeNames"`
	// A Boolean value indicating whether HTTPS redirection is enabled for the load balancer that the TLS certificate is attached to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancertlscertificate.html#cfn-lightsail-loadbalancertlscertificate-httpsredirectionenabled
	//
	HttpsRedirectionEnabled interface{} `field:"optional" json:"httpsRedirectionEnabled" yaml:"httpsRedirectionEnabled"`
	// A Boolean value indicating whether the SSL/TLS certificate is attached to a Lightsail load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancertlscertificate.html#cfn-lightsail-loadbalancertlscertificate-isattached
	//
	IsAttached interface{} `field:"optional" json:"isAttached" yaml:"isAttached"`
}

