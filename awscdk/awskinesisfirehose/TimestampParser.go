package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Value class that wraps a Joda Time format string.
//
// Use this with the Hive JSON input format for data record format conversion to parse custom timestamp formats.
//
// Example:
//   inputFormat := firehose.NewHiveJsonInputFormat(&HiveJsonInputFormatProps{
//   	TimestampParsers: []TimestampParser{
//   		firehose.TimestampParser_FromFormatString(jsii.String("yyyy-MM-dd")),
//   		firehose.TimestampParser_EPOCH_MILLIS(),
//   	},
//   })
//
type TimestampParser interface {
	// The format string to use in Hive JSON input format configuration.
	Format() *string
}

// The jsii proxy struct for TimestampParser
type jsiiProxy_TimestampParser struct {
	_ byte // padding
}

func (j *jsiiProxy_TimestampParser) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}


// Creates a TimestampParser from the given format string.
//
// The format string should be a valid Joda Time pattern string.
// See [Class DateTimeFormat](https://docs.aws.amazon.com/https://www.joda.org/joda-time/apidocs/org/joda/time/format/DateTimeFormat.html) for more details
func TimestampParser_FromFormatString(format *string) TimestampParser {
	_init_.Initialize()

	if err := validateTimestampParser_FromFormatStringParameters(format); err != nil {
		panic(err)
	}
	var returns TimestampParser

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.TimestampParser",
		"fromFormatString",
		[]interface{}{format},
		&returns,
	)

	return returns
}

func TimestampParser_EPOCH_MILLIS() TimestampParser {
	_init_.Initialize()
	var returns TimestampParser
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.TimestampParser",
		"EPOCH_MILLIS",
		&returns,
	)
	return returns
}

