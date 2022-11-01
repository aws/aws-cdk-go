package awscloudfront


// Options for creating a new behavior.
//
// Example:
//   var s3Bucket bucket
//   // Add a cloudfront Function to a Distribution
//   cfFunction := cloudfront.NewFunction(this, jsii.String("Function"), &functionProps{
//   	code: cloudfront.functionCode.fromInline(jsii.String("function handler(event) { return event.request }")),
//   })
//   cloudfront.NewDistribution(this, jsii.String("distro"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewS3Origin(s3Bucket),
//   		functionAssociations: []functionAssociation{
//   			&functionAssociation{
//   				function: cfFunction,
//   				eventType: cloudfront.functionEventType_VIEWER_REQUEST,
//   			},
//   		},
//   	},
//   })
//
type BehaviorOptions struct {
	// HTTP methods to allow for this behavior.
	AllowedMethods AllowedMethods `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// HTTP methods to cache for this behavior.
	CachedMethods CachedMethods `field:"optional" json:"cachedMethods" yaml:"cachedMethods"`
	// The cache policy for this behavior.
	//
	// The cache policy determines what values are included in the cache key,
	// and the time-to-live (TTL) values for the cache.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html.
	//
	CachePolicy ICachePolicy `field:"optional" json:"cachePolicy" yaml:"cachePolicy"`
	// Whether you want CloudFront to automatically compress certain files for this cache behavior.
	//
	// See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/ServingCompressedFiles.html#compressed-content-cloudfront-file-types
	// for file types CloudFront will compress.
	Compress *bool `field:"optional" json:"compress" yaml:"compress"`
	// The Lambda@Edge functions to invoke before serving the contents.
	// See: https://aws.amazon.com/lambda/edge
	//
	EdgeLambdas *[]*EdgeLambda `field:"optional" json:"edgeLambdas" yaml:"edgeLambdas"`
	// The CloudFront functions to invoke before serving the contents.
	FunctionAssociations *[]*FunctionAssociation `field:"optional" json:"functionAssociations" yaml:"functionAssociations"`
	// The origin request policy for this behavior.
	//
	// The origin request policy determines which values (e.g., headers, cookies)
	// are included in requests that CloudFront sends to the origin.
	OriginRequestPolicy IOriginRequestPolicy `field:"optional" json:"originRequestPolicy" yaml:"originRequestPolicy"`
	// The response headers policy for this behavior.
	//
	// The response headers policy determines which headers are included in responses.
	ResponseHeadersPolicy IResponseHeadersPolicy `field:"optional" json:"responseHeadersPolicy" yaml:"responseHeadersPolicy"`
	// Set this to true to indicate you want to distribute media files in the Microsoft Smooth Streaming format using this behavior.
	SmoothStreaming *bool `field:"optional" json:"smoothStreaming" yaml:"smoothStreaming"`
	// A list of Key Groups that CloudFront can use to validate signed URLs or signed cookies.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
	//
	TrustedKeyGroups *[]IKeyGroup `field:"optional" json:"trustedKeyGroups" yaml:"trustedKeyGroups"`
	// The protocol that viewers can use to access the files controlled by this behavior.
	ViewerProtocolPolicy ViewerProtocolPolicy `field:"optional" json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
	// The origin that you want CloudFront to route requests to when they match this behavior.
	Origin IOrigin `field:"required" json:"origin" yaml:"origin"`
}

