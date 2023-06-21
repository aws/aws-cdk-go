package awss3deployment

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Used for HTTP cache-control header, which influences downstream caches.
//
// Example:
//   var destinationBucket bucket
//
//   s3deploy.NewBucketDeployment(this, jsii.String("BucketDeployment"), &BucketDeploymentProps{
//   	Sources: []iSource{
//   		s3deploy.Source_Asset(jsii.String("./website"), &AssetOptions{
//   			Exclude: []*string{
//   				jsii.String("index.html"),
//   			},
//   		}),
//   	},
//   	DestinationBucket: DestinationBucket,
//   	CacheControl: []cacheControl{
//   		s3deploy.*cacheControl_MaxAge(awscdk.Duration_Days(jsii.Number(365))),
//   		s3deploy.*cacheControl_Immutable(),
//   	},
//   	Prune: jsii.Boolean(false),
//   })
//
//   s3deploy.NewBucketDeployment(this, jsii.String("HTMLBucketDeployment"), &BucketDeploymentProps{
//   	Sources: []*iSource{
//   		s3deploy.Source_*Asset(jsii.String("./website"), &AssetOptions{
//   			Exclude: []*string{
//   				jsii.String("*"),
//   				jsii.String("!index.html"),
//   			},
//   		}),
//   	},
//   	DestinationBucket: DestinationBucket,
//   	CacheControl: []*cacheControl{
//   		s3deploy.*cacheControl_*MaxAge(awscdk.Duration_Seconds(jsii.Number(0))),
//   	},
//   	Prune: jsii.Boolean(false),
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
//
type CacheControl interface {
	// The raw cache control setting.
	Value() interface{}
}

// The jsii proxy struct for CacheControl
type jsiiProxy_CacheControl struct {
	_ byte // padding
}

func (j *jsiiProxy_CacheControl) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Constructs a custom cache control key from the literal value.
func CacheControl_FromString(s *string) CacheControl {
	_init_.Initialize()

	if err := validateCacheControl_FromStringParameters(s); err != nil {
		panic(err)
	}
	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"fromString",
		[]interface{}{s},
		&returns,
	)

	return returns
}

// Sets 'immutable'.
func CacheControl_Immutable() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"immutable",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets 'max-age=<duration-in-seconds>'.
func CacheControl_MaxAge(t awscdk.Duration) CacheControl {
	_init_.Initialize()

	if err := validateCacheControl_MaxAgeParameters(t); err != nil {
		panic(err)
	}
	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"maxAge",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Sets 'must-revalidate'.
func CacheControl_MustRevalidate() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"mustRevalidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets 'must-understand'.
func CacheControl_MustUnderstand() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"mustUnderstand",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets 'no-cache'.
func CacheControl_NoCache() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"noCache",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets 'no-store'.
func CacheControl_NoStore() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"noStore",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets 'no-transform'.
func CacheControl_NoTransform() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"noTransform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets 'proxy-revalidate'.
func CacheControl_ProxyRevalidate() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"proxyRevalidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets 'private'.
func CacheControl_SetPrivate() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"setPrivate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets 'public'.
func CacheControl_SetPublic() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"setPublic",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets 's-maxage=<duration-in-seconds>'.
func CacheControl_SMaxAge(t awscdk.Duration) CacheControl {
	_init_.Initialize()

	if err := validateCacheControl_SMaxAgeParameters(t); err != nil {
		panic(err)
	}
	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"sMaxAge",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Sets 'stale-if-error=<duration-in-seconds>'.
func CacheControl_StaleIfError(t awscdk.Duration) CacheControl {
	_init_.Initialize()

	if err := validateCacheControl_StaleIfErrorParameters(t); err != nil {
		panic(err)
	}
	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"staleIfError",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Sets 'stale-while-revalidate=<duration-in-seconds>'.
func CacheControl_StaleWhileRevalidate(t awscdk.Duration) CacheControl {
	_init_.Initialize()

	if err := validateCacheControl_StaleWhileRevalidateParameters(t); err != nil {
		panic(err)
	}
	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"staleWhileRevalidate",
		[]interface{}{t},
		&returns,
	)

	return returns
}

