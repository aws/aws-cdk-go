//go:build !no_runtime_type_checking

package previewawsec2events

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
)

func (i *jsiiProxy_InstanceEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_InstanceEvents) validateEc2InstanceStateChangeNotificationPatternParameters(options *EC2InstanceStateChangeNotification_EC2InstanceStateChangeNotificationProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_InstanceEvents) validateEc2SpotInstanceInterruptionWarningPatternParameters(options *EC2SpotInstanceInterruptionWarning_EC2SpotInstanceInterruptionWarningProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInstanceEvents_FromInstanceParameters(instanceRef interfacesawsec2.IInstanceRef) error {
	if instanceRef == nil {
		return fmt.Errorf("parameter instanceRef is required, but nil was provided")
	}

	return nil
}

