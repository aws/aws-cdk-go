package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@RunStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   runStatusChange := awscdkmixinspreview.Events.NewRunStatusChange()
//
// Experimental.
type RunStatusChange interface {
}

// The jsii proxy struct for RunStatusChange
type jsiiProxy_RunStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewRunStatusChange() RunStatusChange {
	_init_.Initialize()

	j := jsiiProxy_RunStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.RunStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRunStatusChange_Override(r RunStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.RunStatusChange",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for Run Status Change.
// Experimental.
func RunStatusChange_RunStatusChangePattern(options *RunStatusChange_RunStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateRunStatusChange_RunStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.RunStatusChange",
		"runStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

