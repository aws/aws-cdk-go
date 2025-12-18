//go:build !no_runtime_type_checking

package previewawslogs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_S3LogsDelivery) validateBindParameters(scope constructs.IConstruct, logType *string, sourceResourceArn *string) error {
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

func validateNewS3LogsDeliveryParameters(bucket interfacesawss3.IBucketRef, props *S3LogsDeliveryProps) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

