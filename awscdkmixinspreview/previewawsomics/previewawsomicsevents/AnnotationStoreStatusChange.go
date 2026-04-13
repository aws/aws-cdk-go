package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@AnnotationStoreStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   annotationStoreStatusChange := awscdkmixinspreview.Events.NewAnnotationStoreStatusChange()
//
// Experimental.
type AnnotationStoreStatusChange interface {
}

// The jsii proxy struct for AnnotationStoreStatusChange
type jsiiProxy_AnnotationStoreStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewAnnotationStoreStatusChange() AnnotationStoreStatusChange {
	_init_.Initialize()

	j := jsiiProxy_AnnotationStoreStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.AnnotationStoreStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAnnotationStoreStatusChange_Override(a AnnotationStoreStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.AnnotationStoreStatusChange",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for Annotation Store Status Change.
// Experimental.
func AnnotationStoreStatusChange_EventPattern(options *AnnotationStoreStatusChange_AnnotationStoreStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAnnotationStoreStatusChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.AnnotationStoreStatusChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

