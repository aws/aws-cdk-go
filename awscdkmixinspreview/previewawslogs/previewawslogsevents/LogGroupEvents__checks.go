//go:build !no_runtime_type_checking

package previewawslogsevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

func (l *jsiiProxy_LogGroupEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *LogGroupEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateLogGroupEvents_FromLogGroupParameters(logGroupRef interfacesawslogs.ILogGroupRef) error {
	if logGroupRef == nil {
		return fmt.Errorf("parameter logGroupRef is required, but nil was provided")
	}

	return nil
}

