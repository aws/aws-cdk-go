package previewawsec2events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ec2@EBSMultiVolumeSnapshotsCompletionStatus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eBSMultiVolumeSnapshotsCompletionStatus := awscdkmixinspreview.Events.NewEBSMultiVolumeSnapshotsCompletionStatus()
//
// Experimental.
type EBSMultiVolumeSnapshotsCompletionStatus interface {
}

// The jsii proxy struct for EBSMultiVolumeSnapshotsCompletionStatus
type jsiiProxy_EBSMultiVolumeSnapshotsCompletionStatus struct {
	_ byte // padding
}

// Experimental.
func NewEBSMultiVolumeSnapshotsCompletionStatus() EBSMultiVolumeSnapshotsCompletionStatus {
	_init_.Initialize()

	j := jsiiProxy_EBSMultiVolumeSnapshotsCompletionStatus{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.EBSMultiVolumeSnapshotsCompletionStatus",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEBSMultiVolumeSnapshotsCompletionStatus_Override(e EBSMultiVolumeSnapshotsCompletionStatus) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.EBSMultiVolumeSnapshotsCompletionStatus",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EBS Multi-Volume Snapshots Completion Status.
// Experimental.
func EBSMultiVolumeSnapshotsCompletionStatus_EbsMultiVolumeSnapshotsCompletionStatusPattern(options *EBSMultiVolumeSnapshotsCompletionStatus_EBSMultiVolumeSnapshotsCompletionStatusProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEBSMultiVolumeSnapshotsCompletionStatus_EbsMultiVolumeSnapshotsCompletionStatusPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.events.EBSMultiVolumeSnapshotsCompletionStatus",
		"ebsMultiVolumeSnapshotsCompletionStatusPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

