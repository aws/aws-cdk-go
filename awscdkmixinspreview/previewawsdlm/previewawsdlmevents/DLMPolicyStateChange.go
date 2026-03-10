package previewawsdlmevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.dlm@DLMPolicyStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dLMPolicyStateChange := awscdkmixinspreview.Events.NewDLMPolicyStateChange()
//
// Experimental.
type DLMPolicyStateChange interface {
}

// The jsii proxy struct for DLMPolicyStateChange
type jsiiProxy_DLMPolicyStateChange struct {
	_ byte // padding
}

// Experimental.
func NewDLMPolicyStateChange() DLMPolicyStateChange {
	_init_.Initialize()

	j := jsiiProxy_DLMPolicyStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dlm.events.DLMPolicyStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDLMPolicyStateChange_Override(d DLMPolicyStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dlm.events.DLMPolicyStateChange",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for DLM Policy State Change.
// Experimental.
func DLMPolicyStateChange_DlmPolicyStateChangePattern(options *DLMPolicyStateChange_DLMPolicyStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDLMPolicyStateChange_DlmPolicyStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_dlm.events.DLMPolicyStateChange",
		"dlmPolicyStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

