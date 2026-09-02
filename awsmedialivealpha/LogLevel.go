package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The log level for the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   logLevel := medialive_alpha.LogLevel_DEBUG()
//
// Experimental.
type LogLevel interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for LogLevel
type jsiiProxy_LogLevel struct {
	_ byte // padding
}

func (j *jsiiProxy_LogLevel) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func LogLevel_Of(value *string) LogLevel {
	_init_.Initialize()

	if err := validateLogLevel_OfParameters(value); err != nil {
		panic(err)
	}
	var returns LogLevel

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.LogLevel",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func LogLevel_DEBUG() LogLevel {
	_init_.Initialize()
	var returns LogLevel
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.LogLevel",
		"DEBUG",
		&returns,
	)
	return returns
}

func LogLevel_DISABLED() LogLevel {
	_init_.Initialize()
	var returns LogLevel
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.LogLevel",
		"DISABLED",
		&returns,
	)
	return returns
}

func LogLevel_ERROR() LogLevel {
	_init_.Initialize()
	var returns LogLevel
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.LogLevel",
		"ERROR",
		&returns,
	)
	return returns
}

func LogLevel_INFO() LogLevel {
	_init_.Initialize()
	var returns LogLevel
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.LogLevel",
		"INFO",
		&returns,
	)
	return returns
}

func LogLevel_WARNING() LogLevel {
	_init_.Initialize()
	var returns LogLevel
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.LogLevel",
		"WARNING",
		&returns,
	)
	return returns
}

