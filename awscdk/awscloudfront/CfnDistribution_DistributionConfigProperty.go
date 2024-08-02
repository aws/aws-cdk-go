package awscloudfront


// A distribution configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   distributionConfigProperty := &DistributionConfigProperty{
//   	DefaultCacheBehavior: &DefaultCacheBehaviorProperty{
//   		TargetOriginId: jsii.String("targetOriginId"),
//   		ViewerProtocolPolicy: jsii.String("viewerProtocolPolicy"),
//
//   		// the properties below are optional
//   		AllowedMethods: []*string{
//   			jsii.String("allowedMethods"),
//   		},
//   		CachedMethods: []*string{
//   			jsii.String("cachedMethods"),
//   		},
//   		CachePolicyId: jsii.String("cachePolicyId"),
//   		Compress: jsii.Boolean(false),
//   		DefaultTtl: jsii.Number(123),
//   		FieldLevelEncryptionId: jsii.String("fieldLevelEncryptionId"),
//   		ForwardedValues: &ForwardedValuesProperty{
//   			QueryString: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			Cookies: &CookiesProperty{
//   				Forward: jsii.String("forward"),
//
//   				// the properties below are optional
//   				WhitelistedNames: []*string{
//   					jsii.String("whitelistedNames"),
//   				},
//   			},
//   			Headers: []*string{
//   				jsii.String("headers"),
//   			},
//   			QueryStringCacheKeys: []*string{
//   				jsii.String("queryStringCacheKeys"),
//   			},
//   		},
//   		FunctionAssociations: []interface{}{
//   			&FunctionAssociationProperty{
//   				EventType: jsii.String("eventType"),
//   				FunctionArn: jsii.String("functionArn"),
//   			},
//   		},
//   		LambdaFunctionAssociations: []interface{}{
//   			&LambdaFunctionAssociationProperty{
//   				EventType: jsii.String("eventType"),
//   				IncludeBody: jsii.Boolean(false),
//   				LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   			},
//   		},
//   		MaxTtl: jsii.Number(123),
//   		MinTtl: jsii.Number(123),
//   		OriginRequestPolicyId: jsii.String("originRequestPolicyId"),
//   		RealtimeLogConfigArn: jsii.String("realtimeLogConfigArn"),
//   		ResponseHeadersPolicyId: jsii.String("responseHeadersPolicyId"),
//   		SmoothStreaming: jsii.Boolean(false),
//   		TrustedKeyGroups: []*string{
//   			jsii.String("trustedKeyGroups"),
//   		},
//   		TrustedSigners: []*string{
//   			jsii.String("trustedSigners"),
//   		},
//   	},
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	Aliases: []*string{
//   		jsii.String("aliases"),
//   	},
//   	CacheBehaviors: []interface{}{
//   		&CacheBehaviorProperty{
//   			PathPattern: jsii.String("pathPattern"),
//   			TargetOriginId: jsii.String("targetOriginId"),
//   			ViewerProtocolPolicy: jsii.String("viewerProtocolPolicy"),
//
//   			// the properties below are optional
//   			AllowedMethods: []*string{
//   				jsii.String("allowedMethods"),
//   			},
//   			CachedMethods: []*string{
//   				jsii.String("cachedMethods"),
//   			},
//   			CachePolicyId: jsii.String("cachePolicyId"),
//   			Compress: jsii.Boolean(false),
//   			DefaultTtl: jsii.Number(123),
//   			FieldLevelEncryptionId: jsii.String("fieldLevelEncryptionId"),
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
//   			FunctionAssociations: []interface{}{
//   				&FunctionAssociationProperty{
//   					EventType: jsii.String("eventType"),
//   					FunctionArn: jsii.String("functionArn"),
//   				},
//   			},
//   			LambdaFunctionAssociations: []interface{}{
//   				&LambdaFunctionAssociationProperty{
//   					EventType: jsii.String("eventType"),
//   					IncludeBody: jsii.Boolean(false),
//   					LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   				},
//   			},
//   			MaxTtl: jsii.Number(123),
//   			MinTtl: jsii.Number(123),
//   			OriginRequestPolicyId: jsii.String("originRequestPolicyId"),
//   			RealtimeLogConfigArn: jsii.String("realtimeLogConfigArn"),
//   			ResponseHeadersPolicyId: jsii.String("responseHeadersPolicyId"),
//   			SmoothStreaming: jsii.Boolean(false),
//   			TrustedKeyGroups: []*string{
//   				jsii.String("trustedKeyGroups"),
//   			},
//   			TrustedSigners: []*string{
//   				jsii.String("trustedSigners"),
//   			},
//   		},
//   	},
//   	CnamEs: []*string{
//   		jsii.String("cnamEs"),
//   	},
//   	Comment: jsii.String("comment"),
//   	ContinuousDeploymentPolicyId: jsii.String("continuousDeploymentPolicyId"),
//   	CustomErrorResponses: []interface{}{
//   		&CustomErrorResponseProperty{
//   			ErrorCode: jsii.Number(123),
//
//   			// the properties below are optional
//   			ErrorCachingMinTtl: jsii.Number(123),
//   			ResponseCode: jsii.Number(123),
//   			ResponsePagePath: jsii.String("responsePagePath"),
//   		},
//   	},
//   	CustomOrigin: &LegacyCustomOriginProperty{
//   		DnsName: jsii.String("dnsName"),
//   		OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//   		OriginSslProtocols: []*string{
//   			jsii.String("originSslProtocols"),
//   		},
//
//   		// the properties below are optional
//   		HttpPort: jsii.Number(123),
//   		HttpsPort: jsii.Number(123),
//   	},
//   	DefaultRootObject: jsii.String("defaultRootObject"),
//   	HttpVersion: jsii.String("httpVersion"),
//   	Ipv6Enabled: jsii.Boolean(false),
//   	Logging: &LoggingProperty{
//   		Bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		IncludeCookies: jsii.Boolean(false),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	OriginGroups: &OriginGroupsProperty{
//   		Quantity: jsii.Number(123),
//
//   		// the properties below are optional
//   		Items: []interface{}{
//   			&OriginGroupProperty{
//   				FailoverCriteria: &OriginGroupFailoverCriteriaProperty{
//   					StatusCodes: &StatusCodesProperty{
//   						Items: []interface{}{
//   							jsii.Number(123),
//   						},
//   						Quantity: jsii.Number(123),
//   					},
//   				},
//   				Id: jsii.String("id"),
//   				Members: &OriginGroupMembersProperty{
//   					Items: []interface{}{
//   						&OriginGroupMemberProperty{
//   							OriginId: jsii.String("originId"),
//   						},
//   					},
//   					Quantity: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	Origins: []interface{}{
//   		&OriginProperty{
//   			DomainName: jsii.String("domainName"),
//   			Id: jsii.String("id"),
//
//   			// the properties below are optional
//   			ConnectionAttempts: jsii.Number(123),
//   			ConnectionTimeout: jsii.Number(123),
//   			CustomOriginConfig: &CustomOriginConfigProperty{
//   				OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//
//   				// the properties below are optional
//   				HttpPort: jsii.Number(123),
//   				HttpsPort: jsii.Number(123),
//   				OriginKeepaliveTimeout: jsii.Number(123),
//   				OriginReadTimeout: jsii.Number(123),
//   				OriginSslProtocols: []*string{
//   					jsii.String("originSslProtocols"),
//   				},
//   			},
//   			OriginAccessControlId: jsii.String("originAccessControlId"),
//   			OriginCustomHeaders: []interface{}{
//   				&OriginCustomHeaderProperty{
//   					HeaderName: jsii.String("headerName"),
//   					HeaderValue: jsii.String("headerValue"),
//   				},
//   			},
//   			OriginPath: jsii.String("originPath"),
//   			OriginShield: &OriginShieldProperty{
//   				Enabled: jsii.Boolean(false),
//   				OriginShieldRegion: jsii.String("originShieldRegion"),
//   			},
//   			S3OriginConfig: &S3OriginConfigProperty{
//   				OriginAccessIdentity: jsii.String("originAccessIdentity"),
//   			},
//   		},
//   	},
//   	PriceClass: jsii.String("priceClass"),
//   	Restrictions: &RestrictionsProperty{
//   		GeoRestriction: &GeoRestrictionProperty{
//   			RestrictionType: jsii.String("restrictionType"),
//
//   			// the properties below are optional
//   			Locations: []*string{
//   				jsii.String("locations"),
//   			},
//   		},
//   	},
//   	S3Origin: &LegacyS3OriginProperty{
//   		DnsName: jsii.String("dnsName"),
//
//   		// the properties below are optional
//   		OriginAccessIdentity: jsii.String("originAccessIdentity"),
//   	},
//   	Staging: jsii.Boolean(false),
//   	ViewerCertificate: &ViewerCertificateProperty{
//   		AcmCertificateArn: jsii.String("acmCertificateArn"),
//   		CloudFrontDefaultCertificate: jsii.Boolean(false),
//   		IamCertificateId: jsii.String("iamCertificateId"),
//   		MinimumProtocolVersion: jsii.String("minimumProtocolVersion"),
//   		SslSupportMethod: jsii.String("sslSupportMethod"),
//   	},
//   	WebAclId: jsii.String("webAclId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html
//
type CfnDistribution_DistributionConfigProperty struct {
	// A complex type that describes the default cache behavior if you don't specify a `CacheBehavior` element or if files don't match any of the values of `PathPattern` in `CacheBehavior` elements.
	//
	// You must create exactly one default cache behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-defaultcachebehavior
	//
	DefaultCacheBehavior interface{} `field:"required" json:"defaultCacheBehavior" yaml:"defaultCacheBehavior"`
	// From this field, you can enable or disable the selected distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// A complex type that contains information about CNAMEs (alternate domain names), if any, for this distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-aliases
	//
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// A complex type that contains zero or more `CacheBehavior` elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-cachebehaviors
	//
	CacheBehaviors interface{} `field:"optional" json:"cacheBehaviors" yaml:"cacheBehaviors"`
	// An alias for the CloudFront distribution's domain name.
	//
	// > This property is legacy. We recommend that you use [Aliases](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-aliases) instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-cnames
	//
	CnamEs *[]*string `field:"optional" json:"cnamEs" yaml:"cnamEs"`
	// A comment to describe the distribution.
	//
	// The comment cannot be longer than 128 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-comment
	//
	// Default: - "".
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The identifier of a continuous deployment policy.
	//
	// For more information, see `CreateContinuousDeploymentPolicy` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-continuousdeploymentpolicyid
	//
	ContinuousDeploymentPolicyId *string `field:"optional" json:"continuousDeploymentPolicyId" yaml:"continuousDeploymentPolicyId"`
	// A complex type that controls the following:.
	//
	// - Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range with custom error messages before returning the response to the viewer.
	// - How long CloudFront caches HTTP status codes in the 4xx and 5xx range.
	//
	// For more information about custom error pages, see [Customizing Error Responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-customerrorresponses
	//
	CustomErrorResponses interface{} `field:"optional" json:"customErrorResponses" yaml:"customErrorResponses"`
	// The user-defined HTTP server that serves as the origin for content that CloudFront distributes.
	//
	// > This property is legacy. We recommend that you use [Origin](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html) instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-customorigin
	//
	CustomOrigin interface{} `field:"optional" json:"customOrigin" yaml:"customOrigin"`
	// The object that you want CloudFront to request from your origin (for example, `index.html` ) when a viewer requests the root URL for your distribution ( `https://www.example.com` ) instead of an object in your distribution ( `https://www.example.com/product-description.html` ). Specifying a default root object avoids exposing the contents of your distribution.
	//
	// Specify only the object name, for example, `index.html` . Don't add a `/` before the object name.
	//
	// If you don't want to specify a default root object when you create a distribution, include an empty `DefaultRootObject` element.
	//
	// To delete the default root object from an existing distribution, update the distribution configuration and include an empty `DefaultRootObject` element.
	//
	// To replace the default root object, update the distribution configuration and specify the new object.
	//
	// For more information about the default root object, see [Creating a Default Root Object](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DefaultRootObject.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-defaultrootobject
	//
	// Default: - "".
	//
	DefaultRootObject *string `field:"optional" json:"defaultRootObject" yaml:"defaultRootObject"`
	// (Optional) Specify the HTTP version(s) that you want viewers to use to communicate with CloudFront .
	//
	// The default value for new distributions is `http1.1` .
	//
	// For viewers and CloudFront to use HTTP/2, viewers must support TLSv1.2 or later, and must support Server Name Indication (SNI).
	//
	// For viewers and CloudFront to use HTTP/3, viewers must support TLSv1.3 and Server Name Indication (SNI). CloudFront supports HTTP/3 connection migration to allow the viewer to switch networks without losing connection. For more information about connection migration, see [Connection Migration](https://docs.aws.amazon.com/https://www.rfc-editor.org/rfc/rfc9000.html#name-connection-migration) at RFC 9000. For more information about supported TLSv1.3 ciphers, see [Supported protocols and ciphers between viewers and CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/secure-connections-supported-viewer-protocols-ciphers.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-httpversion
	//
	// Default: - "http1.1"
	//
	HttpVersion *string `field:"optional" json:"httpVersion" yaml:"httpVersion"`
	// If you want CloudFront to respond to IPv6 DNS requests with an IPv6 address for your distribution, specify `true` .
	//
	// If you specify `false` , CloudFront responds to IPv6 DNS requests with the DNS response code `NOERROR` and with no IP addresses. This allows viewers to submit a second request, for an IPv4 address for your distribution.
	//
	// In general, you should enable IPv6 if you have users on IPv6 networks who want to access your content. However, if you're using signed URLs or signed cookies to restrict access to your content, and if you're using a custom policy that includes the `IpAddress` parameter to restrict the IP addresses that can access your content, don't enable IPv6. If you want to restrict access to some content by IP address and not restrict access to other content (or restrict access but not by IP address), you can create two distributions. For more information, see [Creating a Signed URL Using a Custom Policy](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-creating-signed-url-custom-policy.html) in the *Amazon CloudFront Developer Guide* .
	//
	// If you're using an Amazon Route 53 AWS Integration alias resource record set to route traffic to your CloudFront distribution, you need to create a second alias resource record set when both of the following are true:
	//
	// - You enable IPv6 for the distribution
	// - You're using alternate domain names in the URLs for your objects
	//
	// For more information, see [Routing Traffic to an Amazon CloudFront Web Distribution by Using Your Domain Name](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-to-cloudfront-distribution.html) in the *Amazon Route 53 AWS Integration Developer Guide* .
	//
	// If you created a CNAME resource record set, either with Amazon Route 53 AWS Integration or with another DNS service, you don't need to make any changes. A CNAME record will route traffic to your distribution regardless of the IP address format of the viewer request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-ipv6enabled
	//
	Ipv6Enabled interface{} `field:"optional" json:"ipv6Enabled" yaml:"ipv6Enabled"`
	// A complex type that controls whether access logs are written for the distribution.
	//
	// For more information about logging, see [Access Logs](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-logging
	//
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// A complex type that contains information about origin groups for this distribution.
	//
	// Specify a value for either the `Origins` or `OriginGroups` property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-origingroups
	//
	OriginGroups interface{} `field:"optional" json:"originGroups" yaml:"originGroups"`
	// A complex type that contains information about origins for this distribution.
	//
	// Specify a value for either the `Origins` or `OriginGroups` property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-origins
	//
	Origins interface{} `field:"optional" json:"origins" yaml:"origins"`
	// The price class that corresponds with the maximum price that you want to pay for CloudFront service.
	//
	// If you specify `PriceClass_All` , CloudFront responds to requests for your objects from all CloudFront edge locations.
	//
	// If you specify a price class other than `PriceClass_All` , CloudFront serves your objects from the CloudFront edge location that has the lowest latency among the edge locations in your price class. Viewers who are in or near regions that are excluded from your specified price class may encounter slower performance.
	//
	// For more information about price classes, see [Choosing the Price Class for a CloudFront Distribution](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PriceClass.html) in the *Amazon CloudFront Developer Guide* . For information about CloudFront pricing, including how price classes (such as Price Class 100) map to CloudFront regions, see [Amazon CloudFront Pricing](https://docs.aws.amazon.com/cloudfront/pricing/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-priceclass
	//
	// Default: - "PriceClass_All".
	//
	PriceClass *string `field:"optional" json:"priceClass" yaml:"priceClass"`
	// A complex type that identifies ways in which you want to restrict distribution of your content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-restrictions
	//
	Restrictions interface{} `field:"optional" json:"restrictions" yaml:"restrictions"`
	// The origin as an Amazon S3 bucket.
	//
	// > This property is legacy. We recommend that you use [Origin](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html) instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-s3origin
	//
	S3Origin interface{} `field:"optional" json:"s3Origin" yaml:"s3Origin"`
	// A Boolean that indicates whether this is a staging distribution.
	//
	// When this value is `true` , this is a staging distribution. When this value is `false` , this is not a staging distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-staging
	//
	Staging interface{} `field:"optional" json:"staging" yaml:"staging"`
	// A complex type that determines the distribution's SSL/TLS configuration for communicating with viewers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-viewercertificate
	//
	ViewerCertificate interface{} `field:"optional" json:"viewerCertificate" yaml:"viewerCertificate"`
	// A unique identifier that specifies the AWS WAF web ACL, if any, to associate with this distribution.
	//
	// To specify a web ACL created using the latest version of AWS WAF , use the ACL ARN, for example `arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111` . To specify a web ACL created using AWS WAF Classic, use the ACL ID, for example `a1b2c3d4-5678-90ab-cdef-EXAMPLE11111` .
	//
	// AWS WAF is a web application firewall that lets you monitor the HTTP and HTTPS requests that are forwarded to CloudFront, and lets you control access to your content. Based on conditions that you specify, such as the IP addresses that requests originate from or the values of query strings, CloudFront responds to requests either with the requested content or with an HTTP 403 status code (Forbidden). You can also configure CloudFront to return a custom error page when a request is blocked. For more information about AWS WAF , see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/what-is-aws-waf.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-webaclid
	//
	// Default: - "".
	//
	WebAclId *string `field:"optional" json:"webAclId" yaml:"webAclId"`
}

