//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"
)

func validateLinkedChannelSettings_FollowerParameters(primaryChannel IChannel) error {
	if primaryChannel == nil {
		return fmt.Errorf("parameter primaryChannel is required, but nil was provided")
	}

	return nil
}

