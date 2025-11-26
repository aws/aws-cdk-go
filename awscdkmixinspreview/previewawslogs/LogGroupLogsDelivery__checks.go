//go:build !no_runtime_type_checking

package previewawslogs

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (l *jsiiProxy_LogGroupLogsDelivery) validateBindParameters(scope constructs.IConstruct, deliverySource interfacesawslogs.IDeliverySourceRef, sourceResourceArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if deliverySource == nil {
		return fmt.Errorf("parameter deliverySource is required, but nil was provided")
	}

	if sourceResourceArn == nil {
		return fmt.Errorf("parameter sourceResourceArn is required, but nil was provided")
	}

	return nil
}

func validateNewLogGroupLogsDeliveryParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	if logGroup == nil {
		return fmt.Errorf("parameter logGroup is required, but nil was provided")
	}

	return nil
}

