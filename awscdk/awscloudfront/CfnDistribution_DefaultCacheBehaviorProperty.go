package awscloudfront


// A complex type that describes the default cache behavior if you don't specify a `CacheBehavior` element or if request URLs don't match any of the values of `PathPattern` in `CacheBehavior` elements.
//
// You must create exactly one default cache behavior.
//
// Example:
//   // Create the simple Origin
//   myBucket := s3.NewBucket(this, jsii.String("myBucket"))
//   s3Origin := origins.S3BucketOrigin_WithOriginAccessControl(myBucket, &S3BucketOriginWithOACProps{
//   	OriginAccessLevels: []accessLevel{
//   		cloudfront.*accessLevel_READ,
//   		cloudfront.*accessLevel_LIST,
//   	},
//   })
//
//   // Create the Distribution construct
//   myMultiTenantDistribution := cloudfront.NewDistribution(this, jsii.String("cf-hosted-distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: s3Origin,
//   	},
//   	DefaultRootObject: jsii.String("index.html"),
//   })
//
//   // Access the underlying L1 CfnDistribution to configure SaaS Manager properties which are not yet available in the L2 Distribution construct
//   cfnDistribution := myMultiTenantDistribution.Node.defaultChild.(cfnDistribution)
//
//   defaultCacheBehavior := &DefaultCacheBehaviorProperty{
//   	TargetOriginId: myBucket.BucketArn,
//   	ViewerProtocolPolicy: jsii.String("allow-all"),
//   	Compress: jsii.Boolean(false),
//   	AllowedMethods: []*string{
//   		jsii.String("GET"),
//   		jsii.String("HEAD"),
//   	},
//   	CachePolicyId: cloudfront.CachePolicy_CACHING_OPTIMIZED().CachePolicyId,
//   }
//   // Create the updated distributionConfig
//   distributionConfig := &DistributionConfigProperty{
//   	DefaultCacheBehavior: defaultCacheBehavior,
//   	Enabled: jsii.Boolean(true),
//   	// the properties below are optional
//   	ConnectionMode: jsii.String("tenant-only"),
//   	Origins: []interface{}{
//   		&OriginProperty{
//   			Id: myBucket.*BucketArn,
//   			DomainName: myBucket.BucketDomainName,
//   			S3OriginConfig: &S3OriginConfigProperty{
//   			},
//   			OriginPath: jsii.String("/{{tenantName}}"),
//   		},
//   	},
//   	TenantConfig: &TenantConfigProperty{
//   		ParameterDefinitions: []interface{}{
//   			&ParameterDefinitionProperty{
//   				Definition: &DefinitionProperty{
//   					StringSchema: &StringSchemaProperty{
//   						Required: jsii.Boolean(false),
//   						// the properties below are optional
//   						Comment: jsii.String("tenantName"),
//   						DefaultValue: jsii.String("root"),
//   					},
//   				},
//   				Name: jsii.String("tenantName"),
//   			},
//   		},
//   	},
//   }
//
//   // Override the distribution configuration to enable multi-tenancy.
//   cfnDistribution.DistributionConfig = distributionConfig
//
//   // Create a connection group so we have access to the RoutingEndpoint associated with the tenant we are about to create
//   connectionGroup := cloudfront.NewCfnConnectionGroup(this, jsii.String("self-hosted-connection-group"), &CfnConnectionGroupProps{
//   	Enabled: jsii.Boolean(true),
//   	Ipv6Enabled: jsii.Boolean(true),
//   	Name: jsii.String("self-hosted-connection-group"),
//   })
//
//   // Export the RoutingEndpoint, skip this step if you'd prefer to fetch it from the CloudFront console or via Cloudfront.ListConnectionGroups API
//   // Export the RoutingEndpoint, skip this step if you'd prefer to fetch it from the CloudFront console or via Cloudfront.ListConnectionGroups API
//   awscdk.NewCfnOutput(this, jsii.String("RoutingEndpoint"), &CfnOutputProps{
//   	Value: connectionGroup.AttrRoutingEndpoint,
//   	Description: jsii.String("CloudFront Routing Endpoint to be added to my hosted zone CNAME records"),
//   })
//
//   // Create a distribution tenant with a self-hosted domain.
//   selfHostedTenant := cloudfront.NewCfnDistributionTenant(this, jsii.String("self-hosted-tenant"), &CfnDistributionTenantProps{
//   	DistributionId: myMultiTenantDistribution.DistributionId,
//   	ConnectionGroupId: connectionGroup.AttrId,
//   	Name: jsii.String("self-hosted-tenant"),
//   	Domains: []*string{
//   		jsii.String("self-hosted-tenant.my.domain.com"),
//   	},
//   	Enabled: jsii.Boolean(true),
//   	ManagedCertificateRequest: &ManagedCertificateRequestProperty{
//   		PrimaryDomainName: jsii.String("self-hosted-tenant.my.domain.com"),
//   		ValidationTokenHost: jsii.String("self-hosted"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html
//
type CfnDistribution_DefaultCacheBehaviorProperty struct {
	// The value of `ID` for the origin that you want CloudFront to route requests to when they use the default cache behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-targetoriginid
	//
	TargetOriginId *string `field:"required" json:"targetOriginId" yaml:"targetOriginId"`
	// The protocol that viewers can use to access the files in the origin specified by `TargetOriginId` when a request matches the path pattern in `PathPattern` .
	//
	// You can specify the following options:
	//
	// - `allow-all` : Viewers can use HTTP or HTTPS.
	// - `redirect-to-https` : If a viewer submits an HTTP request, CloudFront returns an HTTP status code of 301 (Moved Permanently) to the viewer along with the HTTPS URL. The viewer then resubmits the request using the new URL.
	// - `https-only` : If a viewer sends an HTTP request, CloudFront returns an HTTP status code of 403 (Forbidden).
	//
	// For more information about requiring the HTTPS protocol, see [Requiring HTTPS Between Viewers and CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https-viewers-to-cloudfront.html) in the *Amazon CloudFront Developer Guide* .
	//
	// > The only way to guarantee that viewers retrieve an object that was fetched from the origin using HTTPS is never to use any other protocol to fetch the object. If you have recently changed from HTTP to HTTPS, we recommend that you clear your objects' cache because cached objects are protocol agnostic. That means that an edge location will return an object from the cache regardless of whether the current request protocol matches the protocol used previously. For more information, see [Managing Cache Expiration](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-viewerprotocolpolicy
	//
	ViewerProtocolPolicy *string `field:"required" json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
	// A complex type that controls which HTTP methods CloudFront processes and forwards to your Amazon S3 bucket or your custom origin.
	//
	// There are three choices:
	//
	// - CloudFront forwards only `GET` and `HEAD` requests.
	// - CloudFront forwards only `GET` , `HEAD` , and `OPTIONS` requests.
	// - CloudFront forwards `GET, HEAD, OPTIONS, PUT, PATCH, POST` , and `DELETE` requests.
	//
	// If you pick the third choice, you may need to restrict access to your Amazon S3 bucket or to your custom origin so users can't perform operations that you don't want them to. For example, you might not want users to have permissions to delete objects from your origin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-allowedmethods
	//
	AllowedMethods *[]*string `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// A complex type that controls whether CloudFront caches the response to requests using the specified HTTP methods.
	//
	// There are two choices:
	//
	// - CloudFront caches responses to `GET` and `HEAD` requests.
	// - CloudFront caches responses to `GET` , `HEAD` , and `OPTIONS` requests.
	//
	// If you pick the second choice for your Amazon S3 Origin, you may need to forward Access-Control-Request-Method, Access-Control-Request-Headers, and Origin headers for the responses to be cached correctly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-cachedmethods
	//
	CachedMethods *[]*string `field:"optional" json:"cachedMethods" yaml:"cachedMethods"`
	// The unique identifier of the cache policy that is attached to the default cache behavior.
	//
	// For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// A `DefaultCacheBehavior` must include either a `CachePolicyId` or `ForwardedValues` . We recommend that you use a `CachePolicyId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-cachepolicyid
	//
	// Default: - "".
	//
	CachePolicyId *string `field:"optional" json:"cachePolicyId" yaml:"cachePolicyId"`
	// Whether you want CloudFront to automatically compress certain files for this cache behavior.
	//
	// If so, specify `true` ; if not, specify `false` . For more information, see [Serving Compressed Files](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/ServingCompressedFiles.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-compress
	//
	// Default: - false.
	//
	Compress interface{} `field:"optional" json:"compress" yaml:"compress"`
	// > This field only supports standard distributions.
	//
	// You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-config-options.html#unsupported-saas) in the *Amazon CloudFront Developer Guide* .
	//
	// This field is deprecated. We recommend that you use the `DefaultTTL` field in a cache policy instead of this field. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// The default amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated. The value that you specify applies only when your origin does not add HTTP headers such as `Cache-Control max-age` , `Cache-Control s-maxage` , and `Expires` to objects. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-defaultttl
	//
	// Default: - 86400.
	//
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// The value of `ID` for the field-level encryption configuration that you want CloudFront to use for encrypting specific fields of data for the default cache behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-fieldlevelencryptionid
	//
	// Default: - "".
	//
	FieldLevelEncryptionId *string `field:"optional" json:"fieldLevelEncryptionId" yaml:"fieldLevelEncryptionId"`
	// This field is deprecated.
	//
	// We recommend that you use a cache policy or an origin request policy instead of this field. For more information, see [Working with policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/working-with-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// If you want to include values in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// If you want to send values to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) or [Using the managed origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-origin-request-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// A `DefaultCacheBehavior` must include either a `CachePolicyId` or `ForwardedValues` . We recommend that you use a `CachePolicyId` .
	//
	// A complex type that specifies how CloudFront handles query strings, cookies, and HTTP headers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-forwardedvalues
	//
	ForwardedValues interface{} `field:"optional" json:"forwardedValues" yaml:"forwardedValues"`
	// A list of CloudFront functions that are associated with this cache behavior.
	//
	// Your functions must be published to the `LIVE` stage to associate them with a cache behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-functionassociations
	//
	FunctionAssociations interface{} `field:"optional" json:"functionAssociations" yaml:"functionAssociations"`
	// The gRPC configuration for your cache behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-grpcconfig
	//
	GrpcConfig interface{} `field:"optional" json:"grpcConfig" yaml:"grpcConfig"`
	// A complex type that contains zero or more Lambda@Edge function associations for a cache behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-lambdafunctionassociations
	//
	LambdaFunctionAssociations interface{} `field:"optional" json:"lambdaFunctionAssociations" yaml:"lambdaFunctionAssociations"`
	// > This field only supports standard distributions.
	//
	// You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-config-options.html#unsupported-saas) in the *Amazon CloudFront Developer Guide* .
	//
	// This field is deprecated. We recommend that you use the `MaxTTL` field in a cache policy instead of this field. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// The maximum amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated. The value that you specify applies only when your origin adds HTTP headers such as `Cache-Control max-age` , `Cache-Control s-maxage` , and `Expires` to objects. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-maxttl
	//
	// Default: - 31536000.
	//
	MaxTtl *float64 `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// > This field only supports standard distributions.
	//
	// You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-config-options.html#unsupported-saas) in the *Amazon CloudFront Developer Guide* .
	//
	// This field is deprecated. We recommend that you use the `MinTTL` field in a cache policy instead of this field. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// The minimum amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	//
	// You must specify `0` for `MinTTL` if you configure CloudFront to forward all headers to your origin (under `Headers` , if you specify `1` for `Quantity` and `*` for `Name` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-minttl
	//
	// Default: - 0.
	//
	MinTtl *float64 `field:"optional" json:"minTtl" yaml:"minTtl"`
	// The unique identifier of the origin request policy that is attached to the default cache behavior.
	//
	// For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) or [Using the managed origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-origin-request-policies.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-originrequestpolicyid
	//
	// Default: - "".
	//
	OriginRequestPolicyId *string `field:"optional" json:"originRequestPolicyId" yaml:"originRequestPolicyId"`
	// The Amazon Resource Name (ARN) of the real-time log configuration that is attached to this cache behavior.
	//
	// For more information, see [Real-time logs](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-realtimelogconfigarn
	//
	// Default: - "".
	//
	RealtimeLogConfigArn *string `field:"optional" json:"realtimeLogConfigArn" yaml:"realtimeLogConfigArn"`
	// The identifier for a response headers policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-responseheaderspolicyid
	//
	// Default: - "".
	//
	ResponseHeadersPolicyId *string `field:"optional" json:"responseHeadersPolicyId" yaml:"responseHeadersPolicyId"`
	// > This field only supports standard distributions.
	//
	// You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-config-options.html#unsupported-saas) in the *Amazon CloudFront Developer Guide* .
	//
	// Indicates whether you want to distribute media files in the Microsoft Smooth Streaming format using the origin that is associated with this cache behavior. If so, specify `true` ; if not, specify `false` . If you specify `true` for `SmoothStreaming` , you can still distribute other content using this cache behavior if the content matches the value of `PathPattern` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-smoothstreaming
	//
	// Default: - false.
	//
	SmoothStreaming interface{} `field:"optional" json:"smoothStreaming" yaml:"smoothStreaming"`
	// A list of key groups that CloudFront can use to validate signed URLs or signed cookies.
	//
	// When a cache behavior contains trusted key groups, CloudFront requires signed URLs or signed cookies for all requests that match the cache behavior. The URLs or cookies must be signed with a private key whose corresponding public key is in the key group. The signed URL or cookie contains information about which public key CloudFront should use to verify the signature. For more information, see [Serving private content](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-trustedkeygroups
	//
	TrustedKeyGroups *[]*string `field:"optional" json:"trustedKeyGroups" yaml:"trustedKeyGroups"`
	// > We recommend using `TrustedKeyGroups` instead of `TrustedSigners` .
	//
	// > This field only supports standard distributions. You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-config-options.html#unsupported-saas) in the *Amazon CloudFront Developer Guide* .
	//
	// A list of AWS account IDs whose public keys CloudFront can use to validate signed URLs or signed cookies.
	//
	// When a cache behavior contains trusted signers, CloudFront requires signed URLs or signed cookies for all requests that match the cache behavior. The URLs or cookies must be signed with the private key of a CloudFront key pair in a trusted signer's AWS account . The signed URL or cookie contains information about which public key CloudFront should use to verify the signature. For more information, see [Serving private content](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html#cfn-cloudfront-distribution-defaultcachebehavior-trustedsigners
	//
	TrustedSigners *[]*string `field:"optional" json:"trustedSigners" yaml:"trustedSigners"`
}

