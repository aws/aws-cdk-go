package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront"
)

// Options for creating a new behavior.
//
// Example:
//   var s3Bucket Bucket
//   // Add a cloudfront Function to a Distribution
//   cfFunction := cloudfront.NewFunction(this, jsii.String("Function"), &FunctionProps{
//   	Code: cloudfront.FunctionCode_FromInline(jsii.String("function handler(event) { return event.request }")),
//   	Runtime: cloudfront.FunctionRuntime_JS_2_0(),
//   })
//   cloudfront.NewDistribution(this, jsii.String("distro"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewS3Origin(s3Bucket),
//   		FunctionAssociations: []FunctionAssociation{
//   			&FunctionAssociation{
//   				Function: cfFunction,
//   				EventType: cloudfront.FunctionEventType_VIEWER_REQUEST,
//   			},
//   		},
//   	},
//   })
//
type BehaviorOptions struct {
	// HTTP methods to allow for this behavior.
	// Default: AllowedMethods.ALLOW_GET_HEAD
	//
	AllowedMethods AllowedMethods `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// HTTP methods to cache for this behavior.
	// Default: CachedMethods.CACHE_GET_HEAD
	//
	CachedMethods CachedMethods `field:"optional" json:"cachedMethods" yaml:"cachedMethods"`
	// The cache policy for this behavior.
	//
	// The cache policy determines what values are included in the cache key,
	// and the time-to-live (TTL) values for the cache.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html.
	//
	// Default: CachePolicy.CACHING_OPTIMIZED
	//
	CachePolicy interfacesawscloudfront.ICachePolicyRef `field:"optional" json:"cachePolicy" yaml:"cachePolicy"`
	// Whether you want CloudFront to automatically compress certain files for this cache behavior.
	//
	// See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/ServingCompressedFiles.html#compressed-content-cloudfront-file-types
	// for file types CloudFront will compress.
	// Default: true.
	//
	Compress *bool `field:"optional" json:"compress" yaml:"compress"`
	// The Lambda@Edge functions to invoke before serving the contents.
	// See: https://aws.amazon.com/lambda/edge
	//
	// Default: - no Lambda functions will be invoked.
	//
	EdgeLambdas *[]*EdgeLambda `field:"optional" json:"edgeLambdas" yaml:"edgeLambdas"`
	// Enables your CloudFront distribution to receive gRPC requests and to proxy them directly to your origins.
	//
	// If the `enableGrpc` is set to true, the following restrictions apply:
	// - The `allowedMethods` property must be `AllowedMethods.ALLOW_ALL` to include POST method because gRPC only supports POST method.
	// - The `httpVersion` property must be `HttpVersion.HTTP2` or `HttpVersion.HTTP2_AND_3` because gRPC only supports versions including HTTP/2.
	// - The `edgeLambdas` property can't be specified because gRPC is not supported with Lambda@Edge.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-using-grpc.html
	//
	// Default: false.
	//
	EnableGrpc *bool `field:"optional" json:"enableGrpc" yaml:"enableGrpc"`
	// The CloudFront functions to invoke before serving the contents.
	// Default: - no functions will be invoked.
	//
	FunctionAssociations *[]*FunctionAssociation `field:"optional" json:"functionAssociations" yaml:"functionAssociations"`
	// The origin request policy for this behavior.
	//
	// The origin request policy determines which values (e.g., headers, cookies)
	// are included in requests that CloudFront sends to the origin.
	// Default: - none.
	//
	OriginRequestPolicy interfacesawscloudfront.IOriginRequestPolicyRef `field:"optional" json:"originRequestPolicy" yaml:"originRequestPolicy"`
	// The real-time log configuration to be attached to this cache behavior.
	// Default: - none.
	//
	RealtimeLogConfig interfacesawscloudfront.IRealtimeLogConfigRef `field:"optional" json:"realtimeLogConfig" yaml:"realtimeLogConfig"`
	// The response headers policy for this behavior.
	//
	// The response headers policy determines which headers are included in responses.
	// Default: - none.
	//
	ResponseHeadersPolicy interfacesawscloudfront.IResponseHeadersPolicyRef `field:"optional" json:"responseHeadersPolicy" yaml:"responseHeadersPolicy"`
	// Set this to true to indicate you want to distribute media files in the Microsoft Smooth Streaming format using this behavior.
	// Default: false.
	//
	SmoothStreaming *bool `field:"optional" json:"smoothStreaming" yaml:"smoothStreaming"`
	// A list of Key Groups that CloudFront can use to validate signed URLs or signed cookies.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
	//
	// Default: - no KeyGroups are associated with cache behavior.
	//
	TrustedKeyGroups *[]interfacesawscloudfront.IKeyGroupRef `field:"optional" json:"trustedKeyGroups" yaml:"trustedKeyGroups"`
	// The protocol that viewers can use to access the files controlled by this behavior.
	// Default: ViewerProtocolPolicy.ALLOW_ALL
	//
	ViewerProtocolPolicy ViewerProtocolPolicy `field:"optional" json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
	// The origin that you want CloudFront to route requests to when they match this behavior.
	Origin IOrigin `field:"required" json:"origin" yaml:"origin"`
}

