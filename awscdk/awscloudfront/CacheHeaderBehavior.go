package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Determines whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin.
//
// Example:
//   // Creating a custom cache policy for a Distribution -- all parameters optional
//   var bucketOrigin s3Origin
//
//   myCachePolicy := cloudfront.NewCachePolicy(this, jsii.String("myCachePolicy"), &CachePolicyProps{
//   	CachePolicyName: jsii.String("MyPolicy"),
//   	Comment: jsii.String("A default policy"),
//   	DefaultTtl: awscdk.Duration_Days(jsii.Number(2)),
//   	MinTtl: awscdk.Duration_Minutes(jsii.Number(1)),
//   	MaxTtl: awscdk.Duration_*Days(jsii.Number(10)),
//   	CookieBehavior: cloudfront.CacheCookieBehavior_All(),
//   	HeaderBehavior: cloudfront.CacheHeaderBehavior_AllowList(jsii.String("X-CustomHeader")),
//   	QueryStringBehavior: cloudfront.CacheQueryStringBehavior_DenyList(jsii.String("username")),
//   	EnableAcceptEncodingGzip: jsii.Boolean(true),
//   	EnableAcceptEncodingBrotli: jsii.Boolean(true),
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: bucketOrigin,
//   		CachePolicy: myCachePolicy,
//   	},
//   })
//
type CacheHeaderBehavior interface {
	// If no headers will be passed, or an allow list of headers.
	Behavior() *string
	// The headers for the allow/deny list, if applicable.
	Headers() *[]*string
}

// The jsii proxy struct for CacheHeaderBehavior
type jsiiProxy_CacheHeaderBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_CacheHeaderBehavior) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CacheHeaderBehavior) Headers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}


// Listed headers are included in the cache key and are automatically included in requests that CloudFront sends to the origin.
func CacheHeaderBehavior_AllowList(headers ...*string) CacheHeaderBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range headers {
		args = append(args, a)
	}

	var returns CacheHeaderBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CacheHeaderBehavior",
		"allowList",
		args,
		&returns,
	)

	return returns
}

// HTTP headers are not included in the cache key and are not automatically included in requests that CloudFront sends to the origin.
func CacheHeaderBehavior_None() CacheHeaderBehavior {
	_init_.Initialize()

	var returns CacheHeaderBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CacheHeaderBehavior",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

