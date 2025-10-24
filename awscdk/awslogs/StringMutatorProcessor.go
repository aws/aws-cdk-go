package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Processor for string mutation operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringMutatorProcessor := awscdk.Aws_logs.NewStringMutatorProcessor(&StringMutatorProps{
//   	Type: awscdk.*Aws_logs.StringMutatorType_LOWER_CASE,
//
//   	// the properties below are optional
//   	LowerCaseKeys: []*string{
//   		jsii.String("lowerCaseKeys"),
//   	},
//   	SplitOptions: &SplitStringProperty{
//   		Entries: []SplitStringEntryProperty{
//   			&SplitStringEntryProperty{
//   				Delimiter: awscdk.*Aws_logs.DelimiterCharacter_COMMA,
//   				Source: jsii.String("source"),
//   			},
//   		},
//   	},
//   	SubstituteOptions: &SubstituteStringProperty{
//   		Entries: []SubstituteStringEntryProperty{
//   			&SubstituteStringEntryProperty{
//   				From: jsii.String("from"),
//   				Source: jsii.String("source"),
//   				To: jsii.String("to"),
//   			},
//   		},
//   	},
//   	TrimKeys: []*string{
//   		jsii.String("trimKeys"),
//   	},
//   	UpperCaseKeys: []*string{
//   		jsii.String("upperCaseKeys"),
//   	},
//   })
//
type StringMutatorProcessor interface {
	IProcessor
	// The type of string mutation operation.
	Type() StringMutatorType
	SetType(val StringMutatorType)
}

// The jsii proxy struct for StringMutatorProcessor
type jsiiProxy_StringMutatorProcessor struct {
	jsiiProxy_IProcessor
}

func (j *jsiiProxy_StringMutatorProcessor) Type() StringMutatorType {
	var returns StringMutatorType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates a new string mutator processor.
func NewStringMutatorProcessor(props *StringMutatorProps) StringMutatorProcessor {
	_init_.Initialize()

	if err := validateNewStringMutatorProcessorParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StringMutatorProcessor{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.StringMutatorProcessor",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Creates a new string mutator processor.
func NewStringMutatorProcessor_Override(s StringMutatorProcessor, props *StringMutatorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.StringMutatorProcessor",
		[]interface{}{props},
		s,
	)
}

func (j *jsiiProxy_StringMutatorProcessor)SetType(val StringMutatorType) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

