package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Failover Mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   failoverMode := mediaconnect_alpha.FailoverMode_Of(jsii.String("value"))
//
// Experimental.
type FailoverMode interface {
	// The failover mode string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FailoverMode
type jsiiProxy_FailoverMode struct {
	_ byte // padding
}

func (j *jsiiProxy_FailoverMode) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom failover mode value.
// Experimental.
func FailoverMode_Of(value *string) FailoverMode {
	_init_.Initialize()

	if err := validateFailoverMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns FailoverMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.FailoverMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func FailoverMode_FAILOVER() FailoverMode {
	_init_.Initialize()
	var returns FailoverMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.FailoverMode",
		"FAILOVER",
		&returns,
	)
	return returns
}

func FailoverMode_MERGE() FailoverMode {
	_init_.Initialize()
	var returns FailoverMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.FailoverMode",
		"MERGE",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FailoverMode) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

