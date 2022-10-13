package awscloudfront


// Example:
//   s3BucketSource := s3.NewBucket(this, jsii.String("Bucket"))
//
//   distribution := cloudfront.NewCloudFrontWebDistribution(this, jsii.String("AnAmazingWebsiteProbably"), &cloudFrontWebDistributionProps{
//   	originConfigs: []sourceConfiguration{
//   		&sourceConfiguration{
//   			s3OriginSource: &s3OriginConfig{
//   				s3BucketSource: s3BucketSource,
//   			},
//   			behaviors: []behavior{
//   				&behavior{
//   					isDefaultBehavior: jsii.Boolean(true),
//   				},
//   			},
//   		},
//   	},
//   	viewerCertificate: cloudfront.viewerCertificate.fromIamCertificate(jsii.String("certificateId"), &viewerCertificateOptions{
//   		aliases: []*string{
//   			jsii.String("example.com"),
//   		},
//   		securityPolicy: cloudfront.securityPolicyProtocol_SSL_V3,
//   		 // default
//   		sslMethod: cloudfront.sSLMethod_SNI,
//   	}),
//   })
//
type ViewerCertificateOptions struct {
	// Domain names on the certificate (both main domain name and Subject Alternative names).
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// The minimum version of the SSL protocol that you want CloudFront to use for HTTPS connections.
	//
	// CloudFront serves your objects only to browsers or devices that support at
	// least the SSL version that you specify.
	SecurityPolicy SecurityPolicyProtocol `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// How CloudFront should serve HTTPS requests.
	//
	// See the notes on SSLMethod if you wish to use other SSL termination types.
	// See: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ViewerCertificate.html
	//
	SslMethod SSLMethod `field:"optional" json:"sslMethod" yaml:"sslMethod"`
}

