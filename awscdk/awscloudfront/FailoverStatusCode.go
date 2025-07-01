package awscloudfront


// HTTP status code to failover to second origin.
//
// Example:
//   // Configuring origin fallback options for the CloudFrontWebDistribution
//   // Configuring origin fallback options for the CloudFrontWebDistribution
//   cloudfront.NewCloudFrontWebDistribution(this, jsii.String("ADistribution"), &CloudFrontWebDistributionProps{
//   	OriginConfigs: []sourceConfiguration{
//   		&sourceConfiguration{
//   			S3OriginSource: &S3OriginConfig{
//   				S3BucketSource: s3.Bucket_FromBucketName(this, jsii.String("aBucket"), jsii.String("amzn-s3-demo-bucket")),
//   				OriginPath: jsii.String("/"),
//   				OriginHeaders: map[string]*string{
//   					"myHeader": jsii.String("42"),
//   				},
//   				OriginShieldRegion: jsii.String("us-west-2"),
//   			},
//   			FailoverS3OriginSource: &S3OriginConfig{
//   				S3BucketSource: s3.Bucket_*FromBucketName(this, jsii.String("aBucketFallback"), jsii.String("amzn-s3-demo-bucket1")),
//   				OriginPath: jsii.String("/somewhere"),
//   				OriginHeaders: map[string]*string{
//   					"myHeader2": jsii.String("21"),
//   				},
//   				OriginShieldRegion: jsii.String("us-east-1"),
//   			},
//   			FailoverCriteriaStatusCodes: []failoverStatusCode{
//   				cloudfront.*failoverStatusCode_INTERNAL_SERVER_ERROR,
//   			},
//   			Behaviors: []behavior{
//   				&behavior{
//   					IsDefaultBehavior: jsii.Boolean(true),
//   				},
//   			},
//   		},
//   	},
//   })
//
type FailoverStatusCode string

const (
	// Forbidden (403).
	FailoverStatusCode_FORBIDDEN FailoverStatusCode = "FORBIDDEN"
	// Not found (404).
	FailoverStatusCode_NOT_FOUND FailoverStatusCode = "NOT_FOUND"
	// Internal Server Error (500).
	FailoverStatusCode_INTERNAL_SERVER_ERROR FailoverStatusCode = "INTERNAL_SERVER_ERROR"
	// Bad Gateway (502).
	FailoverStatusCode_BAD_GATEWAY FailoverStatusCode = "BAD_GATEWAY"
	// Service Unavailable (503).
	FailoverStatusCode_SERVICE_UNAVAILABLE FailoverStatusCode = "SERVICE_UNAVAILABLE"
	// Gateway Timeout (504).
	FailoverStatusCode_GATEWAY_TIMEOUT FailoverStatusCode = "GATEWAY_TIMEOUT"
)

