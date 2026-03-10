package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@ReadSetImportJobStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   readSetImportJobStatusChange := awscdkmixinspreview.Events.NewReadSetImportJobStatusChange()
//
// Experimental.
type ReadSetImportJobStatusChange interface {
}

// The jsii proxy struct for ReadSetImportJobStatusChange
type jsiiProxy_ReadSetImportJobStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewReadSetImportJobStatusChange() ReadSetImportJobStatusChange {
	_init_.Initialize()

	j := jsiiProxy_ReadSetImportJobStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReadSetImportJobStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewReadSetImportJobStatusChange_Override(r ReadSetImportJobStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReadSetImportJobStatusChange",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for Read Set Import Job Status Change.
// Experimental.
func ReadSetImportJobStatusChange_ReadSetImportJobStatusChangePattern(options *ReadSetImportJobStatusChange_ReadSetImportJobStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateReadSetImportJobStatusChange_ReadSetImportJobStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.ReadSetImportJobStatusChange",
		"readSetImportJobStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

