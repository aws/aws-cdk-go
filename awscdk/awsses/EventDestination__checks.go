//go:build !no_runtime_type_checking

package awsses

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

func validateEventDestination_CloudWatchDimensionsParameters(dimensions *[]*CloudWatchDimension) error {
	if dimensions == nil {
		return fmt.Errorf("parameter dimensions is required, but nil was provided")
	}
	for idx_70b27d, v := range *dimensions {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter dimensions[%#v]", idx_70b27d) }); err != nil {
			return err
		}
	}

	return nil
}

func validateEventDestination_EventBusParameters(eventBus awsevents.IEventBus) error {
	if eventBus == nil {
		return fmt.Errorf("parameter eventBus is required, but nil was provided")
	}

	return nil
}

func validateEventDestination_FirehoseDeliveryStreamParameters(stream *FirehoseDeliveryStreamDestination) error {
	if stream == nil {
		return fmt.Errorf("parameter stream is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(stream, func() string { return "parameter stream" }); err != nil {
		return err
	}

	return nil
}

func validateEventDestination_SnsTopicParameters(topic awssns.ITopic) error {
	if topic == nil {
		return fmt.Errorf("parameter topic is required, but nil was provided")
	}

	return nil
}

