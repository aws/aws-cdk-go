package awscloudfront


// Example:
//   s3BucketSource := s3.NewBucket(this, jsii.String("Bucket"))
//
//   distribution := cloudfront.NewCloudFrontWebDistribution(this, jsii.String("AnAmazingWebsiteProbably"), &CloudFrontWebDistributionProps{
//   	OriginConfigs: []SourceConfiguration{
//   		&SourceConfiguration{
//   			S3OriginSource: &S3OriginConfig{
//   				S3BucketSource: *S3BucketSource,
//   			},
//   			Behaviors: []Behavior{
//   				&Behavior{
//   					IsDefaultBehavior: jsii.Boolean(true),
//   				},
//   			},
//   		},
//   	},
//   	ViewerCertificate: cloudfront.ViewerCertificate_FromIamCertificate(jsii.String("certificateId"), &ViewerCertificateOptions{
//   		Aliases: []*string{
//   			jsii.String("example.com"),
//   		},
//   		SecurityPolicy: cloudfront.SecurityPolicyProtocol_SSL_V3,
//   		 // default
//   		SslMethod: cloudfront.SSLMethod_SNI,
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
	// Default: - SSLv3 if sslMethod VIP, TLSv1 if sslMethod SNI.
	//
	SecurityPolicy SecurityPolicyProtocol `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// How CloudFront should serve HTTPS requests.
	//
	// See the notes on SSLMethod if you wish to use other SSL termination types.
	// See: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ViewerCertificate.html
	//
	// Default: SSLMethod.SNI
	//
	SslMethod SSLMethod `field:"optional" json:"sslMethod" yaml:"sslMethod"`
}

