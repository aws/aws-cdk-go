//go:build !no_runtime_type_checking

package awscdkappconfigalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

func validateNewSnsDestinationParameters(topic awssns.ITopic) error {
	if topic == nil {
		return fmt.Errorf("parameter topic is required, but nil was provided")
	}

	return nil
}

