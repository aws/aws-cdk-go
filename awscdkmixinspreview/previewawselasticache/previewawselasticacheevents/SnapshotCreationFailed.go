package previewawselasticacheevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.elasticache@SnapshotCreationFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   snapshotCreationFailed := awscdkmixinspreview.Events.NewSnapshotCreationFailed()
//
// Experimental.
type SnapshotCreationFailed interface {
}

// The jsii proxy struct for SnapshotCreationFailed
type jsiiProxy_SnapshotCreationFailed struct {
	_ byte // padding
}

// Experimental.
func NewSnapshotCreationFailed() SnapshotCreationFailed {
	_init_.Initialize()

	j := jsiiProxy_SnapshotCreationFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.SnapshotCreationFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSnapshotCreationFailed_Override(s SnapshotCreationFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.SnapshotCreationFailed",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Snapshot Creation Failed.
// Experimental.
func SnapshotCreationFailed_EventPattern(options *SnapshotCreationFailed_SnapshotCreationFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSnapshotCreationFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticache.events.SnapshotCreationFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

