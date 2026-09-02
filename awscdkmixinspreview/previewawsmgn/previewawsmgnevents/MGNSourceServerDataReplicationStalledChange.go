package previewawsmgnevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.mgn@MGNSourceServerDataReplicationStalledChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mGNSourceServerDataReplicationStalledChange := awscdkmixinspreview.Events.NewMGNSourceServerDataReplicationStalledChange()
//
// Experimental.
type MGNSourceServerDataReplicationStalledChange interface {
}

// The jsii proxy struct for MGNSourceServerDataReplicationStalledChange
type jsiiProxy_MGNSourceServerDataReplicationStalledChange struct {
	_ byte // padding
}

// Experimental.
func NewMGNSourceServerDataReplicationStalledChange() MGNSourceServerDataReplicationStalledChange {
	_init_.Initialize()

	j := jsiiProxy_MGNSourceServerDataReplicationStalledChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mgn.events.MGNSourceServerDataReplicationStalledChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMGNSourceServerDataReplicationStalledChange_Override(m MGNSourceServerDataReplicationStalledChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mgn.events.MGNSourceServerDataReplicationStalledChange",
		nil, // no parameters
		m,
	)
}

// EventBridge event pattern for MGN Source Server Data Replication Stalled Change.
// Experimental.
func MGNSourceServerDataReplicationStalledChange_EventPattern(options *MGNSourceServerDataReplicationStalledChange_MGNSourceServerDataReplicationStalledChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateMGNSourceServerDataReplicationStalledChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mgn.events.MGNSourceServerDataReplicationStalledChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

