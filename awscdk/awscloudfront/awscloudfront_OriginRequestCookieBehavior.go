package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Determines whether any cookies in viewer requests (and if so, which cookies) are included in requests that CloudFront sends to the origin.
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
type OriginRequestCookieBehavior interface {
	// The behavior of cookies: allow all, none or an allow list.
	Behavior() *string
	// The cookies to allow, if the behavior is an allow list.
	Cookies() *[]*string
}

// The jsii proxy struct for OriginRequestCookieBehavior
type jsiiProxy_OriginRequestCookieBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_OriginRequestCookieBehavior) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginRequestCookieBehavior) Cookies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}


// All cookies in viewer requests are included in requests that CloudFront sends to the origin.
func OriginRequestCookieBehavior_All() OriginRequestCookieBehavior {
	_init_.Initialize()

	var returns OriginRequestCookieBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestCookieBehavior",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Only the provided `cookies` are included in requests that CloudFront sends to the origin.
func OriginRequestCookieBehavior_AllowList(cookies ...*string) OriginRequestCookieBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range cookies {
		args = append(args, a)
	}

	var returns OriginRequestCookieBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestCookieBehavior",
		"allowList",
		args,
		&returns,
	)

	return returns
}

// Cookies in viewer requests are not included in requests that CloudFront sends to the origin.
//
// Any cookies that are listed in a CachePolicy are still included in origin requests.
func OriginRequestCookieBehavior_None() OriginRequestCookieBehavior {
	_init_.Initialize()

	var returns OriginRequestCookieBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestCookieBehavior",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

