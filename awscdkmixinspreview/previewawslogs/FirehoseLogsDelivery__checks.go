//go:build !no_runtime_type_checking

package previewawslogs

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
)

func (f *jsiiProxy_FirehoseLogsDelivery) validateBindParameters(scope constructs.IConstruct, logType *string, sourceResourceArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if logType == nil {
		return fmt.Errorf("parameter logType is required, but nil was provided")
	}

	if sourceResourceArn == nil {
		return fmt.Errorf("parameter sourceResourceArn is required, but nil was provided")
	}

	return nil
}

func validateNewFirehoseLogsDeliveryParameters(stream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	if stream == nil {
		return fmt.Errorf("parameter stream is required, but nil was provided")
	}

	return nil
}

