package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@ReadSetStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   readSetStatusChange := awscdkmixinspreview.Events.NewReadSetStatusChange()
//
// Experimental.
type ReadSetStatusChange interface {
}

// The jsii proxy struct for ReadSetStatusChange
type jsiiProxy_ReadSetStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewReadSetStatusChange() ReadSetStatusChange {
	_init_.Initialize()

	j := jsiiProxy_ReadSetStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReadSetStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewReadSetStatusChange_Override(r ReadSetStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReadSetStatusChange",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for Read Set Status Change.
// Experimental.
func ReadSetStatusChange_ReadSetStatusChangePattern(options *ReadSetStatusChange_ReadSetStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateReadSetStatusChange_ReadSetStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.ReadSetStatusChange",
		"readSetStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

