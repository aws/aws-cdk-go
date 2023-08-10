//go:build !no_runtime_type_checking

package awscdkappconfigalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

func validateNewSqsDestinationParameters(queue awssqs.IQueue) error {
	if queue == nil {
		return fmt.Errorf("parameter queue is required, but nil was provided")
	}

	return nil
}

