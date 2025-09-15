package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Parser processor for AWS vended logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vendedLogParser := awscdk.Aws_logs.NewVendedLogParser(&VendedLogParserProps{
//   	LogType: awscdk.*Aws_logs.VendedLogType_CLOUDFRONT,
//
//   	// the properties below are optional
//   	Source: jsii.String("source"),
//   })
//
type VendedLogParser interface {
	IProcessor
	// The type of AWS vended log.
	LogType() VendedLogType
	SetLogType(val VendedLogType)
}

// The jsii proxy struct for VendedLogParser
type jsiiProxy_VendedLogParser struct {
	jsiiProxy_IProcessor
}

func (j *jsiiProxy_VendedLogParser) LogType() VendedLogType {
	var returns VendedLogType
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Creates a new vended log parser processor.
func NewVendedLogParser(props *VendedLogParserProps) VendedLogParser {
	_init_.Initialize()

	if err := validateNewVendedLogParserParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_VendedLogParser{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.VendedLogParser",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Creates a new vended log parser processor.
func NewVendedLogParser_Override(v VendedLogParser, props *VendedLogParserProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.VendedLogParser",
		[]interface{}{props},
		v,
	)
}

func (j *jsiiProxy_VendedLogParser)SetLogType(val VendedLogType) {
	if err := j.validateSetLogTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logType",
		val,
	)
}

