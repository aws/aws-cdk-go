package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@AnnotationImportJobStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   annotationImportJobStatusChange := awscdkmixinspreview.Events.NewAnnotationImportJobStatusChange()
//
// Experimental.
type AnnotationImportJobStatusChange interface {
}

// The jsii proxy struct for AnnotationImportJobStatusChange
type jsiiProxy_AnnotationImportJobStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewAnnotationImportJobStatusChange() AnnotationImportJobStatusChange {
	_init_.Initialize()

	j := jsiiProxy_AnnotationImportJobStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.AnnotationImportJobStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAnnotationImportJobStatusChange_Override(a AnnotationImportJobStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.AnnotationImportJobStatusChange",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for Annotation Import Job Status Change.
// Experimental.
func AnnotationImportJobStatusChange_EventPattern(options *AnnotationImportJobStatusChange_AnnotationImportJobStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAnnotationImportJobStatusChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.AnnotationImportJobStatusChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

