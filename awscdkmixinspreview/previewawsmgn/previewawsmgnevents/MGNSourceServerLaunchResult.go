package previewawsmgnevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.mgn@MGNSourceServerLaunchResult.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mGNSourceServerLaunchResult := awscdkmixinspreview.Events.NewMGNSourceServerLaunchResult()
//
// Experimental.
type MGNSourceServerLaunchResult interface {
}

// The jsii proxy struct for MGNSourceServerLaunchResult
type jsiiProxy_MGNSourceServerLaunchResult struct {
	_ byte // padding
}

// Experimental.
func NewMGNSourceServerLaunchResult() MGNSourceServerLaunchResult {
	_init_.Initialize()

	j := jsiiProxy_MGNSourceServerLaunchResult{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mgn.events.MGNSourceServerLaunchResult",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMGNSourceServerLaunchResult_Override(m MGNSourceServerLaunchResult) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mgn.events.MGNSourceServerLaunchResult",
		nil, // no parameters
		m,
	)
}

// EventBridge event pattern for MGN Source Server Launch Result.
// Experimental.
func MGNSourceServerLaunchResult_EventPattern(options *MGNSourceServerLaunchResult_MGNSourceServerLaunchResultProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateMGNSourceServerLaunchResult_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mgn.events.MGNSourceServerLaunchResult",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

