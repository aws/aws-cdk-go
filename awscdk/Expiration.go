package awscdk

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a date of expiration.
//
// The amount can be specified either as a Date object, timestamp, Duration or string.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   expiration := cdk.Expiration_After(cdk.Duration_Minutes(jsii.Number(30)))
//
type Expiration interface {
	// Expiration value as a Date object.
	Date() *time.Time
	// Check if Expiration expires after input.
	IsAfter(t Duration) *bool
	// Check if Expiration expires before input.
	IsBefore(t Duration) *bool
	// Expiration Value in a formatted Unix Epoch Time in seconds.
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
func Expiration_After(t Duration) Expiration {
	_init_.Initialize()

	if err := validateExpiration_AfterParameters(t); err != nil {
		panic(err)
	}
	var returns Expiration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Expiration",
		"after",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Expire at the specified date.
func Expiration_AtDate(d *time.Time) Expiration {
	_init_.Initialize()

	if err := validateExpiration_AtDateParameters(d); err != nil {
		panic(err)
	}
	var returns Expiration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Expiration",
		"atDate",
		[]interface{}{d},
		&returns,
	)

	return returns
}

// Expire at the specified timestamp.
func Expiration_AtTimestamp(t *float64) Expiration {
	_init_.Initialize()

	if err := validateExpiration_AtTimestampParameters(t); err != nil {
		panic(err)
	}
	var returns Expiration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Expiration",
		"atTimestamp",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Expire at specified date, represented as a string.
func Expiration_FromString(s *string) Expiration {
	_init_.Initialize()

	if err := validateExpiration_FromStringParameters(s); err != nil {
		panic(err)
	}
	var returns Expiration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Expiration",
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

