package previewawselasticacheevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.elasticache@SnapshotCreated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   snapshotCreated := awscdkmixinspreview.Events.NewSnapshotCreated()
//
// Experimental.
type SnapshotCreated interface {
}

// The jsii proxy struct for SnapshotCreated
type jsiiProxy_SnapshotCreated struct {
	_ byte // padding
}

// Experimental.
func NewSnapshotCreated() SnapshotCreated {
	_init_.Initialize()

	j := jsiiProxy_SnapshotCreated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.SnapshotCreated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSnapshotCreated_Override(s SnapshotCreated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.SnapshotCreated",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Snapshot Created.
// Experimental.
func SnapshotCreated_EventPattern(options *SnapshotCreated_SnapshotCreatedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSnapshotCreated_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticache.events.SnapshotCreated",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

