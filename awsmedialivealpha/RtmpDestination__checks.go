//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateRtmpDestination_UrlParameters(url *string, streamName *string, options *OutputDestinationOptions) error {
	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	if streamName == nil {
		return fmt.Errorf("parameter streamName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

