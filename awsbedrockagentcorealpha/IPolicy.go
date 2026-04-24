package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore"
	"github.com/aws/constructs-go/constructs/v10"
)

// Minimal reference interface for Policy resources.
//
// Used for resource identification and ARN construction.
// Experimental.
type IPolicy interface {
	awsiam.IGrantable
	interfacesawsbedrockagentcore.IPolicyRef
	awscdk.IResource
	// Grants IAM actions to the IAM Principal.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grants read permissions on the Policy (data plane).
	//
	// This grants runtime read access to policy configuration. Use this for monitoring,
	// audit, or read-only administrative roles that need to inspect policy definitions.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this policy.
	// Experimental.
	Metric(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric measuring the evaluation latency for this policy.
	//
	// This metric represents the time taken to evaluate this specific policy.
	// Experimental.
	MetricEvaluationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the total number of evaluations for this policy.
	//
	// This metric tracks how many times this policy has been evaluated.
	// Experimental.
	MetricEvaluations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The description of the policy.
	// Experimental.
	Description() *string
	// The ARN of the policy resource.
	// Experimental.
	PolicyArn() *string
	// The policy engine this policy belongs to.
	// Experimental.
	PolicyEngine() IPolicyEngine
	// The ID of the policy.
	// Experimental.
	PolicyId() *string
	// The name of the policy.
	// Experimental.
	PolicyName() *string
	// The validation mode for the policy.
	// Experimental.
	ValidationMode() PolicyValidationMode
}

// The jsii proxy for IPolicy
type jsiiProxy_IPolicy struct {
	internal.Type__awsiamIGrantable
	internal.Type__interfacesawsbedrockagentcoreIPolicyRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IPolicy) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IPolicy) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPolicy) Metric(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, dimensions, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, dimensions, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPolicy) MetricEvaluationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricEvaluationLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEvaluationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPolicy) MetricEvaluations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricEvaluationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEvaluations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IPolicy) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IPolicy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicy) PolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicy) PolicyEngine() IPolicyEngine {
	var returns IPolicyEngine
	_jsii_.Get(
		j,
		"policyEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicy) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicy) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicy) ValidationMode() PolicyValidationMode {
	var returns PolicyValidationMode
	_jsii_.Get(
		j,
		"validationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicy) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicy) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicy) PolicyRef() *interfacesawsbedrockagentcore.PolicyReference {
	var returns *interfacesawsbedrockagentcore.PolicyReference
	_jsii_.Get(
		j,
		"policyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

