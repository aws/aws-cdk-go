package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@ReadSetExportJobStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   readSetExportJobStatusChange := awscdkmixinspreview.Events.NewReadSetExportJobStatusChange()
//
// Experimental.
type ReadSetExportJobStatusChange interface {
}

// The jsii proxy struct for ReadSetExportJobStatusChange
type jsiiProxy_ReadSetExportJobStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewReadSetExportJobStatusChange() ReadSetExportJobStatusChange {
	_init_.Initialize()

	j := jsiiProxy_ReadSetExportJobStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReadSetExportJobStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewReadSetExportJobStatusChange_Override(r ReadSetExportJobStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReadSetExportJobStatusChange",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for Read Set Export Job Status Change.
// Experimental.
func ReadSetExportJobStatusChange_ReadSetExportJobStatusChangePattern(options *ReadSetExportJobStatusChange_ReadSetExportJobStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateReadSetExportJobStatusChange_ReadSetExportJobStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.ReadSetExportJobStatusChange",
		"readSetExportJobStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

