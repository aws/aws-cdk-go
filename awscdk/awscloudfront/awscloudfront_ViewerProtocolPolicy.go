package awscloudfront


// How HTTPs should be handled with your distribution.
//
// Example:
//   // Create a Distribution with configured HTTP methods and viewer protocol policy of the cache.
//   var myBucket bucket
//
//   myWebDistribution := cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewS3Origin(myBucket),
//   		AllowedMethods: cloudfront.AllowedMethods_ALLOW_ALL(),
//   		ViewerProtocolPolicy: cloudfront.ViewerProtocolPolicy_REDIRECT_TO_HTTPS,
//   	},
//   })
//
// Experimental.
type ViewerProtocolPolicy string

const (
	// HTTPS only.
	// Experimental.
	ViewerProtocolPolicy_HTTPS_ONLY ViewerProtocolPolicy = "HTTPS_ONLY"
	// Will redirect HTTP requests to HTTPS.
	// Experimental.
	ViewerProtocolPolicy_REDIRECT_TO_HTTPS ViewerProtocolPolicy = "REDIRECT_TO_HTTPS"
	// Both HTTP and HTTPS supported.
	// Experimental.
	ViewerProtocolPolicy_ALLOW_ALL ViewerProtocolPolicy = "ALLOW_ALL"
)

