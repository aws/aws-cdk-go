package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore"
	"github.com/aws/constructs-go/constructs/v10"
)

// Contains all properties and methods for both created and imported policy engines.
// Experimental.
type IPolicyEngine interface {
	awsiam.IGrantable
	interfacesawsbedrockagentcore.IPolicyEngineRef
	awscdk.IResource
	// Grants IAM actions to the IAM Principal.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grants permissions to evaluate policies at runtime .
	//
	// This is the primary permission needed by Gateway execution roles to evaluate
	// authorization decisions during agent requests. Grant this to roles that need
	// to call AuthorizeAction or PartiallyAuthorizeActions.
	// Experimental.
	GrantEvaluate(grantee awsiam.IGrantable) awsiam.Grant
	// Grants the full set of permissions required for a gateway execution role to use this policy engine, correctly scoped to both the policy engine and gateway ARNs.
	//
	// Per the AWS docs, `AuthorizeAction` and `PartiallyAuthorizeActions` require
	// both the policy engine ARN and the gateway ARN as resources, while
	// `GetPolicyEngine` only needs the policy engine ARN.
	//
	// This follows the same pattern as Lambda's `grantInvokeVersion(grantee, version)`.
	// Experimental.
	GrantEvaluateForGateway(grantee awsiam.IGrantable, gateway IGateway) awsiam.Grant
	// Grants read permissions on the PolicyEngine.
	//
	// This grants runtime read access to policy engine configuration. Use this for
	// monitoring, observability, or read-only administrative roles.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this policy engine.
	// Experimental.
	Metric(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric measuring the authorization latency for this policy engine.
	//
	// This metric represents the time taken to evaluate authorization policies.
	// Experimental.
	MetricAuthorizationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of denied authorization requests for this policy engine.
	//
	// This metric tracks authorization requests that were explicitly denied by policies.
	// Experimental.
	MetricDeniedRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of errors during authorization for this policy engine.
	//
	// This metric tracks errors encountered during policy evaluation.
	// Experimental.
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The description of the policy engine.
	// Experimental.
	Description() *string
	// The KMS key used for encryption.
	// Experimental.
	KmsKey() awskms.IKey
	// The ARN of the policy engine resource.
	// Experimental.
	PolicyEngineArn() *string
	// The ID of the policy engine.
	// Experimental.
	PolicyEngineId() *string
	// The name of the policy engine.
	// Experimental.
	PolicyEngineName() *string
}

// The jsii proxy for IPolicyEngine
type jsiiProxy_IPolicyEngine struct {
	internal.Type__awsiamIGrantable
	internal.Type__interfacesawsbedrockagentcoreIPolicyEngineRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IPolicyEngine) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IPolicyEngine) GrantEvaluate(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantEvaluateParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantEvaluate",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPolicyEngine) GrantEvaluateForGateway(grantee awsiam.IGrantable, gateway IGateway) awsiam.Grant {
	if err := i.validateGrantEvaluateForGatewayParameters(grantee, gateway); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantEvaluateForGateway",
		[]interface{}{grantee, gateway},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPolicyEngine) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IPolicyEngine) Metric(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IPolicyEngine) MetricAuthorizationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricAuthorizationLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricAuthorizationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPolicyEngine) MetricDeniedRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDeniedRequestsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDeniedRequests",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPolicyEngine) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPolicyEngine) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IPolicyEngine) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPolicyEngine) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyEngine) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyEngine) PolicyEngineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyEngineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyEngine) PolicyEngineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyEngineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyEngine) PolicyEngineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyEngineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyEngine) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyEngine) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyEngine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyEngine) PolicyEngineRef() *interfacesawsbedrockagentcore.PolicyEngineReference {
	var returns *interfacesawsbedrockagentcore.PolicyEngineReference
	_jsii_.Get(
		j,
		"policyEngineRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyEngine) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

