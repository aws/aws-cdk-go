package previewawselasticacheevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.elasticache@SnapshotExportFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   snapshotExportFailed := awscdkmixinspreview.Events.NewSnapshotExportFailed()
//
// Experimental.
type SnapshotExportFailed interface {
}

// The jsii proxy struct for SnapshotExportFailed
type jsiiProxy_SnapshotExportFailed struct {
	_ byte // padding
}

// Experimental.
func NewSnapshotExportFailed() SnapshotExportFailed {
	_init_.Initialize()

	j := jsiiProxy_SnapshotExportFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.SnapshotExportFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSnapshotExportFailed_Override(s SnapshotExportFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.events.SnapshotExportFailed",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Snapshot Export Failed.
// Experimental.
func SnapshotExportFailed_EventPattern(options *SnapshotExportFailed_SnapshotExportFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSnapshotExportFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticache.events.SnapshotExportFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

