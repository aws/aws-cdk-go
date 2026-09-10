//go:build !no_runtime_type_checking

package awscdkgluealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue"
)

func validateCondition_CrawlerParameters(crawler interfacesawsglue.ICrawlerRef, crawlState CrawlerState, options *ConditionOptions) error {
	if crawler == nil {
		return fmt.Errorf("parameter crawler is required, but nil was provided")
	}

	if crawlState == "" {
		return fmt.Errorf("parameter crawlState is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateCondition_JobParameters(job interfacesawsglue.IJobRef, state JobState, options *ConditionOptions) error {
	if job == nil {
		return fmt.Errorf("parameter job is required, but nil was provided")
	}

	if state == "" {
		return fmt.Errorf("parameter state is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

