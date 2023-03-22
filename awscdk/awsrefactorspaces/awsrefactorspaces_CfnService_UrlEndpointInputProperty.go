package awsrefactorspaces


// The configuration for the URL endpoint type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   urlEndpointInputProperty := &urlEndpointInputProperty{
//   	url: jsii.String("url"),
//
//   	// the properties below are optional
//   	healthUrl: jsii.String("healthUrl"),
//   }
//
type CfnService_UrlEndpointInputProperty struct {
	// The URL to route traffic to.
	//
	// The URL must be an [rfc3986-formatted URL](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc3986) . If the host is a domain name, the name must be resolvable over the public internet. If the scheme is `https` , the top level domain of the host must be listed in the [IANA root zone database](https://docs.aws.amazon.com/https://www.iana.org/domains/root/db) .
	Url *string `field:"required" json:"url" yaml:"url"`
	// The health check URL of the URL endpoint type.
	//
	// If the URL is a public endpoint, the `HealthUrl` must also be a public endpoint. If the URL is a private endpoint inside a virtual private cloud (VPC), the health URL must also be a private endpoint, and the host must be the same as the URL.
	HealthUrl *string `field:"optional" json:"healthUrl" yaml:"healthUrl"`
}

