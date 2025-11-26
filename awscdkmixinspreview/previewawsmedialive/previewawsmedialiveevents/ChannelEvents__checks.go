//go:build !no_runtime_type_checking

package previewawsmedialiveevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
)

func (c *jsiiProxy_ChannelEvents) validateMediaLiveChannelInputChangePatternParameters(options *ChannelEvents_MediaLiveChannelInputChange_MediaLiveChannelInputChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateChannelEvents_FromChannelParameters(channelRef interfacesawsmedialive.IChannelRef) error {
	if channelRef == nil {
		return fmt.Errorf("parameter channelRef is required, but nil was provided")
	}

	return nil
}

