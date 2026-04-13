package previewawsbatchevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.batch@BatchJobStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   batchJobStateChange := awscdkmixinspreview.Events.NewBatchJobStateChange()
//
// Experimental.
type BatchJobStateChange interface {
}

// The jsii proxy struct for BatchJobStateChange
type jsiiProxy_BatchJobStateChange struct {
	_ byte // padding
}

// Experimental.
func NewBatchJobStateChange() BatchJobStateChange {
	_init_.Initialize()

	j := jsiiProxy_BatchJobStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_batch.events.BatchJobStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewBatchJobStateChange_Override(b BatchJobStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_batch.events.BatchJobStateChange",
		nil, // no parameters
		b,
	)
}

// EventBridge event pattern for Batch Job State Change.
// Experimental.
func BatchJobStateChange_EventPattern(options *BatchJobStateChange_BatchJobStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateBatchJobStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_batch.events.BatchJobStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

