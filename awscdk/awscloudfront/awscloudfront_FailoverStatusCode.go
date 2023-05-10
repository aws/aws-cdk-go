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
//   				S3BucketSource: s3.Bucket_FromBucketName(this, jsii.String("aBucket"), jsii.String("myoriginbucket")),
//   				OriginPath: jsii.String("/"),
//   				OriginHeaders: map[string]*string{
//   					"myHeader": jsii.String("42"),
//   				},
//   				OriginShieldRegion: jsii.String("us-west-2"),
//   			},
//   			FailoverS3OriginSource: &S3OriginConfig{
//   				S3BucketSource: s3.Bucket_*FromBucketName(this, jsii.String("aBucketFallback"), jsii.String("myoriginbucketfallback")),
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
// Experimental.
type FailoverStatusCode string

const (
	// Forbidden (403).
	// Experimental.
	FailoverStatusCode_FORBIDDEN FailoverStatusCode = "FORBIDDEN"
	// Not found (404).
	// Experimental.
	FailoverStatusCode_NOT_FOUND FailoverStatusCode = "NOT_FOUND"
	// Internal Server Error (500).
	// Experimental.
	FailoverStatusCode_INTERNAL_SERVER_ERROR FailoverStatusCode = "INTERNAL_SERVER_ERROR"
	// Bad Gateway (502).
	// Experimental.
	FailoverStatusCode_BAD_GATEWAY FailoverStatusCode = "BAD_GATEWAY"
	// Service Unavailable (503).
	// Experimental.
	FailoverStatusCode_SERVICE_UNAVAILABLE FailoverStatusCode = "SERVICE_UNAVAILABLE"
	// Gateway Timeout (504).
	// Experimental.
	FailoverStatusCode_GATEWAY_TIMEOUT FailoverStatusCode = "GATEWAY_TIMEOUT"
)

