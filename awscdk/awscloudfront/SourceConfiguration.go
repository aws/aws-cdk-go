package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// A source configuration is a wrapper for CloudFront origins and behaviors.
//
// An origin is what CloudFront will "be in front of" - that is, CloudFront will pull its assets from an origin.
//
// If you're using s3 as a source - pass the `s3Origin` property, otherwise, pass the `customOriginSource` property.
//
// One or the other must be passed, and it is invalid to pass both in the same SourceConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   sourceConfiguration := &SourceConfiguration{
//   	Behaviors: []behavior{
//   		&behavior{
//   			AllowedMethods: awscdk.Aws_cloudfront.CloudFrontAllowedMethods_GET_HEAD,
//   			CachedMethods: awscdk.*Aws_cloudfront.CloudFrontAllowedCachedMethods_GET_HEAD,
//   			Compress: jsii.Boolean(false),
//   			DefaultTtl: cdk.Duration_Minutes(jsii.Number(30)),
//   			ForwardedValues: &ForwardedValuesProperty{
//   				QueryString: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				Cookies: &CookiesProperty{
//   					Forward: jsii.String("forward"),
//
//   					// the properties below are optional
//   					WhitelistedNames: []*string{
//   						jsii.String("whitelistedNames"),
//   					},
//   				},
//   				Headers: []*string{
//   					jsii.String("headers"),
//   				},
//   				QueryStringCacheKeys: []*string{
//   					jsii.String("queryStringCacheKeys"),
//   				},
//   			},
//   			FunctionAssociations: []functionAssociation{
//   				&functionAssociation{
//   					EventType: awscdk.*Aws_cloudfront.FunctionEventType_VIEWER_REQUEST,
//   					Function: function_,
//   				},
//   			},
//   			IsDefaultBehavior: jsii.Boolean(false),
//   			LambdaFunctionAssociations: []lambdaFunctionAssociation{
//   				&lambdaFunctionAssociation{
//   					EventType: awscdk.*Aws_cloudfront.LambdaEdgeEventType_ORIGIN_REQUEST,
//   					LambdaFunction: version,
//
//   					// the properties below are optional
//   					IncludeBody: jsii.Boolean(false),
//   				},
//   			},
//   			MaxTtl: cdk.Duration_*Minutes(jsii.Number(30)),
//   			MinTtl: cdk.Duration_*Minutes(jsii.Number(30)),
//   			PathPattern: jsii.String("pathPattern"),
//   			TrustedKeyGroups: []iKeyGroup{
//   				keyGroup,
//   			},
//   			TrustedSigners: []*string{
//   				jsii.String("trustedSigners"),
//   			},
//   			ViewerProtocolPolicy: awscdk.*Aws_cloudfront.ViewerProtocolPolicy_HTTPS_ONLY,
//   		},
//   	},
//
//   	// the properties below are optional
//   	ConnectionAttempts: jsii.Number(123),
//   	ConnectionTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	CustomOriginSource: &CustomOriginConfig{
//   		DomainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		AllowedOriginSSLVersions: []originSslPolicy{
//   			awscdk.*Aws_cloudfront.*originSslPolicy_SSL_V3,
//   		},
//   		HttpPort: jsii.Number(123),
//   		HttpsPort: jsii.Number(123),
//   		OriginHeaders: map[string]*string{
//   			"originHeadersKey": jsii.String("originHeaders"),
//   		},
//   		OriginKeepaliveTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   		OriginPath: jsii.String("originPath"),
//   		OriginProtocolPolicy: awscdk.*Aws_cloudfront.OriginProtocolPolicy_HTTP_ONLY,
//   		OriginReadTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   		OriginShieldRegion: jsii.String("originShieldRegion"),
//   	},
//   	FailoverCriteriaStatusCodes: []failoverStatusCode{
//   		awscdk.*Aws_cloudfront.*failoverStatusCode_FORBIDDEN,
//   	},
//   	FailoverCustomOriginSource: &CustomOriginConfig{
//   		DomainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		AllowedOriginSSLVersions: []*originSslPolicy{
//   			awscdk.*Aws_cloudfront.*originSslPolicy_SSL_V3,
//   		},
//   		HttpPort: jsii.Number(123),
//   		HttpsPort: jsii.Number(123),
//   		OriginHeaders: map[string]*string{
//   			"originHeadersKey": jsii.String("originHeaders"),
//   		},
//   		OriginKeepaliveTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   		OriginPath: jsii.String("originPath"),
//   		OriginProtocolPolicy: awscdk.*Aws_cloudfront.OriginProtocolPolicy_HTTP_ONLY,
//   		OriginReadTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   		OriginShieldRegion: jsii.String("originShieldRegion"),
//   	},
//   	FailoverS3OriginSource: &S3OriginConfig{
//   		S3BucketSource: bucket,
//
//   		// the properties below are optional
//   		OriginAccessIdentity: originAccessIdentity,
//   		OriginHeaders: map[string]*string{
//   			"originHeadersKey": jsii.String("originHeaders"),
//   		},
//   		OriginPath: jsii.String("originPath"),
//   		OriginShieldRegion: jsii.String("originShieldRegion"),
//   	},
//   	OriginShieldRegion: jsii.String("originShieldRegion"),
//   	S3OriginSource: &S3OriginConfig{
//   		S3BucketSource: bucket,
//
//   		// the properties below are optional
//   		OriginAccessIdentity: originAccessIdentity,
//   		OriginHeaders: map[string]*string{
//   			"originHeadersKey": jsii.String("originHeaders"),
//   		},
//   		OriginPath: jsii.String("originPath"),
//   		OriginShieldRegion: jsii.String("originShieldRegion"),
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
	// Default: 3.
	//
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// You can specify a number of seconds between 1 and 10 (inclusive).
	// Default: cdk.Duration.seconds(10)
	//
	ConnectionTimeout awscdk.Duration `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// A custom origin source - for all non-s3 sources.
	CustomOriginSource *CustomOriginConfig `field:"optional" json:"customOriginSource" yaml:"customOriginSource"`
	// HTTP status code to failover to second origin.
	// Default: [500, 502, 503, 504].
	//
	FailoverCriteriaStatusCodes *[]FailoverStatusCode `field:"optional" json:"failoverCriteriaStatusCodes" yaml:"failoverCriteriaStatusCodes"`
	// A custom origin source for failover in case the s3OriginSource returns invalid status code.
	// Default: - no failover configuration.
	//
	FailoverCustomOriginSource *CustomOriginConfig `field:"optional" json:"failoverCustomOriginSource" yaml:"failoverCustomOriginSource"`
	// An s3 origin source for failover in case the s3OriginSource returns invalid status code.
	// Default: - no failover configuration.
	//
	FailoverS3OriginSource *S3OriginConfig `field:"optional" json:"failoverS3OriginSource" yaml:"failoverS3OriginSource"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html
	//
	// Default: - origin shield not enabled.
	//
	OriginShieldRegion *string `field:"optional" json:"originShieldRegion" yaml:"originShieldRegion"`
	// An s3 origin source - if you're using s3 for your assets.
	S3OriginSource *S3OriginConfig `field:"optional" json:"s3OriginSource" yaml:"s3OriginSource"`
}

