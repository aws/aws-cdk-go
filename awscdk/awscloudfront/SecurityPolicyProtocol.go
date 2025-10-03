package awscloudfront


// The minimum version of the SSL protocol that you want CloudFront to use for HTTPS connections.
//
// CloudFront serves your objects only to browsers or devices that support at least the SSL version that you specify.
//
// Example:
//   s3BucketSource := s3.NewBucket(this, jsii.String("Bucket"))
//
//   distribution := cloudfront.NewCloudFrontWebDistribution(this, jsii.String("AnAmazingWebsiteProbably"), &CloudFrontWebDistributionProps{
//   	OriginConfigs: []sourceConfiguration{
//   		&sourceConfiguration{
//   			S3OriginSource: &S3OriginConfig{
//   				S3BucketSource: *S3BucketSource,
//   			},
//   			Behaviors: []behavior{
//   				&behavior{
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
type SecurityPolicyProtocol string

const (
	SecurityPolicyProtocol_SSL_V3 SecurityPolicyProtocol = "SSL_V3"
	SecurityPolicyProtocol_TLS_V1 SecurityPolicyProtocol = "TLS_V1"
	SecurityPolicyProtocol_TLS_V1_2016 SecurityPolicyProtocol = "TLS_V1_2016"
	SecurityPolicyProtocol_TLS_V1_1_2016 SecurityPolicyProtocol = "TLS_V1_1_2016"
	SecurityPolicyProtocol_TLS_V1_2_2018 SecurityPolicyProtocol = "TLS_V1_2_2018"
	SecurityPolicyProtocol_TLS_V1_2_2019 SecurityPolicyProtocol = "TLS_V1_2_2019"
	SecurityPolicyProtocol_TLS_V1_2_2021 SecurityPolicyProtocol = "TLS_V1_2_2021"
)

