package previewawsmedialiveevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
)

// EventBridge event patterns for Channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var channelRef IChannelRef
//
//   channelEvents := awscdkmixinspreview.Events.ChannelEvents_FromChannel(channelRef)
//
// Experimental.
type ChannelEvents interface {
	// EventBridge event pattern for Channel MediaLive Channel Input Change.
	// Experimental.
	MediaLiveChannelInputChangePattern(options *ChannelEvents_MediaLiveChannelInputChange_MediaLiveChannelInputChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for ChannelEvents
type jsiiProxy_ChannelEvents struct {
	_ byte // padding
}

// Create ChannelEvents from a Channel reference.
// Experimental.
func ChannelEvents_FromChannel(channelRef interfacesawsmedialive.IChannelRef) ChannelEvents {
	_init_.Initialize()

	if err := validateChannelEvents_FromChannelParameters(channelRef); err != nil {
		panic(err)
	}
	var returns ChannelEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_medialive.events.ChannelEvents",
		"fromChannel",
		[]interface{}{channelRef},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChannelEvents) MediaLiveChannelInputChangePattern(options *ChannelEvents_MediaLiveChannelInputChange_MediaLiveChannelInputChangeProps) *awsevents.EventPattern {
	if err := c.validateMediaLiveChannelInputChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"mediaLiveChannelInputChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

