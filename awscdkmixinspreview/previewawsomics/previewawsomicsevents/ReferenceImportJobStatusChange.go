package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@ReferenceImportJobStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   referenceImportJobStatusChange := awscdkmixinspreview.Events.NewReferenceImportJobStatusChange()
//
// Experimental.
type ReferenceImportJobStatusChange interface {
}

// The jsii proxy struct for ReferenceImportJobStatusChange
type jsiiProxy_ReferenceImportJobStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewReferenceImportJobStatusChange() ReferenceImportJobStatusChange {
	_init_.Initialize()

	j := jsiiProxy_ReferenceImportJobStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceImportJobStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewReferenceImportJobStatusChange_Override(r ReferenceImportJobStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceImportJobStatusChange",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for Reference Import Job Status Change.
// Experimental.
func ReferenceImportJobStatusChange_EventPattern(options *ReferenceImportJobStatusChange_ReferenceImportJobStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateReferenceImportJobStatusChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceImportJobStatusChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

