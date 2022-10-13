package awscloudfront


// Enum representing possible values of the Referrer-Policy HTTP response header.
//
// Example:
//   // Using an existing managed response headers policy
//   var bucketOrigin s3Origin
//
//   cloudfront.NewDistribution(this, jsii.String("myDistManagedPolicy"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: bucketOrigin,
//   		responseHeadersPolicy: cloudfront.responseHeadersPolicy_CORS_ALLOW_ALL_ORIGINS(),
//   	},
//   })
//
//   // Creating a custom response headers policy -- all parameters optional
//   myResponseHeadersPolicy := cloudfront.NewResponseHeadersPolicy(this, jsii.String("ResponseHeadersPolicy"), &responseHeadersPolicyProps{
//   	responseHeadersPolicyName: jsii.String("MyPolicy"),
//   	comment: jsii.String("A default policy"),
//   	corsBehavior: &responseHeadersCorsBehavior{
//   		accessControlAllowCredentials: jsii.Boolean(false),
//   		accessControlAllowHeaders: []*string{
//   			jsii.String("X-Custom-Header-1"),
//   			jsii.String("X-Custom-Header-2"),
//   		},
//   		accessControlAllowMethods: []*string{
//   			jsii.String("GET"),
//   			jsii.String("POST"),
//   		},
//   		accessControlAllowOrigins: []*string{
//   			jsii.String("*"),
//   		},
//   		accessControlExposeHeaders: []*string{
//   			jsii.String("X-Custom-Header-1"),
//   			jsii.String("X-Custom-Header-2"),
//   		},
//   		accessControlMaxAge: awscdk.Duration.seconds(jsii.Number(600)),
//   		originOverride: jsii.Boolean(true),
//   	},
//   	customHeadersBehavior: &responseCustomHeadersBehavior{
//   		customHeaders: []responseCustomHeader{
//   			&responseCustomHeader{
//   				header: jsii.String("X-Amz-Date"),
//   				value: jsii.String("some-value"),
//   				override: jsii.Boolean(true),
//   			},
//   			&responseCustomHeader{
//   				header: jsii.String("X-Amz-Security-Token"),
//   				value: jsii.String("some-value"),
//   				override: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	securityHeadersBehavior: &responseSecurityHeadersBehavior{
//   		contentSecurityPolicy: &responseHeadersContentSecurityPolicy{
//   			contentSecurityPolicy: jsii.String("default-src https:;"),
//   			override: jsii.Boolean(true),
//   		},
//   		contentTypeOptions: &responseHeadersContentTypeOptions{
//   			override: jsii.Boolean(true),
//   		},
//   		frameOptions: &responseHeadersFrameOptions{
//   			frameOption: cloudfront.headersFrameOption_DENY,
//   			override: jsii.Boolean(true),
//   		},
//   		referrerPolicy: &responseHeadersReferrerPolicy{
//   			referrerPolicy: cloudfront.headersReferrerPolicy_NO_REFERRER,
//   			override: jsii.Boolean(true),
//   		},
//   		strictTransportSecurity: &responseHeadersStrictTransportSecurity{
//   			accessControlMaxAge: awscdk.Duration.seconds(jsii.Number(600)),
//   			includeSubdomains: jsii.Boolean(true),
//   			override: jsii.Boolean(true),
//   		},
//   		xssProtection: &responseHeadersXSSProtection{
//   			protection: jsii.Boolean(true),
//   			modeBlock: jsii.Boolean(true),
//   			reportUri: jsii.String("https://example.com/csp-report"),
//   			override: jsii.Boolean(true),
//   		},
//   	},
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: bucketOrigin,
//   		responseHeadersPolicy: myResponseHeadersPolicy,
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

