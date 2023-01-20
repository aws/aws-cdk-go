package awscloudfront


// Properties for creating a Origin Request Policy.
//
// Example:
//   // Creating a custom origin request policy for a Distribution -- all parameters optional
//   var bucketOrigin s3Origin
//
//   myOriginRequestPolicy := cloudfront.NewOriginRequestPolicy(this, jsii.String("OriginRequestPolicy"), &originRequestPolicyProps{
//   	originRequestPolicyName: jsii.String("MyPolicy"),
//   	comment: jsii.String("A default policy"),
//   	cookieBehavior: cloudfront.originRequestCookieBehavior.none(),
//   	headerBehavior: cloudfront.originRequestHeaderBehavior.all(jsii.String("CloudFront-Is-Android-Viewer")),
//   	queryStringBehavior: cloudfront.originRequestQueryStringBehavior.allowList(jsii.String("username")),
//   })
//
//   cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: bucketOrigin,
//   		originRequestPolicy: myOriginRequestPolicy,
//   	},
//   })
//
type OriginRequestPolicyProps struct {
	// A comment to describe the origin request policy.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The cookies from viewer requests to include in origin requests.
	CookieBehavior OriginRequestCookieBehavior `field:"optional" json:"cookieBehavior" yaml:"cookieBehavior"`
	// The HTTP headers to include in origin requests.
	//
	// These can include headers from viewer requests and additional headers added by CloudFront.
	HeaderBehavior OriginRequestHeaderBehavior `field:"optional" json:"headerBehavior" yaml:"headerBehavior"`
	// A unique name to identify the origin request policy.
	//
	// The name must only include '-', '_', or alphanumeric characters.
	OriginRequestPolicyName *string `field:"optional" json:"originRequestPolicyName" yaml:"originRequestPolicyName"`
	// The URL query strings from viewer requests to include in origin requests.
	QueryStringBehavior OriginRequestQueryStringBehavior `field:"optional" json:"queryStringBehavior" yaml:"queryStringBehavior"`
}

