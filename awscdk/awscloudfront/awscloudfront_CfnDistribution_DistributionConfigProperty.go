package awscloudfront


// A distribution configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   distributionConfigProperty := &distributionConfigProperty{
//   	defaultCacheBehavior: &defaultCacheBehaviorProperty{
//   		targetOriginId: jsii.String("targetOriginId"),
//   		viewerProtocolPolicy: jsii.String("viewerProtocolPolicy"),
//
//   		// the properties below are optional
//   		allowedMethods: []*string{
//   			jsii.String("allowedMethods"),
//   		},
//   		cachedMethods: []*string{
//   			jsii.String("cachedMethods"),
//   		},
//   		cachePolicyId: jsii.String("cachePolicyId"),
//   		compress: jsii.Boolean(false),
//   		defaultTtl: jsii.Number(123),
//   		fieldLevelEncryptionId: jsii.String("fieldLevelEncryptionId"),
//   		forwardedValues: &forwardedValuesProperty{
//   			queryString: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			cookies: &cookiesProperty{
//   				forward: jsii.String("forward"),
//
//   				// the properties below are optional
//   				whitelistedNames: []*string{
//   					jsii.String("whitelistedNames"),
//   				},
//   			},
//   			headers: []*string{
//   				jsii.String("headers"),
//   			},
//   			queryStringCacheKeys: []*string{
//   				jsii.String("queryStringCacheKeys"),
//   			},
//   		},
//   		functionAssociations: []interface{}{
//   			&functionAssociationProperty{
//   				eventType: jsii.String("eventType"),
//   				functionArn: jsii.String("functionArn"),
//   			},
//   		},
//   		lambdaFunctionAssociations: []interface{}{
//   			&lambdaFunctionAssociationProperty{
//   				eventType: jsii.String("eventType"),
//   				includeBody: jsii.Boolean(false),
//   				lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   			},
//   		},
//   		maxTtl: jsii.Number(123),
//   		minTtl: jsii.Number(123),
//   		originRequestPolicyId: jsii.String("originRequestPolicyId"),
//   		realtimeLogConfigArn: jsii.String("realtimeLogConfigArn"),
//   		responseHeadersPolicyId: jsii.String("responseHeadersPolicyId"),
//   		smoothStreaming: jsii.Boolean(false),
//   		trustedKeyGroups: []*string{
//   			jsii.String("trustedKeyGroups"),
//   		},
//   		trustedSigners: []*string{
//   			jsii.String("trustedSigners"),
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	aliases: []*string{
//   		jsii.String("aliases"),
//   	},
//   	cacheBehaviors: []interface{}{
//   		&cacheBehaviorProperty{
//   			pathPattern: jsii.String("pathPattern"),
//   			targetOriginId: jsii.String("targetOriginId"),
//   			viewerProtocolPolicy: jsii.String("viewerProtocolPolicy"),
//
//   			// the properties below are optional
//   			allowedMethods: []*string{
//   				jsii.String("allowedMethods"),
//   			},
//   			cachedMethods: []*string{
//   				jsii.String("cachedMethods"),
//   			},
//   			cachePolicyId: jsii.String("cachePolicyId"),
//   			compress: jsii.Boolean(false),
//   			defaultTtl: jsii.Number(123),
//   			fieldLevelEncryptionId: jsii.String("fieldLevelEncryptionId"),
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
//   			functionAssociations: []interface{}{
//   				&functionAssociationProperty{
//   					eventType: jsii.String("eventType"),
//   					functionArn: jsii.String("functionArn"),
//   				},
//   			},
//   			lambdaFunctionAssociations: []interface{}{
//   				&lambdaFunctionAssociationProperty{
//   					eventType: jsii.String("eventType"),
//   					includeBody: jsii.Boolean(false),
//   					lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   				},
//   			},
//   			maxTtl: jsii.Number(123),
//   			minTtl: jsii.Number(123),
//   			originRequestPolicyId: jsii.String("originRequestPolicyId"),
//   			realtimeLogConfigArn: jsii.String("realtimeLogConfigArn"),
//   			responseHeadersPolicyId: jsii.String("responseHeadersPolicyId"),
//   			smoothStreaming: jsii.Boolean(false),
//   			trustedKeyGroups: []*string{
//   				jsii.String("trustedKeyGroups"),
//   			},
//   			trustedSigners: []*string{
//   				jsii.String("trustedSigners"),
//   			},
//   		},
//   	},
//   	cnamEs: []*string{
//   		jsii.String("cnamEs"),
//   	},
//   	comment: jsii.String("comment"),
//   	continuousDeploymentPolicyId: jsii.String("continuousDeploymentPolicyId"),
//   	customErrorResponses: []interface{}{
//   		&customErrorResponseProperty{
//   			errorCode: jsii.Number(123),
//
//   			// the properties below are optional
//   			errorCachingMinTtl: jsii.Number(123),
//   			responseCode: jsii.Number(123),
//   			responsePagePath: jsii.String("responsePagePath"),
//   		},
//   	},
//   	customOrigin: &legacyCustomOriginProperty{
//   		dnsName: jsii.String("dnsName"),
//   		originProtocolPolicy: jsii.String("originProtocolPolicy"),
//   		originSslProtocols: []*string{
//   			jsii.String("originSslProtocols"),
//   		},
//
//   		// the properties below are optional
//   		httpPort: jsii.Number(123),
//   		httpsPort: jsii.Number(123),
//   	},
//   	defaultRootObject: jsii.String("defaultRootObject"),
//   	httpVersion: jsii.String("httpVersion"),
//   	ipv6Enabled: jsii.Boolean(false),
//   	logging: &loggingProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		includeCookies: jsii.Boolean(false),
//   		prefix: jsii.String("prefix"),
//   	},
//   	originGroups: &originGroupsProperty{
//   		quantity: jsii.Number(123),
//
//   		// the properties below are optional
//   		items: []interface{}{
//   			&originGroupProperty{
//   				failoverCriteria: &originGroupFailoverCriteriaProperty{
//   					statusCodes: &statusCodesProperty{
//   						items: []interface{}{
//   							jsii.Number(123),
//   						},
//   						quantity: jsii.Number(123),
//   					},
//   				},
//   				id: jsii.String("id"),
//   				members: &originGroupMembersProperty{
//   					items: []interface{}{
//   						&originGroupMemberProperty{
//   							originId: jsii.String("originId"),
//   						},
//   					},
//   					quantity: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	origins: []interface{}{
//   		&originProperty{
//   			domainName: jsii.String("domainName"),
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			connectionAttempts: jsii.Number(123),
//   			connectionTimeout: jsii.Number(123),
//   			customOriginConfig: &customOriginConfigProperty{
//   				originProtocolPolicy: jsii.String("originProtocolPolicy"),
//
//   				// the properties below are optional
//   				httpPort: jsii.Number(123),
//   				httpsPort: jsii.Number(123),
//   				originKeepaliveTimeout: jsii.Number(123),
//   				originReadTimeout: jsii.Number(123),
//   				originSslProtocols: []*string{
//   					jsii.String("originSslProtocols"),
//   				},
//   			},
//   			originAccessControlId: jsii.String("originAccessControlId"),
//   			originCustomHeaders: []interface{}{
//   				&originCustomHeaderProperty{
//   					headerName: jsii.String("headerName"),
//   					headerValue: jsii.String("headerValue"),
//   				},
//   			},
//   			originPath: jsii.String("originPath"),
//   			originShield: &originShieldProperty{
//   				enabled: jsii.Boolean(false),
//   				originShieldRegion: jsii.String("originShieldRegion"),
//   			},
//   			s3OriginConfig: &s3OriginConfigProperty{
//   				originAccessIdentity: jsii.String("originAccessIdentity"),
//   			},
//   		},
//   	},
//   	priceClass: jsii.String("priceClass"),
//   	restrictions: &restrictionsProperty{
//   		geoRestriction: &geoRestrictionProperty{
//   			restrictionType: jsii.String("restrictionType"),
//
//   			// the properties below are optional
//   			locations: []*string{
//   				jsii.String("locations"),
//   			},
//   		},
//   	},
//   	s3Origin: &legacyS3OriginProperty{
//   		dnsName: jsii.String("dnsName"),
//
//   		// the properties below are optional
//   		originAccessIdentity: jsii.String("originAccessIdentity"),
//   	},
//   	staging: jsii.Boolean(false),
//   	viewerCertificate: &viewerCertificateProperty{
//   		acmCertificateArn: jsii.String("acmCertificateArn"),
//   		cloudFrontDefaultCertificate: jsii.Boolean(false),
//   		iamCertificateId: jsii.String("iamCertificateId"),
//   		minimumProtocolVersion: jsii.String("minimumProtocolVersion"),
//   		sslSupportMethod: jsii.String("sslSupportMethod"),
//   	},
//   	webAclId: jsii.String("webAclId"),
//   }
//
type CfnDistribution_DistributionConfigProperty struct {
	// A complex type that describes the default cache behavior if you don't specify a `CacheBehavior` element or if files don't match any of the values of `PathPattern` in `CacheBehavior` elements.
	//
	// You must create exactly one default cache behavior.
	DefaultCacheBehavior interface{} `field:"required" json:"defaultCacheBehavior" yaml:"defaultCacheBehavior"`
	// From this field, you can enable or disable the selected distribution.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// A complex type that contains information about CNAMEs (alternate domain names), if any, for this distribution.
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// A complex type that contains zero or more `CacheBehavior` elements.
	CacheBehaviors interface{} `field:"optional" json:"cacheBehaviors" yaml:"cacheBehaviors"`
	// `CfnDistribution.DistributionConfigProperty.CNAMEs`.
	CnamEs *[]*string `field:"optional" json:"cnamEs" yaml:"cnamEs"`
	// An optional comment to describe the distribution.
	//
	// The comment cannot be longer than 128 characters.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// `CfnDistribution.DistributionConfigProperty.ContinuousDeploymentPolicyId`.
	ContinuousDeploymentPolicyId *string `field:"optional" json:"continuousDeploymentPolicyId" yaml:"continuousDeploymentPolicyId"`
	// A complex type that controls the following:.
	//
	// - Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range with custom error messages before returning the response to the viewer.
	// - How long CloudFront caches HTTP status codes in the 4xx and 5xx range.
	//
	// For more information about custom error pages, see [Customizing Error Responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html) in the *Amazon CloudFront Developer Guide* .
	CustomErrorResponses interface{} `field:"optional" json:"customErrorResponses" yaml:"customErrorResponses"`
	// `CfnDistribution.DistributionConfigProperty.CustomOrigin`.
	CustomOrigin interface{} `field:"optional" json:"customOrigin" yaml:"customOrigin"`
	// The object that you want CloudFront to request from your origin (for example, `index.html` ) when a viewer requests the root URL for your distribution ( `http://www.example.com` ) instead of an object in your distribution ( `http://www.example.com/product-description.html` ). Specifying a default root object avoids exposing the contents of your distribution.
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
	DefaultRootObject *string `field:"optional" json:"defaultRootObject" yaml:"defaultRootObject"`
	// (Optional) Specify the maximum HTTP version that you want viewers to use to communicate with CloudFront .
	//
	// The default value for new web distributions is `http1.1` .
	//
	// For viewers and CloudFront to use HTTP/2, viewers must support TLS 1.2 or later, and must support server name identification (SNI).
	//
	// In general, configuring CloudFront to communicate with viewers using HTTP/2 reduces latency. You can improve performance by optimizing for HTTP/2.
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
	Ipv6Enabled interface{} `field:"optional" json:"ipv6Enabled" yaml:"ipv6Enabled"`
	// A complex type that controls whether access logs are written for the distribution.
	//
	// For more information about logging, see [Access Logs](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html) in the *Amazon CloudFront Developer Guide* .
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// A complex type that contains information about origin groups for this distribution.
	OriginGroups interface{} `field:"optional" json:"originGroups" yaml:"originGroups"`
	// A complex type that contains information about origins for this distribution.
	Origins interface{} `field:"optional" json:"origins" yaml:"origins"`
	// The price class that corresponds with the maximum price that you want to pay for CloudFront service.
	//
	// If you specify `PriceClass_All` , CloudFront responds to requests for your objects from all CloudFront edge locations.
	//
	// If you specify a price class other than `PriceClass_All` , CloudFront serves your objects from the CloudFront edge location that has the lowest latency among the edge locations in your price class. Viewers who are in or near regions that are excluded from your specified price class may encounter slower performance.
	//
	// For more information about price classes, see [Choosing the Price Class for a CloudFront Distribution](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PriceClass.html) in the *Amazon CloudFront Developer Guide* . For information about CloudFront pricing, including how price classes (such as Price Class 100) map to CloudFront regions, see [Amazon CloudFront Pricing](https://docs.aws.amazon.com/cloudfront/pricing/) .
	PriceClass *string `field:"optional" json:"priceClass" yaml:"priceClass"`
	// A complex type that identifies ways in which you want to restrict distribution of your content.
	Restrictions interface{} `field:"optional" json:"restrictions" yaml:"restrictions"`
	// `CfnDistribution.DistributionConfigProperty.S3Origin`.
	S3Origin interface{} `field:"optional" json:"s3Origin" yaml:"s3Origin"`
	// `CfnDistribution.DistributionConfigProperty.Staging`.
	Staging interface{} `field:"optional" json:"staging" yaml:"staging"`
	// A complex type that determines the distribution’s SSL/TLS configuration for communicating with viewers.
	ViewerCertificate interface{} `field:"optional" json:"viewerCertificate" yaml:"viewerCertificate"`
	// A unique identifier that specifies the AWS WAF web ACL, if any, to associate with this distribution.
	//
	// To specify a web ACL created using the latest version of AWS WAF , use the ACL ARN, for example `arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/473e64fd-f30b-4765-81a0-62ad96dd167a` . To specify a web ACL created using AWS WAF Classic, use the ACL ID, for example `473e64fd-f30b-4765-81a0-62ad96dd167a` .
	//
	// AWS WAF is a web application firewall that lets you monitor the HTTP and HTTPS requests that are forwarded to CloudFront, and lets you control access to your content. Based on conditions that you specify, such as the IP addresses that requests originate from or the values of query strings, CloudFront responds to requests either with the requested content or with an HTTP 403 status code (Forbidden). You can also configure CloudFront to return a custom error page when a request is blocked. For more information about AWS WAF , see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/what-is-aws-waf.html) .
	WebAclId *string `field:"optional" json:"webAclId" yaml:"webAclId"`
}

