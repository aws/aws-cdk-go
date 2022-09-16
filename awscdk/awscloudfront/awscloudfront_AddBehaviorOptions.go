package awscloudfront


// Options for adding a new behavior to a Distribution.
//
// Example:
//   // Add a behavior to a Distribution after initial creation.
//   var myBucket bucket
//   var myWebDistribution distribution
//
//   myWebDistribution.addBehavior(jsii.String("/images/*.jpg"), origins.NewS3Origin(myBucket), &addBehaviorOptions{
//   	viewerProtocolPolicy: cloudfront.viewerProtocolPolicy_REDIRECT_TO_HTTPS,
//   })
//
type AddBehaviorOptions struct {
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
}

