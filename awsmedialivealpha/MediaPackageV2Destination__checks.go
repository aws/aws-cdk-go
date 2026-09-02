//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2"
)

func validateMediaPackageV2Destination_ChannelParameters(channel awsmediapackagev2alpha.IChannel) error {
	if channel == nil {
		return fmt.Errorf("parameter channel is required, but nil was provided")
	}

	return nil
}

