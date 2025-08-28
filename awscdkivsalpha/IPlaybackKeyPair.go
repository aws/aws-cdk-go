package awscdkivsalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkivsalpha/v2/internal"
)

// Represents an IVS Playback Key Pair.
// Experimental.
type IPlaybackKeyPair interface {
	awscdk.IResource
	// Key-pair ARN.
	//
	// For example: arn:aws:ivs:us-west-2:693991300569:playback-key/f99cde61-c2b0-4df3-8941-ca7d38acca1a.
	// Experimental.
	PlaybackKeyPairArn() *string
}

// The jsii proxy for IPlaybackKeyPair
type jsiiProxy_IPlaybackKeyPair struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IPlaybackKeyPair) PlaybackKeyPairArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playbackKeyPairArn",
		&returns,
	)
	return returns
}

