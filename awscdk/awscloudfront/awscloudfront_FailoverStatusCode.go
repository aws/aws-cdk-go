package awscloudfront


// HTTP status code to failover to second origin.
//
// Example:
//   // Configuring origin fallback options for the CloudFrontWebDistribution
//   // Configuring origin fallback options for the CloudFrontWebDistribution
//   cloudfront.NewCloudFrontWebDistribution(this, jsii.String("ADistribution"), &cloudFrontWebDistributionProps{
//   	originConfigs: []sourceConfiguration{
//   		&sourceConfiguration{
//   			s3OriginSource: &s3OriginConfig{
//   				s3BucketSource: s3.bucket.fromBucketName(this, jsii.String("aBucket"), jsii.String("myoriginbucket")),
//   				originPath: jsii.String("/"),
//   				originHeaders: map[string]*string{
//   					"myHeader": jsii.String("42"),
//   				},
//   				originShieldRegion: jsii.String("us-west-2"),
//   			},
//   			failoverS3OriginSource: &s3OriginConfig{
//   				s3BucketSource: s3.*bucket.fromBucketName(this, jsii.String("aBucketFallback"), jsii.String("myoriginbucketfallback")),
//   				originPath: jsii.String("/somewhere"),
//   				originHeaders: map[string]*string{
//   					"myHeader2": jsii.String("21"),
//   				},
//   				originShieldRegion: jsii.String("us-east-1"),
//   			},
//   			failoverCriteriaStatusCodes: []failoverStatusCode{
//   				cloudfront.*failoverStatusCode_INTERNAL_SERVER_ERROR,
//   			},
//   			behaviors: []behavior{
//   				&behavior{
//   					isDefaultBehavior: jsii.Boolean(true),
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

