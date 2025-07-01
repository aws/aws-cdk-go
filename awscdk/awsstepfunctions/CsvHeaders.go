package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for CSV header options for a CSV Item Reader.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   csvHeaders := awscdk.Aws_stepfunctions.CsvHeaders_Use([]*string{
//   	jsii.String("headers"),
//   })
//
type CsvHeaders interface {
	// Location of headers in CSV file.
	HeaderLocation() CsvHeaderLocation
	// List of headers if `headerLocation` is `GIVEN`.
	Headers() *[]*string
}

// The jsii proxy struct for CsvHeaders
type jsiiProxy_CsvHeaders struct {
	_ byte // padding
}

func (j *jsiiProxy_CsvHeaders) HeaderLocation() CsvHeaderLocation {
	var returns CsvHeaderLocation
	_jsii_.Get(
		j,
		"headerLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsvHeaders) Headers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}


// Configures S3CsvItemReader to use the headers provided in the `headers` parameter.
//
// Returns: - CsvHeaders.
func CsvHeaders_Use(headers *[]*string) CsvHeaders {
	_init_.Initialize()

	if err := validateCsvHeaders_UseParameters(headers); err != nil {
		panic(err)
	}
	var returns CsvHeaders

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.CsvHeaders",
		"use",
		[]interface{}{headers},
		&returns,
	)

	return returns
}

// Configures S3CsvItemReader to read headers from the first row of the CSV file.
//
// Returns: - CsvHeaders.
func CsvHeaders_UseFirstRow() CsvHeaders {
	_init_.Initialize()

	var returns CsvHeaders

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.CsvHeaders",
		"useFirstRow",
		nil, // no parameters
		&returns,
	)

	return returns
}

