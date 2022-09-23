package awscloudfront


// Example:
//   var sourceBucket bucket
//
//   viewerCertificate := cloudfront.viewerCertificate.fromIamCertificate(jsii.String("MYIAMROLEIDENTIFIER"), &viewerCertificateOptions{
//   	aliases: []*string{
//   		jsii.String("MYALIAS"),
//   	},
//   })
//
//   cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyCfWebDistribution"), &cloudFrontWebDistributionProps{
//   	originConfigs: []sourceConfiguration{
//   		&sourceConfiguration{
//   			s3OriginSource: &s3OriginConfig{
//   				s3BucketSource: sourceBucket,
//   			},
//   			behaviors: []behavior{
//   				&behavior{
//   					isDefaultBehavior: jsii.Boolean(true),
//   				},
//   			},
//   		},
//   	},
//   	viewerCertificate: viewerCertificate,
//   })
//
type CloudFrontWebDistributionProps struct {
	// The origin configurations for this distribution.
	//
	// Behaviors are a part of the origin.
	OriginConfigs *[]*SourceConfiguration `field:"required" json:"originConfigs" yaml:"originConfigs"`
	// A comment for this distribution in the CloudFront console.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The default object to serve.
	DefaultRootObject *string `field:"optional" json:"defaultRootObject" yaml:"defaultRootObject"`
	// Enable or disable the distribution.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// If your distribution should have IPv6 enabled.
	EnableIpV6 *bool `field:"optional" json:"enableIpV6" yaml:"enableIpV6"`
	// How CloudFront should handle requests that are not successful (eg PageNotFound).
	//
	// By default, CloudFront does not replace HTTP status codes in the 4xx and 5xx range
	// with custom error messages. CloudFront does not cache HTTP status codes.
	ErrorConfigurations *[]*CfnDistribution_CustomErrorResponseProperty `field:"optional" json:"errorConfigurations" yaml:"errorConfigurations"`
	// Controls the countries in which your content is distributed.
	GeoRestriction GeoRestriction `field:"optional" json:"geoRestriction" yaml:"geoRestriction"`
	// The max supported HTTP Versions.
	HttpVersion HttpVersion `field:"optional" json:"httpVersion" yaml:"httpVersion"`
	// Optional - if we should enable logging.
	//
	// You can pass an empty object ({}) to have us auto create a bucket for logging.
	// Omission of this property indicates no logging is to be enabled.
	LoggingConfig *LoggingConfiguration `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// The price class for the distribution (this impacts how many locations CloudFront uses for your distribution, and billing).
	PriceClass PriceClass `field:"optional" json:"priceClass" yaml:"priceClass"`
	// Specifies whether you want viewers to use HTTP or HTTPS to request your objects, whether you're using an alternate domain name with HTTPS, and if so, if you're using AWS Certificate Manager (ACM) or a third-party certificate authority.
	// See: https://aws.amazon.com/premiumsupport/knowledge-center/custom-ssl-certificate-cloudfront/
	//
	ViewerCertificate ViewerCertificate `field:"optional" json:"viewerCertificate" yaml:"viewerCertificate"`
	// The default viewer policy for incoming clients.
	ViewerProtocolPolicy ViewerProtocolPolicy `field:"optional" json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
	// Unique identifier that specifies the AWS WAF web ACL to associate with this CloudFront distribution.
	//
	// To specify a web ACL created using the latest version of AWS WAF, use the ACL ARN, for example
	// `arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/473e64fd-f30b-4765-81a0-62ad96dd167a`.
	//
	// To specify a web ACL created using AWS WAF Classic, use the ACL ID, for example `473e64fd-f30b-4765-81a0-62ad96dd167a`.
	// See: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CreateDistribution.html#API_CreateDistribution_RequestParameters.
	//
	WebACLId *string `field:"optional" json:"webACLId" yaml:"webACLId"`
}

