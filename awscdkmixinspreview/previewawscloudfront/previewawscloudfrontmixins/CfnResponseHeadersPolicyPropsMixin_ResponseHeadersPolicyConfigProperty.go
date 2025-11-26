package previewawscloudfrontmixins


// A response headers policy configuration.
//
// A response headers policy configuration contains metadata about the response headers policy, and configurations for sets of HTTP response headers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   responseHeadersPolicyConfigProperty := &ResponseHeadersPolicyConfigProperty{
//   	Comment: jsii.String("comment"),
//   	CorsConfig: &CorsConfigProperty{
//   		AccessControlAllowCredentials: jsii.Boolean(false),
//   		AccessControlAllowHeaders: &AccessControlAllowHeadersProperty{
//   			Items: []*string{
//   				jsii.String("items"),
//   			},
//   		},
//   		AccessControlAllowMethods: &AccessControlAllowMethodsProperty{
//   			Items: []*string{
//   				jsii.String("items"),
//   			},
//   		},
//   		AccessControlAllowOrigins: &AccessControlAllowOriginsProperty{
//   			Items: []*string{
//   				jsii.String("items"),
//   			},
//   		},
//   		AccessControlExposeHeaders: &AccessControlExposeHeadersProperty{
//   			Items: []*string{
//   				jsii.String("items"),
//   			},
//   		},
//   		AccessControlMaxAgeSec: jsii.Number(123),
//   		OriginOverride: jsii.Boolean(false),
//   	},
//   	CustomHeadersConfig: &CustomHeadersConfigProperty{
//   		Items: []interface{}{
//   			&CustomHeaderProperty{
//   				Header: jsii.String("header"),
//   				Override: jsii.Boolean(false),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RemoveHeadersConfig: &RemoveHeadersConfigProperty{
//   		Items: []interface{}{
//   			&RemoveHeaderProperty{
//   				Header: jsii.String("header"),
//   			},
//   		},
//   	},
//   	SecurityHeadersConfig: &SecurityHeadersConfigProperty{
//   		ContentSecurityPolicy: &ContentSecurityPolicyProperty{
//   			ContentSecurityPolicy: jsii.String("contentSecurityPolicy"),
//   			Override: jsii.Boolean(false),
//   		},
//   		ContentTypeOptions: &ContentTypeOptionsProperty{
//   			Override: jsii.Boolean(false),
//   		},
//   		FrameOptions: &FrameOptionsProperty{
//   			FrameOption: jsii.String("frameOption"),
//   			Override: jsii.Boolean(false),
//   		},
//   		ReferrerPolicy: &ReferrerPolicyProperty{
//   			Override: jsii.Boolean(false),
//   			ReferrerPolicy: jsii.String("referrerPolicy"),
//   		},
//   		StrictTransportSecurity: &StrictTransportSecurityProperty{
//   			AccessControlMaxAgeSec: jsii.Number(123),
//   			IncludeSubdomains: jsii.Boolean(false),
//   			Override: jsii.Boolean(false),
//   			Preload: jsii.Boolean(false),
//   		},
//   		XssProtection: &XSSProtectionProperty{
//   			ModeBlock: jsii.Boolean(false),
//   			Override: jsii.Boolean(false),
//   			Protection: jsii.Boolean(false),
//   			ReportUri: jsii.String("reportUri"),
//   		},
//   	},
//   	ServerTimingHeadersConfig: &ServerTimingHeadersConfigProperty{
//   		Enabled: jsii.Boolean(false),
//   		SamplingRate: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-responseheaderspolicyconfig.html
//
type CfnResponseHeadersPolicyPropsMixin_ResponseHeadersPolicyConfigProperty struct {
	// A comment to describe the response headers policy.
	//
	// The comment cannot be longer than 128 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-responseheaderspolicyconfig.html#cfn-cloudfront-responseheaderspolicy-responseheaderspolicyconfig-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// A configuration for a set of HTTP response headers that are used for cross-origin resource sharing (CORS).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-responseheaderspolicyconfig.html#cfn-cloudfront-responseheaderspolicy-responseheaderspolicyconfig-corsconfig
	//
	CorsConfig interface{} `field:"optional" json:"corsConfig" yaml:"corsConfig"`
	// A configuration for a set of custom HTTP response headers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-responseheaderspolicyconfig.html#cfn-cloudfront-responseheaderspolicy-responseheaderspolicyconfig-customheadersconfig
	//
	CustomHeadersConfig interface{} `field:"optional" json:"customHeadersConfig" yaml:"customHeadersConfig"`
	// A name to identify the response headers policy.
	//
	// The name must be unique for response headers policies in this AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-responseheaderspolicyconfig.html#cfn-cloudfront-responseheaderspolicy-responseheaderspolicyconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A configuration for a set of HTTP headers to remove from the HTTP response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-responseheaderspolicyconfig.html#cfn-cloudfront-responseheaderspolicy-responseheaderspolicyconfig-removeheadersconfig
	//
	RemoveHeadersConfig interface{} `field:"optional" json:"removeHeadersConfig" yaml:"removeHeadersConfig"`
	// A configuration for a set of security-related HTTP response headers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-responseheaderspolicyconfig.html#cfn-cloudfront-responseheaderspolicy-responseheaderspolicyconfig-securityheadersconfig
	//
	SecurityHeadersConfig interface{} `field:"optional" json:"securityHeadersConfig" yaml:"securityHeadersConfig"`
	// A configuration for enabling the `Server-Timing` header in HTTP responses sent from CloudFront.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-responseheaderspolicyconfig.html#cfn-cloudfront-responseheaderspolicy-responseheaderspolicyconfig-servertimingheadersconfig
	//
	ServerTimingHeadersConfig interface{} `field:"optional" json:"serverTimingHeadersConfig" yaml:"serverTimingHeadersConfig"`
}

