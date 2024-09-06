package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration for a set of HTTP response headers that are used for cross-origin resource sharing (CORS).
//
// CloudFront adds these headers to HTTP responses that it sends for CORS requests that match a cache behavior
// associated with this response headers policy.
//
// Example:
//   // Using an existing managed response headers policy
//   var bucketOrigin s3Origin
//
//   cloudfront.NewDistribution(this, jsii.String("myDistManagedPolicy"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: bucketOrigin,
//   		ResponseHeadersPolicy: cloudfront.ResponseHeadersPolicy_CORS_ALLOW_ALL_ORIGINS(),
//   	},
//   })
//
//   // Creating a custom response headers policy -- all parameters optional
//   myResponseHeadersPolicy := cloudfront.NewResponseHeadersPolicy(this, jsii.String("ResponseHeadersPolicy"), &ResponseHeadersPolicyProps{
//   	ResponseHeadersPolicyName: jsii.String("MyPolicy"),
//   	Comment: jsii.String("A default policy"),
//   	CorsBehavior: &ResponseHeadersCorsBehavior{
//   		AccessControlAllowCredentials: jsii.Boolean(false),
//   		AccessControlAllowHeaders: []*string{
//   			jsii.String("X-Custom-Header-1"),
//   			jsii.String("X-Custom-Header-2"),
//   		},
//   		AccessControlAllowMethods: []*string{
//   			jsii.String("GET"),
//   			jsii.String("POST"),
//   		},
//   		AccessControlAllowOrigins: []*string{
//   			jsii.String("*"),
//   		},
//   		AccessControlExposeHeaders: []*string{
//   			jsii.String("X-Custom-Header-1"),
//   			jsii.String("X-Custom-Header-2"),
//   		},
//   		AccessControlMaxAge: awscdk.Duration_Seconds(jsii.Number(600)),
//   		OriginOverride: jsii.Boolean(true),
//   	},
//   	CustomHeadersBehavior: &ResponseCustomHeadersBehavior{
//   		CustomHeaders: []responseCustomHeader{
//   			&responseCustomHeader{
//   				Header: jsii.String("X-Amz-Date"),
//   				Value: jsii.String("some-value"),
//   				Override: jsii.Boolean(true),
//   			},
//   			&responseCustomHeader{
//   				Header: jsii.String("X-Amz-Security-Token"),
//   				Value: jsii.String("some-value"),
//   				Override: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	SecurityHeadersBehavior: &ResponseSecurityHeadersBehavior{
//   		ContentSecurityPolicy: &ResponseHeadersContentSecurityPolicy{
//   			ContentSecurityPolicy: jsii.String("default-src https:;"),
//   			Override: jsii.Boolean(true),
//   		},
//   		ContentTypeOptions: &ResponseHeadersContentTypeOptions{
//   			Override: jsii.Boolean(true),
//   		},
//   		FrameOptions: &ResponseHeadersFrameOptions{
//   			FrameOption: cloudfront.HeadersFrameOption_DENY,
//   			Override: jsii.Boolean(true),
//   		},
//   		ReferrerPolicy: &ResponseHeadersReferrerPolicy{
//   			ReferrerPolicy: cloudfront.HeadersReferrerPolicy_NO_REFERRER,
//   			Override: jsii.Boolean(true),
//   		},
//   		StrictTransportSecurity: &ResponseHeadersStrictTransportSecurity{
//   			AccessControlMaxAge: awscdk.Duration_*Seconds(jsii.Number(600)),
//   			IncludeSubdomains: jsii.Boolean(true),
//   			Override: jsii.Boolean(true),
//   		},
//   		XssProtection: &ResponseHeadersXSSProtection{
//   			Protection: jsii.Boolean(true),
//   			ModeBlock: jsii.Boolean(false),
//   			ReportUri: jsii.String("https://example.com/csp-report"),
//   			Override: jsii.Boolean(true),
//   		},
//   	},
//   	RemoveHeaders: []*string{
//   		jsii.String("Server"),
//   	},
//   	ServerTimingSamplingRate: jsii.Number(50),
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: bucketOrigin,
//   		ResponseHeadersPolicy: myResponseHeadersPolicy,
//   	},
//   })
//
type ResponseHeadersCorsBehavior struct {
	// A Boolean that CloudFront uses as the value for the Access-Control-Allow-Credentials HTTP response header.
	AccessControlAllowCredentials *bool `field:"required" json:"accessControlAllowCredentials" yaml:"accessControlAllowCredentials"`
	// A list of HTTP header names that CloudFront includes as values for the Access-Control-Allow-Headers HTTP response header.
	//
	// You can specify `['*']` to allow all headers.
	AccessControlAllowHeaders *[]*string `field:"required" json:"accessControlAllowHeaders" yaml:"accessControlAllowHeaders"`
	// A list of HTTP methods that CloudFront includes as values for the Access-Control-Allow-Methods HTTP response header.
	AccessControlAllowMethods *[]*string `field:"required" json:"accessControlAllowMethods" yaml:"accessControlAllowMethods"`
	// A list of origins (domain names) that CloudFront can use as the value for the Access-Control-Allow-Origin HTTP response header.
	//
	// You can specify `['*']` to allow all origins.
	AccessControlAllowOrigins *[]*string `field:"required" json:"accessControlAllowOrigins" yaml:"accessControlAllowOrigins"`
	// A Boolean that determines whether CloudFront overrides HTTP response headers received from the origin with the ones specified in this response headers policy.
	OriginOverride *bool `field:"required" json:"originOverride" yaml:"originOverride"`
	// A list of HTTP headers that CloudFront includes as values for the Access-Control-Expose-Headers HTTP response header.
	//
	// You can specify `['*']` to expose all headers.
	// Default: - no headers exposed.
	//
	AccessControlExposeHeaders *[]*string `field:"optional" json:"accessControlExposeHeaders" yaml:"accessControlExposeHeaders"`
	// A number that CloudFront uses as the value for the Access-Control-Max-Age HTTP response header.
	// Default: - no max age.
	//
	AccessControlMaxAge awscdk.Duration `field:"optional" json:"accessControlMaxAge" yaml:"accessControlMaxAge"`
}

