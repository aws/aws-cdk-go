package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Determines whether CloudFront includes the Strict-Transport-Security HTTP response header and the header’s value.
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
type ResponseHeadersStrictTransportSecurity struct {
	// A number that CloudFront uses as the value for the max-age directive in the Strict-Transport-Security HTTP response header.
	AccessControlMaxAge awscdk.Duration `field:"required" json:"accessControlMaxAge" yaml:"accessControlMaxAge"`
	// A Boolean that determines whether CloudFront overrides the Strict-Transport-Security HTTP response header received from the origin with the one specified in this response headers policy.
	Override *bool `field:"required" json:"override" yaml:"override"`
	// A Boolean that determines whether CloudFront includes the includeSubDomains directive in the Strict-Transport-Security HTTP response header.
	// Default: false.
	//
	IncludeSubdomains *bool `field:"optional" json:"includeSubdomains" yaml:"includeSubdomains"`
	// A Boolean that determines whether CloudFront includes the preload directive in the Strict-Transport-Security HTTP response header.
	// Default: false.
	//
	Preload *bool `field:"optional" json:"preload" yaml:"preload"`
}

