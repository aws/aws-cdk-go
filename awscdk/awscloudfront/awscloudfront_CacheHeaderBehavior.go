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
//   myCachePolicy := cloudfront.NewCachePolicy(this, jsii.String("myCachePolicy"), &cachePolicyProps{
//   	cachePolicyName: jsii.String("MyPolicy"),
//   	comment: jsii.String("A default policy"),
//   	defaultTtl: awscdk.Duration.days(jsii.Number(2)),
//   	minTtl: awscdk.Duration.minutes(jsii.Number(1)),
//   	maxTtl: awscdk.Duration.days(jsii.Number(10)),
//   	cookieBehavior: cloudfront.cacheCookieBehavior.all(),
//   	headerBehavior: cloudfront.cacheHeaderBehavior.allowList(jsii.String("X-CustomHeader")),
//   	queryStringBehavior: cloudfront.cacheQueryStringBehavior.denyList(jsii.String("username")),
//   	enableAcceptEncodingGzip: jsii.Boolean(true),
//   	enableAcceptEncodingBrotli: jsii.Boolean(true),
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: bucketOrigin,
//   		cachePolicy: myCachePolicy,
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

