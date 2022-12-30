package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The HTTP methods that the Behavior will accept requests on.
//
// Example:
//   // Create a Distribution with configured HTTP methods and viewer protocol policy of the cache.
//   var myBucket bucket
//
//   myWebDistribution := cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewS3Origin(myBucket),
//   		allowedMethods: cloudfront.allowedMethods_ALLOW_ALL(),
//   		viewerProtocolPolicy: cloudfront.viewerProtocolPolicy_REDIRECT_TO_HTTPS,
//   	},
//   })
//
// Experimental.
type AllowedMethods interface {
	// HTTP methods supported.
	// Experimental.
	Methods() *[]*string
}

// The jsii proxy struct for AllowedMethods
type jsiiProxy_AllowedMethods struct {
	_ byte // padding
}

func (j *jsiiProxy_AllowedMethods) Methods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"methods",
		&returns,
	)
	return returns
}


func AllowedMethods_ALLOW_ALL() AllowedMethods {
	_init_.Initialize()
	var returns AllowedMethods
	_jsii_.StaticGet(
		"monocdk.aws_cloudfront.AllowedMethods",
		"ALLOW_ALL",
		&returns,
	)
	return returns
}

func AllowedMethods_ALLOW_GET_HEAD() AllowedMethods {
	_init_.Initialize()
	var returns AllowedMethods
	_jsii_.StaticGet(
		"monocdk.aws_cloudfront.AllowedMethods",
		"ALLOW_GET_HEAD",
		&returns,
	)
	return returns
}

func AllowedMethods_ALLOW_GET_HEAD_OPTIONS() AllowedMethods {
	_init_.Initialize()
	var returns AllowedMethods
	_jsii_.StaticGet(
		"monocdk.aws_cloudfront.AllowedMethods",
		"ALLOW_GET_HEAD_OPTIONS",
		&returns,
	)
	return returns
}

