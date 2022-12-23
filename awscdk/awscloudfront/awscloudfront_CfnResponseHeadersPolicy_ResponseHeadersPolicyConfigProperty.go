package awscloudfront


// A response headers policy configuration.
//
// A response headers policy configuration contains metadata about the response headers policy, and configurations for sets of HTTP response headers and their values. CloudFront adds the headers in the policy to HTTP responses that it sends for requests that match a cache behavior associated with the policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responseHeadersPolicyConfigProperty := &responseHeadersPolicyConfigProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	comment: jsii.String("comment"),
//   	corsConfig: &corsConfigProperty{
//   		accessControlAllowCredentials: jsii.Boolean(false),
//   		accessControlAllowHeaders: &accessControlAllowHeadersProperty{
//   			items: []*string{
//   				jsii.String("items"),
//   			},
//   		},
//   		accessControlAllowMethods: &accessControlAllowMethodsProperty{
//   			items: []*string{
//   				jsii.String("items"),
//   			},
//   		},
//   		accessControlAllowOrigins: &accessControlAllowOriginsProperty{
//   			items: []*string{
//   				jsii.String("items"),
//   			},
//   		},
//   		originOverride: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		accessControlExposeHeaders: &accessControlExposeHeadersProperty{
//   			items: []*string{
//   				jsii.String("items"),
//   			},
//   		},
//   		accessControlMaxAgeSec: jsii.Number(123),
//   	},
//   	customHeadersConfig: &customHeadersConfigProperty{
//   		items: []interface{}{
//   			&customHeaderProperty{
//   				header: jsii.String("header"),
//   				override: jsii.Boolean(false),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	securityHeadersConfig: &securityHeadersConfigProperty{
//   		contentSecurityPolicy: &contentSecurityPolicyProperty{
//   			contentSecurityPolicy: jsii.String("contentSecurityPolicy"),
//   			override: jsii.Boolean(false),
//   		},
//   		contentTypeOptions: &contentTypeOptionsProperty{
//   			override: jsii.Boolean(false),
//   		},
//   		frameOptions: &frameOptionsProperty{
//   			frameOption: jsii.String("frameOption"),
//   			override: jsii.Boolean(false),
//   		},
//   		referrerPolicy: &referrerPolicyProperty{
//   			override: jsii.Boolean(false),
//   			referrerPolicy: jsii.String("referrerPolicy"),
//   		},
//   		strictTransportSecurity: &strictTransportSecurityProperty{
//   			accessControlMaxAgeSec: jsii.Number(123),
//   			override: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			includeSubdomains: jsii.Boolean(false),
//   			preload: jsii.Boolean(false),
//   		},
//   		xssProtection: &xSSProtectionProperty{
//   			override: jsii.Boolean(false),
//   			protection: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			modeBlock: jsii.Boolean(false),
//   			reportUri: jsii.String("reportUri"),
//   		},
//   	},
//   	serverTimingHeadersConfig: &serverTimingHeadersConfigProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		samplingRate: jsii.Number(123),
//   	},
//   }
//
type CfnResponseHeadersPolicy_ResponseHeadersPolicyConfigProperty struct {
	// A name to identify the response headers policy.
	//
	// The name must be unique for response headers policies in this AWS account .
	Name *string `field:"required" json:"name" yaml:"name"`
	// A comment to describe the response headers policy.
	//
	// The comment cannot be longer than 128 characters.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// A configuration for a set of HTTP response headers that are used for cross-origin resource sharing (CORS).
	CorsConfig interface{} `field:"optional" json:"corsConfig" yaml:"corsConfig"`
	// A configuration for a set of custom HTTP response headers.
	CustomHeadersConfig interface{} `field:"optional" json:"customHeadersConfig" yaml:"customHeadersConfig"`
	// A configuration for a set of security-related HTTP response headers.
	SecurityHeadersConfig interface{} `field:"optional" json:"securityHeadersConfig" yaml:"securityHeadersConfig"`
	// `CfnResponseHeadersPolicy.ResponseHeadersPolicyConfigProperty.ServerTimingHeadersConfig`.
	ServerTimingHeadersConfig interface{} `field:"optional" json:"serverTimingHeadersConfig" yaml:"serverTimingHeadersConfig"`
}

