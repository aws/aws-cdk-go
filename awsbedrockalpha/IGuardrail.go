package awsbedrockalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Represents a Guardrail, either created with CDK or imported.
// Experimental.
type IGuardrail interface {
	awscdk.IResource
	// Grant the given principal identity permissions to perform actions on this guardrail.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to apply the guardrail.
	// Experimental.
	GrantApply(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this guardrail.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the invocation client errors metric for this guardrail.
	// Experimental.
	MetricInvocationClientErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the invocation latency metric for this guardrail.
	// Experimental.
	MetricInvocationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the invocations metric for this guardrail.
	// Experimental.
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the invocation server errors metric for this guardrail.
	// Experimental.
	MetricInvocationServerErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the invocations intervened metric for this guardrail.
	// Experimental.
	MetricInvocationsIntervened(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the invocation throttles metric for this guardrail.
	// Experimental.
	MetricInvocationThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the text unit count metric for this guardrail.
	// Experimental.
	MetricTextUnitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ARN of the guardrail.
	// Experimental.
	GuardrailArn() *string
	// The ID of the guardrail.
	// Experimental.
	GuardrailId() *string
	// The version of the guardrail.
	//
	// If no explicit version is created,
	// this will default to "DRAFT".
	// Experimental.
	GuardrailVersion() *string
	// Optional KMS encryption key associated with this guardrail.
	// Experimental.
	KmsKey() awskms.IKey
	// When this guardrail was last updated.
	// Experimental.
	LastUpdated() *string
}

// The jsii proxy for IGuardrail
type jsiiProxy_IGuardrail struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IGuardrail) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGuardrail) GrantApply(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantApplyParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantApply",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGuardrail) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGuardrail) MetricInvocationClientErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInvocationClientErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInvocationClientErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGuardrail) MetricInvocationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInvocationLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInvocationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGuardrail) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGuardrail) MetricInvocationServerErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInvocationServerErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInvocationServerErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGuardrail) MetricInvocationsIntervened(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInvocationsIntervenedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInvocationsIntervened",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGuardrail) MetricInvocationThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInvocationThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInvocationThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGuardrail) MetricTextUnitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTextUnitCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTextUnitCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IGuardrail) GuardrailArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guardrailArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGuardrail) GuardrailId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guardrailId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGuardrail) GuardrailVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guardrailVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGuardrail) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGuardrail) LastUpdated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdated",
		&returns,
	)
	return returns
}

