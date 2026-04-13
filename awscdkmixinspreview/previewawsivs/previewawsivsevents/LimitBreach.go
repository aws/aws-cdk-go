package previewawsivsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ivs@LimitBreach.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   limitBreach := awscdkmixinspreview.Events.NewLimitBreach()
//
// Experimental.
type LimitBreach interface {
}

// The jsii proxy struct for LimitBreach
type jsiiProxy_LimitBreach struct {
	_ byte // padding
}

// Experimental.
func NewLimitBreach() LimitBreach {
	_init_.Initialize()

	j := jsiiProxy_LimitBreach{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.events.LimitBreach",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewLimitBreach_Override(l LimitBreach) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.events.LimitBreach",
		nil, // no parameters
		l,
	)
}

// EventBridge event pattern for IVS Limit Breach.
// Experimental.
func LimitBreach_EventPattern(options *LimitBreach_LimitBreachProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateLimitBreach_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ivs.events.LimitBreach",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

