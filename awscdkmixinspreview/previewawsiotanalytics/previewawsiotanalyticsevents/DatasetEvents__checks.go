//go:build !no_runtime_type_checking

package previewawsiotanalyticsevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotanalytics"
)

func (d *jsiiProxy_DatasetEvents) validateIoTAnalyticsDataSetLifeCycleNotificationPatternParameters(options *DatasetEvents_IoTAnalyticsDataSetLifeCycleNotification_IoTAnalyticsDataSetLifeCycleNotificationProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDatasetEvents_FromDatasetParameters(datasetRef interfacesawsiotanalytics.IDatasetRef) error {
	if datasetRef == nil {
		return fmt.Errorf("parameter datasetRef is required, but nil was provided")
	}

	return nil
}

