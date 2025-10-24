package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Processor for data conversion operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataConverterProcessor := awscdk.Aws_logs.NewDataConverterProcessor(&DataConverterProps{
//   	Type: awscdk.*Aws_logs.DataConverterType_TYPE_CONVERTER,
//
//   	// the properties below are optional
//   	DateTimeConverterOptions: &DateTimeConverterProperty{
//   		Locale: jsii.String("locale"),
//   		MatchPatterns: []*string{
//   			jsii.String("matchPatterns"),
//   		},
//   		Source: jsii.String("source"),
//   		Target: jsii.String("target"),
//
//   		// the properties below are optional
//   		SourceTimezone: jsii.String("sourceTimezone"),
//   		TargetFormat: jsii.String("targetFormat"),
//   		TargetTimezone: jsii.String("targetTimezone"),
//   	},
//   	TypeConverterOptions: &TypeConverterProperty{
//   		Entries: []TypeConverterEntryProperty{
//   			&TypeConverterEntryProperty{
//   				Key: jsii.String("key"),
//   				Type: awscdk.*Aws_logs.TypeConverterType_BOOLEAN,
//   			},
//   		},
//   	},
//   })
//
type DataConverterProcessor interface {
	IProcessor
	// The type of data conversion operation.
	Type() DataConverterType
	SetType(val DataConverterType)
}

// The jsii proxy struct for DataConverterProcessor
type jsiiProxy_DataConverterProcessor struct {
	jsiiProxy_IProcessor
}

func (j *jsiiProxy_DataConverterProcessor) Type() DataConverterType {
	var returns DataConverterType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates a new data converter processor.
func NewDataConverterProcessor(props *DataConverterProps) DataConverterProcessor {
	_init_.Initialize()

	if err := validateNewDataConverterProcessorParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataConverterProcessor{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.DataConverterProcessor",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Creates a new data converter processor.
func NewDataConverterProcessor_Override(d DataConverterProcessor, props *DataConverterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.DataConverterProcessor",
		[]interface{}{props},
		d,
	)
}

func (j *jsiiProxy_DataConverterProcessor)SetType(val DataConverterType) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

