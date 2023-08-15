package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for a Distribution.
//
// Example:
//   var s3Bucket bucket
//   // Add a cloudfront Function to a Distribution
//   cfFunction := cloudfront.NewFunction(this, jsii.String("Function"), &FunctionProps{
//   	Code: cloudfront.FunctionCode_FromInline(jsii.String("function handler(event) { return event.request }")),
//   })
//   cloudfront.NewDistribution(this, jsii.String("distro"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewS3Origin(s3Bucket),
//   		FunctionAssociations: []functionAssociation{
//   			&functionAssociation{
//   				Function: cfFunction,
//   				EventType: cloudfront.FunctionEventType_VIEWER_REQUEST,
//   			},
//   		},
//   	},
//   })
//
type DistributionProps struct {
	// The default behavior for the distribution.
	DefaultBehavior *BehaviorOptions `field:"required" json:"defaultBehavior" yaml:"defaultBehavior"`
	// Additional behaviors for the distribution, mapped by the pathPattern that specifies which requests to apply the behavior to.
	// Default: - no additional behaviors are added.
	//
	AdditionalBehaviors *map[string]*BehaviorOptions `field:"optional" json:"additionalBehaviors" yaml:"additionalBehaviors"`
	// A certificate to associate with the distribution.
	//
	// The certificate must be located in N. Virginia (us-east-1).
	// Default: - the CloudFront wildcard certificate (*.cloudfront.net) will be used.
	//
	Certificate awscertificatemanager.ICertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Any comments you want to include about the distribution.
	// Default: - no comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The object that you want CloudFront to request from your origin (for example, index.html) when a viewer requests the root URL for your distribution. If no default object is set, the request goes to the origin's root (e.g., example.com/).
	// Default: - no default root object.
	//
	DefaultRootObject *string `field:"optional" json:"defaultRootObject" yaml:"defaultRootObject"`
	// Alternative domain names for this distribution.
	//
	// If you want to use your own domain name, such as www.example.com, instead of the cloudfront.net domain name,
	// you can add an alternate domain name to your distribution. If you attach a certificate to the distribution,
	// you must add (at least one of) the domain names of the certificate to this list.
	// Default: - The distribution will only support the default generated name (e.g., d111111abcdef8.cloudfront.net)
	//
	DomainNames *[]*string `field:"optional" json:"domainNames" yaml:"domainNames"`
	// Enable or disable the distribution.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Whether CloudFront will respond to IPv6 DNS requests with an IPv6 address.
	//
	// If you specify false, CloudFront responds to IPv6 DNS requests with the DNS response code NOERROR and with no IP addresses.
	// This allows viewers to submit a second request, for an IPv4 address for your distribution.
	// Default: true.
	//
	EnableIpv6 *bool `field:"optional" json:"enableIpv6" yaml:"enableIpv6"`
	// Enable access logging for the distribution.
	// Default: - false, unless `logBucket` is specified.
	//
	EnableLogging *bool `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// How CloudFront should handle requests that are not successful (e.g., PageNotFound).
	// Default: - No custom error responses.
	//
	ErrorResponses *[]*ErrorResponse `field:"optional" json:"errorResponses" yaml:"errorResponses"`
	// Controls the countries in which your content is distributed.
	// Default: - No geographic restrictions.
	//
	GeoRestriction GeoRestriction `field:"optional" json:"geoRestriction" yaml:"geoRestriction"`
	// Specify the maximum HTTP version that you want viewers to use to communicate with CloudFront.
	//
	// For viewers and CloudFront to use HTTP/2, viewers must support TLS 1.2 or later, and must support server name identification (SNI).
	// Default: HttpVersion.HTTP2
	//
	HttpVersion HttpVersion `field:"optional" json:"httpVersion" yaml:"httpVersion"`
	// The Amazon S3 bucket to store the access logs in.
	//
	// Make sure to set `objectOwnership` to `s3.ObjectOwnership.OBJECT_WRITER` in your custom bucket.
	// Default: - A bucket is created if `enableLogging` is true.
	//
	LogBucket awss3.IBucket `field:"optional" json:"logBucket" yaml:"logBucket"`
	// An optional string that you want CloudFront to prefix to the access log filenames for this distribution.
	// Default: - no prefix.
	//
	LogFilePrefix *string `field:"optional" json:"logFilePrefix" yaml:"logFilePrefix"`
	// Specifies whether you want CloudFront to include cookies in access logs.
	// Default: false.
	//
	LogIncludesCookies *bool `field:"optional" json:"logIncludesCookies" yaml:"logIncludesCookies"`
	// The minimum version of the SSL protocol that you want CloudFront to use for HTTPS connections.
	//
	// CloudFront serves your objects only to browsers or devices that support at
	// least the SSL version that you specify.
	// Default: - SecurityPolicyProtocol.TLS_V1_2_2021 if the '@aws-cdk/aws-cloudfront:defaultSecurityPolicyTLSv1.2_2021' feature flag is set; otherwise, SecurityPolicyProtocol.TLS_V1_2_2019.
	//
	MinimumProtocolVersion SecurityPolicyProtocol `field:"optional" json:"minimumProtocolVersion" yaml:"minimumProtocolVersion"`
	// The price class that corresponds with the maximum price that you want to pay for CloudFront service.
	//
	// If you specify PriceClass_All, CloudFront responds to requests for your objects from all CloudFront edge locations.
	// If you specify a price class other than PriceClass_All, CloudFront serves your objects from the CloudFront edge location
	// that has the lowest latency among the edge locations in your price class.
	// Default: PriceClass.PRICE_CLASS_ALL
	//
	PriceClass PriceClass `field:"optional" json:"priceClass" yaml:"priceClass"`
	// The SSL method CloudFront will use for your distribution.
	//
	// Server Name Indication (SNI) - is an extension to the TLS computer networking protocol by which a client indicates
	// which hostname it is attempting to connect to at the start of the handshaking process. This allows a server to present
	// multiple certificates on the same IP address and TCP port number and hence allows multiple secure (HTTPS) websites
	// (or any other service over TLS) to be served by the same IP address without requiring all those sites to use the same certificate.
	//
	// CloudFront can use SNI to host multiple distributions on the same IP - which a large majority of clients will support.
	//
	// If your clients cannot support SNI however - CloudFront can use dedicated IPs for your distribution - but there is a prorated monthly charge for
	// using this feature. By default, we use SNI - but you can optionally enable dedicated IPs (VIP).
	//
	// See the CloudFront SSL for more details about pricing : https://aws.amazon.com/cloudfront/custom-ssl-domains/
	// Default: SSLMethod.SNI
	//
	SslSupportMethod SSLMethod `field:"optional" json:"sslSupportMethod" yaml:"sslSupportMethod"`
	// Unique identifier that specifies the AWS WAF web ACL to associate with this CloudFront distribution.
	//
	// To specify a web ACL created using the latest version of AWS WAF, use the ACL ARN, for example
	// `arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/473e64fd-f30b-4765-81a0-62ad96dd167a`.
	// To specify a web ACL created using AWS WAF Classic, use the ACL ID, for example `473e64fd-f30b-4765-81a0-62ad96dd167a`.
	// See: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CreateDistribution.html#API_CreateDistribution_RequestParameters.
	//
	// Default: - No AWS Web Application Firewall web access control list (web ACL).
	//
	WebAclId *string `field:"optional" json:"webAclId" yaml:"webAclId"`
}

