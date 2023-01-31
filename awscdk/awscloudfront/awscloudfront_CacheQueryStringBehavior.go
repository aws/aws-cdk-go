package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Determines whether any URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin.
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
// Experimental.
type CacheQueryStringBehavior interface {
	// The behavior of query strings -- allow all, none, only an allow list, or a deny list.
	// Experimental.
	Behavior() *string
	// The query strings to allow or deny, if the behavior is an allow or deny list.
	// Experimental.
	QueryStrings() *[]*string
}

// The jsii proxy struct for CacheQueryStringBehavior
type jsiiProxy_CacheQueryStringBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_CacheQueryStringBehavior) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CacheQueryStringBehavior) QueryStrings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStrings",
		&returns,
	)
	return returns
}


// All query strings in viewer requests are included in the cache key and are automatically included in requests that CloudFront sends to the origin.
// Experimental.
func CacheQueryStringBehavior_All() CacheQueryStringBehavior {
	_init_.Initialize()

	var returns CacheQueryStringBehavior

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.CacheQueryStringBehavior",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Only the provided `queryStrings` are included in the cache key and automatically included in requests that CloudFront sends to the origin.
// Experimental.
func CacheQueryStringBehavior_AllowList(queryStrings ...*string) CacheQueryStringBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range queryStrings {
		args = append(args, a)
	}

	var returns CacheQueryStringBehavior

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.CacheQueryStringBehavior",
		"allowList",
		args,
		&returns,
	)

	return returns
}

// All query strings except the provided `queryStrings` are included in the cache key and automatically included in requests that CloudFront sends to the origin.
// Experimental.
func CacheQueryStringBehavior_DenyList(queryStrings ...*string) CacheQueryStringBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range queryStrings {
		args = append(args, a)
	}

	var returns CacheQueryStringBehavior

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.CacheQueryStringBehavior",
		"denyList",
		args,
		&returns,
	)

	return returns
}

// Query strings in viewer requests are not included in the cache key and are not automatically included in requests that CloudFront sends to the origin.
// Experimental.
func CacheQueryStringBehavior_None() CacheQueryStringBehavior {
	_init_.Initialize()

	var returns CacheQueryStringBehavior

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.CacheQueryStringBehavior",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

