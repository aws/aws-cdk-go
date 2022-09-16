package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Used for HTTP cache-control header, which influences downstream caches.
//
// Use the provided static factory methods to construct instances of this class.
// Used in the {@link S3DeployActionProps.cacheControl} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cacheControl := awscdk.Aws_codepipeline_actions.cacheControl.fromString(jsii.String("s"))
//
// See: https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9
//
type CacheControl interface {
	// the actual text value of the created directive.
	Value() *string
	SetValue(val *string)
}

// The jsii proxy struct for CacheControl
type jsiiProxy_CacheControl struct {
	_ byte // padding
}

func (j *jsiiProxy_CacheControl) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func (j *jsiiProxy_CacheControl)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Allows you to create an arbitrary cache control directive, in case our support is missing a method for a particular directive.
func CacheControl_FromString(s *string) CacheControl {
	_init_.Initialize()

	if err := validateCacheControl_FromStringParameters(s); err != nil {
		panic(err)
	}
	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"fromString",
		[]interface{}{s},
		&returns,
	)

	return returns
}

// The 'max-age' cache control directive.
func CacheControl_MaxAge(t awscdk.Duration) CacheControl {
	_init_.Initialize()

	if err := validateCacheControl_MaxAgeParameters(t); err != nil {
		panic(err)
	}
	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"maxAge",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// The 'must-revalidate' cache control directive.
func CacheControl_MustRevalidate() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"mustRevalidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The 'no-cache' cache control directive.
func CacheControl_NoCache() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"noCache",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The 'no-transform' cache control directive.
func CacheControl_NoTransform() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"noTransform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The 'proxy-revalidate' cache control directive.
func CacheControl_ProxyRevalidate() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"proxyRevalidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The 'private' cache control directive.
func CacheControl_SetPrivate() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"setPrivate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The 'public' cache control directive.
func CacheControl_SetPublic() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"setPublic",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The 's-max-age' cache control directive.
func CacheControl_SMaxAge(t awscdk.Duration) CacheControl {
	_init_.Initialize()

	if err := validateCacheControl_SMaxAgeParameters(t); err != nil {
		panic(err)
	}
	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.CacheControl",
		"sMaxAge",
		[]interface{}{t},
		&returns,
	)

	return returns
}

