package awscloudfront


// The SSL method CloudFront will use for your distribution.
//
// Server Name Indication (SNI) - is an extension to the TLS computer networking protocol by which a client indicates
//   which hostname it is attempting to connect to at the start of the handshaking process. This allows a server to present
//   multiple certificates on the same IP address and TCP port number and hence allows multiple secure (HTTPS) websites
// (or any other service over TLS) to be served by the same IP address without requiring all those sites to use the same certificate.
//
// CloudFront can use SNI to host multiple distributions on the same IP - which a large majority of clients will support.
//
// If your clients cannot support SNI however - CloudFront can use dedicated IPs for your distribution - but there is a prorated monthly charge for
// using this feature. By default, we use SNI - but you can optionally enable dedicated IPs (VIP).
//
// See the CloudFront SSL for more details about pricing : https://aws.amazon.com/cloudfront/custom-ssl-domains/
//
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
type SSLMethod string

const (
	SSLMethod_SNI SSLMethod = "SNI"
	SSLMethod_VIP SSLMethod = "VIP"
)

