package previewawselasticacheevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.elasticache@SnapshotCopyFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   snapshotCopyFailed := awscdkmixinspreview.Events.NewSnapshotCopyFailed()
//
// Experimental.
type SnapshotCopyFailed interface {
}

// The jsii proxy struct for SnapshotCopyFailed
type jsiiProxy_SnapshotCopyFailed struct {
	_ byte // padding
}

// Experimental.
func NewSnapshotCopyFailed() SnapshotCopyFailed {
	_init_.Initialize()

	j := jsiiProxy_SnapshotCopyFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.SnapshotCopyFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSnapshotCopyFailed_Override(s SnapshotCopyFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.SnapshotCopyFailed",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Snapshot Copy Failed.
// Experimental.
func SnapshotCopyFailed_EventPattern(options *SnapshotCopyFailed_SnapshotCopyFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSnapshotCopyFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticache.events.SnapshotCopyFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

