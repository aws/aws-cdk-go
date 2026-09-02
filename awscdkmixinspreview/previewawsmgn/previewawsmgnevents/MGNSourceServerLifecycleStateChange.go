package previewawsmgnevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.mgn@MGNSourceServerLifecycleStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mGNSourceServerLifecycleStateChange := awscdkmixinspreview.Events.NewMGNSourceServerLifecycleStateChange()
//
// Experimental.
type MGNSourceServerLifecycleStateChange interface {
}

// The jsii proxy struct for MGNSourceServerLifecycleStateChange
type jsiiProxy_MGNSourceServerLifecycleStateChange struct {
	_ byte // padding
}

// Experimental.
func NewMGNSourceServerLifecycleStateChange() MGNSourceServerLifecycleStateChange {
	_init_.Initialize()

	j := jsiiProxy_MGNSourceServerLifecycleStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mgn.events.MGNSourceServerLifecycleStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMGNSourceServerLifecycleStateChange_Override(m MGNSourceServerLifecycleStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mgn.events.MGNSourceServerLifecycleStateChange",
		nil, // no parameters
		m,
	)
}

// EventBridge event pattern for MGN Source Server Lifecycle State Change.
// Experimental.
func MGNSourceServerLifecycleStateChange_EventPattern(options *MGNSourceServerLifecycleStateChange_MGNSourceServerLifecycleStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateMGNSourceServerLifecycleStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mgn.events.MGNSourceServerLifecycleStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

