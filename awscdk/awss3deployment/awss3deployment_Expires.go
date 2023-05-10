package awss3deployment

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Used for HTTP expires header, which influences downstream caches.
//
// Does NOT influence deletion of the object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   expires := awscdk.Aws_s3_deployment.Expires_After(duration)
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
//
// Deprecated: use core.Expiration
type Expires interface {
	// The raw expiration date expression.
	// Deprecated: use core.Expiration
	Value() interface{}
}

// The jsii proxy struct for Expires
type jsiiProxy_Expires struct {
	_ byte // padding
}

func (j *jsiiProxy_Expires) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Expire once the specified duration has passed since deployment time.
// Deprecated: use core.Expiration
func Expires_After(t awscdk.Duration) Expires {
	_init_.Initialize()

	if err := validateExpires_AfterParameters(t); err != nil {
		panic(err)
	}
	var returns Expires

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Expires",
		"after",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Expire at the specified date.
// Deprecated: use core.Expiration
func Expires_AtDate(d *time.Time) Expires {
	_init_.Initialize()

	if err := validateExpires_AtDateParameters(d); err != nil {
		panic(err)
	}
	var returns Expires

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Expires",
		"atDate",
		[]interface{}{d},
		&returns,
	)

	return returns
}

// Expire at the specified timestamp.
// Deprecated: use core.Expiration
func Expires_AtTimestamp(t *float64) Expires {
	_init_.Initialize()

	if err := validateExpires_AtTimestampParameters(t); err != nil {
		panic(err)
	}
	var returns Expires

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Expires",
		"atTimestamp",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Create an expiration date from a raw date string.
// Deprecated: use core.Expiration
func Expires_FromString(s *string) Expires {
	_init_.Initialize()

	if err := validateExpires_FromStringParameters(s); err != nil {
		panic(err)
	}
	var returns Expires

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Expires",
		"fromString",
		[]interface{}{s},
		&returns,
	)

	return returns
}

