//go:build !no_runtime_type_checking

package awscdkgluealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (j *jsiiProxy_Job) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Job) validateBuildJobArnParameters(scope constructs.Construct, jobName *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if jobName == nil {
		return fmt.Errorf("parameter jobName is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Job) validateCodeS3ObjectUrlParameters(code Code) error {
	if code == nil {
		return fmt.Errorf("parameter code is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Job) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_Job) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Job) validateMetricParameters(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) error {
	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	if type_ == "" {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_Job) validateMetricFailureParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_Job) validateMetricSuccessParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_Job) validateMetricTimeoutParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_Job) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_Job) validateOnFailureParameters(id *string, options *awsevents.OnEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_Job) validateOnStateChangeParameters(id *string, jobState JobState, options *awsevents.OnEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if jobState == "" {
		return fmt.Errorf("parameter jobState is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_Job) validateOnSuccessParameters(id *string, options *awsevents.OnEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_Job) validateOnTimeoutParameters(id *string, options *awsevents.OnEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_Job) validateSetupContinuousLoggingParameters(role awsiam.IRole, props *ContinuousLoggingProps) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateJob_FromJobAttributesParameters(scope constructs.Construct, id *string, attrs *JobAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validateJob_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateJob_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateJob_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewJobParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

