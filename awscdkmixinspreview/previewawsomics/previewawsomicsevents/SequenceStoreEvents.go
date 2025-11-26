package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsomics"
)

// EventBridge event patterns for SequenceStore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sequenceStoreRef ISequenceStoreRef
//
//   sequenceStoreEvents := awscdkmixinspreview.Events.SequenceStoreEvents_FromSequenceStore(sequenceStoreRef)
//
// Experimental.
type SequenceStoreEvents interface {
	// EventBridge event pattern for SequenceStore Read Set Activation Job Status Change.
	// Experimental.
	ReadSetActivationJobStatusChangePattern(options *SequenceStoreEvents_ReadSetActivationJobStatusChange_ReadSetActivationJobStatusChangeProps) *awsevents.EventPattern
	// EventBridge event pattern for SequenceStore Read Set Export Job Status Change.
	// Experimental.
	ReadSetExportJobStatusChangePattern(options *SequenceStoreEvents_ReadSetExportJobStatusChange_ReadSetExportJobStatusChangeProps) *awsevents.EventPattern
	// EventBridge event pattern for SequenceStore Read Set Import Job Status Change.
	// Experimental.
	ReadSetImportJobStatusChangePattern(options *SequenceStoreEvents_ReadSetImportJobStatusChange_ReadSetImportJobStatusChangeProps) *awsevents.EventPattern
	// EventBridge event pattern for SequenceStore Read Set Status Change.
	// Experimental.
	ReadSetStatusChangePattern(options *SequenceStoreEvents_ReadSetStatusChange_ReadSetStatusChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for SequenceStoreEvents
type jsiiProxy_SequenceStoreEvents struct {
	_ byte // padding
}

// Create SequenceStoreEvents from a SequenceStore reference.
// Experimental.
func SequenceStoreEvents_FromSequenceStore(sequenceStoreRef interfacesawsomics.ISequenceStoreRef) SequenceStoreEvents {
	_init_.Initialize()

	if err := validateSequenceStoreEvents_FromSequenceStoreParameters(sequenceStoreRef); err != nil {
		panic(err)
	}
	var returns SequenceStoreEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents",
		"fromSequenceStore",
		[]interface{}{sequenceStoreRef},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SequenceStoreEvents) ReadSetActivationJobStatusChangePattern(options *SequenceStoreEvents_ReadSetActivationJobStatusChange_ReadSetActivationJobStatusChangeProps) *awsevents.EventPattern {
	if err := s.validateReadSetActivationJobStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		s,
		"readSetActivationJobStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SequenceStoreEvents) ReadSetExportJobStatusChangePattern(options *SequenceStoreEvents_ReadSetExportJobStatusChange_ReadSetExportJobStatusChangeProps) *awsevents.EventPattern {
	if err := s.validateReadSetExportJobStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		s,
		"readSetExportJobStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SequenceStoreEvents) ReadSetImportJobStatusChangePattern(options *SequenceStoreEvents_ReadSetImportJobStatusChange_ReadSetImportJobStatusChangeProps) *awsevents.EventPattern {
	if err := s.validateReadSetImportJobStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		s,
		"readSetImportJobStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SequenceStoreEvents) ReadSetStatusChangePattern(options *SequenceStoreEvents_ReadSetStatusChange_ReadSetStatusChangeProps) *awsevents.EventPattern {
	if err := s.validateReadSetStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		s,
		"readSetStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

