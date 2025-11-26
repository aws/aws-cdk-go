//go:build !no_runtime_type_checking

package previewawsorganizationsevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsorganizations"
)

func (a *jsiiProxy_AccountEvents) validateAwsServiceEventViaCloudTrailPatternParameters(options *AccountEvents_AWSServiceEventViaCloudTrail_AWSServiceEventViaCloudTrailProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateAccountEvents_FromAccountParameters(accountRef interfacesawsorganizations.IAccountRef) error {
	if accountRef == nil {
		return fmt.Errorf("parameter accountRef is required, but nil was provided")
	}

	return nil
}

