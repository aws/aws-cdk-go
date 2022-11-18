package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Determines whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin.
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
type CacheCookieBehavior interface {
	// The behavior of cookies: allow all, none, an allow list, or a deny list.
	// Experimental.
	Behavior() *string
	// The cookies to allow or deny, if the behavior is an allow or deny list.
	// Experimental.
	Cookies() *[]*string
}

// The jsii proxy struct for CacheCookieBehavior
type jsiiProxy_CacheCookieBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_CacheCookieBehavior) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CacheCookieBehavior) Cookies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}


// All cookies in viewer requests are included in the cache key and are automatically included in requests that CloudFront sends to the origin.
// Experimental.
func CacheCookieBehavior_All() CacheCookieBehavior {
	_init_.Initialize()

	var returns CacheCookieBehavior

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.CacheCookieBehavior",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Only the provided `cookies` are included in the cache key and automatically included in requests that CloudFront sends to the origin.
// Experimental.
func CacheCookieBehavior_AllowList(cookies ...*string) CacheCookieBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range cookies {
		args = append(args, a)
	}

	var returns CacheCookieBehavior

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.CacheCookieBehavior",
		"allowList",
		args,
		&returns,
	)

	return returns
}

// All cookies except the provided `cookies` are included in the cache key and automatically included in requests that CloudFront sends to the origin.
// Experimental.
func CacheCookieBehavior_DenyList(cookies ...*string) CacheCookieBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range cookies {
		args = append(args, a)
	}

	var returns CacheCookieBehavior

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.CacheCookieBehavior",
		"denyList",
		args,
		&returns,
	)

	return returns
}

// Cookies in viewer requests are not included in the cache key and are not automatically included in requests that CloudFront sends to the origin.
// Experimental.
func CacheCookieBehavior_None() CacheCookieBehavior {
	_init_.Initialize()

	var returns CacheCookieBehavior

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.CacheCookieBehavior",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

