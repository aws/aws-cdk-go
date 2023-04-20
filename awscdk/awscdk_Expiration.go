// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a date of expiration.
//
// The amount can be specified either as a Date object, timestamp, Duration or string.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   expiration := monocdk.Expiration_After(duration)
//
// Experimental.
type Expiration interface {
	// Expiration value as a Date object.
	// Experimental.
	Date() *time.Time
	// Check if Exipiration expires after input.
	// Experimental.
	IsAfter(t Duration) *bool
	// Check if Exipiration expires before input.
	// Experimental.
	IsBefore(t Duration) *bool
	// Exipration Value in a formatted Unix Epoch Time in seconds.
	// Experimental.
	ToEpoch() *float64
}

// The jsii proxy struct for Expiration
type jsiiProxy_Expiration struct {
	_ byte // padding
}

func (j *jsiiProxy_Expiration) Date() *time.Time {
	var returns *time.Time
	_jsii_.Get(
		j,
		"date",
		&returns,
	)
	return returns
}


// Expire once the specified duration has passed since deployment time.
// Experimental.
func Expiration_After(t Duration) Expiration {
	_init_.Initialize()

	if err := validateExpiration_AfterParameters(t); err != nil {
		panic(err)
	}
	var returns Expiration

	_jsii_.StaticInvoke(
		"monocdk.Expiration",
		"after",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Expire at the specified date.
// Experimental.
func Expiration_AtDate(d *time.Time) Expiration {
	_init_.Initialize()

	if err := validateExpiration_AtDateParameters(d); err != nil {
		panic(err)
	}
	var returns Expiration

	_jsii_.StaticInvoke(
		"monocdk.Expiration",
		"atDate",
		[]interface{}{d},
		&returns,
	)

	return returns
}

// Expire at the specified timestamp.
// Experimental.
func Expiration_AtTimestamp(t *float64) Expiration {
	_init_.Initialize()

	if err := validateExpiration_AtTimestampParameters(t); err != nil {
		panic(err)
	}
	var returns Expiration

	_jsii_.StaticInvoke(
		"monocdk.Expiration",
		"atTimestamp",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Expire at specified date, represented as a string.
// Experimental.
func Expiration_FromString(s *string) Expiration {
	_init_.Initialize()

	if err := validateExpiration_FromStringParameters(s); err != nil {
		panic(err)
	}
	var returns Expiration

	_jsii_.StaticInvoke(
		"monocdk.Expiration",
		"fromString",
		[]interface{}{s},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Expiration) IsAfter(t Duration) *bool {
	if err := e.validateIsAfterParameters(t); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		e,
		"isAfter",
		[]interface{}{t},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Expiration) IsBefore(t Duration) *bool {
	if err := e.validateIsBeforeParameters(t); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		e,
		"isBefore",
		[]interface{}{t},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Expiration) ToEpoch() *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"toEpoch",
		nil, // no parameters
		&returns,
	)

	return returns
}

