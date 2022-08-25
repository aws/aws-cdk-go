package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// A source configuration is a wrapper for CloudFront origins and behaviors.
//
// An origin is what CloudFront will "be in front of" - that is, CloudFront will pull it's assets from an origin.
//
// If you're using s3 as a source - pass the `s3Origin` property, otherwise, pass the `customOriginSource` property.
//
// One or the other must be passed, and it is invalid to pass both in the same SourceConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var function_ function
//   var keyGroup keyGroup
//   var originAccessIdentity originAccessIdentity
//   var version version
//
//   sourceConfiguration := &sourceConfiguration{
//   	behaviors: []behavior{
//   		&behavior{
//   			allowedMethods: awscdk.Aws_cloudfront.cloudFrontAllowedMethods_GET_HEAD,
//   			cachedMethods: awscdk.*Aws_cloudfront.cloudFrontAllowedCachedMethods_GET_HEAD,
//   			compress: jsii.Boolean(false),
//   			defaultTtl: cdk.duration.minutes(jsii.Number(30)),
//   			forwardedValues: &forwardedValuesProperty{
//   				queryString: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				cookies: &cookiesProperty{
//   					forward: jsii.String("forward"),
//
//   					// the properties below are optional
//   					whitelistedNames: []*string{
//   						jsii.String("whitelistedNames"),
//   					},
//   				},
//   				headers: []*string{
//   					jsii.String("headers"),
//   				},
//   				queryStringCacheKeys: []*string{
//   					jsii.String("queryStringCacheKeys"),
//   				},
//   			},
//   			functionAssociations: []functionAssociation{
//   				&functionAssociation{
//   					eventType: awscdk.*Aws_cloudfront.functionEventType_VIEWER_REQUEST,
//   					function: function_,
//   				},
//   			},
//   			isDefaultBehavior: jsii.Boolean(false),
//   			lambdaFunctionAssociations: []lambdaFunctionAssociation{
//   				&lambdaFunctionAssociation{
//   					eventType: awscdk.*Aws_cloudfront.lambdaEdgeEventType_ORIGIN_REQUEST,
//   					lambdaFunction: version,
//
//   					// the properties below are optional
//   					includeBody: jsii.Boolean(false),
//   				},
//   			},
//   			maxTtl: cdk.*duration.minutes(jsii.Number(30)),
//   			minTtl: cdk.*duration.minutes(jsii.Number(30)),
//   			pathPattern: jsii.String("pathPattern"),
//   			trustedKeyGroups: []iKeyGroup{
//   				keyGroup,
//   			},
//   			trustedSigners: []*string{
//   				jsii.String("trustedSigners"),
//   			},
//   			viewerProtocolPolicy: awscdk.*Aws_cloudfront.viewerProtocolPolicy_HTTPS_ONLY,
//   		},
//   	},
//
//   	// the properties below are optional
//   	connectionAttempts: jsii.Number(123),
//   	connectionTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   	customOriginSource: &customOriginConfig{
//   		domainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		allowedOriginSSLVersions: []originSslPolicy{
//   			awscdk.*Aws_cloudfront.*originSslPolicy_SSL_V3,
//   		},
//   		httpPort: jsii.Number(123),
//   		httpsPort: jsii.Number(123),
//   		originHeaders: map[string]*string{
//   			"originHeadersKey": jsii.String("originHeaders"),
//   		},
//   		originKeepaliveTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   		originPath: jsii.String("originPath"),
//   		originProtocolPolicy: awscdk.*Aws_cloudfront.originProtocolPolicy_HTTP_ONLY,
//   		originReadTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   		originShieldRegion: jsii.String("originShieldRegion"),
//   	},
//   	failoverCriteriaStatusCodes: []failoverStatusCode{
//   		awscdk.*Aws_cloudfront.*failoverStatusCode_FORBIDDEN,
//   	},
//   	failoverCustomOriginSource: &customOriginConfig{
//   		domainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		allowedOriginSSLVersions: []*originSslPolicy{
//   			awscdk.*Aws_cloudfront.*originSslPolicy_SSL_V3,
//   		},
//   		httpPort: jsii.Number(123),
//   		httpsPort: jsii.Number(123),
//   		originHeaders: map[string]*string{
//   			"originHeadersKey": jsii.String("originHeaders"),
//   		},
//   		originKeepaliveTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   		originPath: jsii.String("originPath"),
//   		originProtocolPolicy: awscdk.*Aws_cloudfront.*originProtocolPolicy_HTTP_ONLY,
//   		originReadTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   		originShieldRegion: jsii.String("originShieldRegion"),
//   	},
//   	failoverS3OriginSource: &s3OriginConfig{
//   		s3BucketSource: bucket,
//
//   		// the properties below are optional
//   		originAccessIdentity: originAccessIdentity,
//   		originHeaders: map[string]*string{
//   			"originHeadersKey": jsii.String("originHeaders"),
//   		},
//   		originPath: jsii.String("originPath"),
//   		originShieldRegion: jsii.String("originShieldRegion"),
//   	},
//   	originShieldRegion: jsii.String("originShieldRegion"),
//   	s3OriginSource: &s3OriginConfig{
//   		s3BucketSource: bucket,
//
//   		// the properties below are optional
//   		originAccessIdentity: originAccessIdentity,
//   		originHeaders: map[string]*string{
//   			"originHeadersKey": jsii.String("originHeaders"),
//   		},
//   		originPath: jsii.String("originPath"),
//   		originShieldRegion: jsii.String("originShieldRegion"),
//   	},
//   }
//
type SourceConfiguration struct {
	// The behaviors associated with this source.
	//
	// At least one (default) behavior must be included.
	Behaviors *[]*Behavior `field:"required" json:"behaviors" yaml:"behaviors"`
	// The number of times that CloudFront attempts to connect to the origin.
	//
	// You can specify 1, 2, or 3 as the number of attempts.
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// You can specify a number of seconds between 1 and 10 (inclusive).
	ConnectionTimeout awscdk.Duration `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// A custom origin source - for all non-s3 sources.
	CustomOriginSource *CustomOriginConfig `field:"optional" json:"customOriginSource" yaml:"customOriginSource"`
	// HTTP status code to failover to second origin.
	FailoverCriteriaStatusCodes *[]FailoverStatusCode `field:"optional" json:"failoverCriteriaStatusCodes" yaml:"failoverCriteriaStatusCodes"`
	// A custom origin source for failover in case the s3OriginSource returns invalid status code.
	FailoverCustomOriginSource *CustomOriginConfig `field:"optional" json:"failoverCustomOriginSource" yaml:"failoverCustomOriginSource"`
	// An s3 origin source for failover in case the s3OriginSource returns invalid status code.
	FailoverS3OriginSource *S3OriginConfig `field:"optional" json:"failoverS3OriginSource" yaml:"failoverS3OriginSource"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html
	//
	OriginShieldRegion *string `field:"optional" json:"originShieldRegion" yaml:"originShieldRegion"`
	// An s3 origin source - if you're using s3 for your assets.
	S3OriginSource *S3OriginConfig `field:"optional" json:"s3OriginSource" yaml:"s3OriginSource"`
}

