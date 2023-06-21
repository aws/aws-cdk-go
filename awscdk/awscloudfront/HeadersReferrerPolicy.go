package awscloudfront


// Enum representing possible values of the Referrer-Policy HTTP response header.
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
//   			ModeBlock: jsii.Boolean(true),
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
type HeadersReferrerPolicy string

const (
	// The referrer policy is not set.
	HeadersReferrerPolicy_NO_REFERRER HeadersReferrerPolicy = "NO_REFERRER"
	// The referrer policy is no-referrer-when-downgrade.
	HeadersReferrerPolicy_NO_REFERRER_WHEN_DOWNGRADE HeadersReferrerPolicy = "NO_REFERRER_WHEN_DOWNGRADE"
	// The referrer policy is origin.
	HeadersReferrerPolicy_ORIGIN HeadersReferrerPolicy = "ORIGIN"
	// The referrer policy is origin-when-cross-origin.
	HeadersReferrerPolicy_ORIGIN_WHEN_CROSS_ORIGIN HeadersReferrerPolicy = "ORIGIN_WHEN_CROSS_ORIGIN"
	// The referrer policy is same-origin.
	HeadersReferrerPolicy_SAME_ORIGIN HeadersReferrerPolicy = "SAME_ORIGIN"
	// The referrer policy is strict-origin.
	HeadersReferrerPolicy_STRICT_ORIGIN HeadersReferrerPolicy = "STRICT_ORIGIN"
	// The referrer policy is strict-origin-when-cross-origin.
	HeadersReferrerPolicy_STRICT_ORIGIN_WHEN_CROSS_ORIGIN HeadersReferrerPolicy = "STRICT_ORIGIN_WHEN_CROSS_ORIGIN"
	// The referrer policy is unsafe-url.
	HeadersReferrerPolicy_UNSAFE_URL HeadersReferrerPolicy = "UNSAFE_URL"
)

