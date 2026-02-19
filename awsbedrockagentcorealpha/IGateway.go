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

// Interface for Gateway resources.
// Experimental.
type IGateway interface {
	interfacesawsbedrockagentcore.IGatewayRef
	awscdk.IResource
	// Grants IAM actions to the IAM Principal.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grants permission to invoke this Gateway.
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Grants `Create`, `Update`, and `Delete` actions on the Gateway.
	// Experimental.
	GrantManage(grantee awsiam.IGrantable) awsiam.Grant
	// Grants `Get` and `List` actions on the Gateway.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this gateway.
	// Experimental.
	Metric(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric measuring the duration of requests for this gateway.
	//
	// The duration metric represents the total time elapsed between receiving the request
	// and sending the final response token, representing complete end-to-end processing time.
	// Default: - Average statistic over 5 minutes.
	//
	// Experimental.
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the total number of invocations for this gateway.
	//
	// This metric tracks all successful invocations of the gateway.
	// Default: - Sum statistic over 5 minutes.
	//
	// Experimental.
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric measuring the latency of requests for this gateway.
	//
	// The latency metric represents the time elapsed between when the service receives
	// the request and when it begins sending the first response token.
	// Default: - Average statistic over 5 minutes.
	//
	// Experimental.
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of system errors (5xx status code) for this gateway.
	//
	// This metric tracks internal server errors and system failures.
	// Default: - Sum statistic over 5 minutes.
	//
	// Experimental.
	MetricSystemErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric measuring the target execution time for this gateway.
	//
	// This metric helps determine the contribution of the target (Lambda, OpenAPI, etc.)
	// to the total latency.
	// Default: - Average statistic over 5 minutes.
	//
	// Experimental.
	MetricTargetExecutionTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of requests served by each target type for this gateway.
	// Default: - Sum statistic over 5 minutes.
	//
	// Experimental.
	MetricTargetType(targetType *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of throttled requests (429 status code) for this gateway.
	//
	// This metric helps identify when the gateway is rate limiting requests.
	// Default: - Sum statistic over 5 minutes.
	//
	// Experimental.
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of user errors (4xx status code, excluding 429) for this gateway.
	//
	// This metric tracks client errors like bad requests, unauthorized access, etc.
	// Default: - Sum statistic over 5 minutes.
	//
	// Experimental.
	MetricUserErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The authorizer configuration for the gateway.
	// Experimental.
	AuthorizerConfiguration() IGatewayAuthorizerConfig
	// Timestamp when the gateway was created.
	// Experimental.
	CreatedAt() *string
	// The description of the gateway.
	// Experimental.
	Description() *string
	// The exception level for the gateway.
	// Experimental.
	ExceptionLevel() GatewayExceptionLevel
	// The ARN of the gateway resource.
	// Experimental.
	GatewayArn() *string
	// The id of the gateway.
	// Experimental.
	GatewayId() *string
	// The URL endpoint for the gateway.
	// Experimental.
	GatewayUrl() *string
	// The KMS key used for encryption.
	// Experimental.
	KmsKey() awskms.IKey
	// The name of the gateway.
	// Experimental.
	Name() *string
	// The protocol configuration for the gateway.
	// Experimental.
	ProtocolConfiguration() IGatewayProtocolConfig
	// The IAM role that provides permissions for the gateway to access AWS services.
	// Experimental.
	Role() awsiam.IRole
	// The status of the gateway.
	// Experimental.
	Status() *string
	// The status reasons for the gateway.
	// Experimental.
	StatusReason() *[]*string
	// Timestamp when the gateway was last updated.
	// Experimental.
	UpdatedAt() *string
}

// The jsii proxy for IGateway
type jsiiProxy_IGateway struct {
	internal.Type__interfacesawsbedrockagentcoreIGatewayRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IGateway) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IGateway) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGateway) GrantManage(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantManageParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantManage",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGateway) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IGateway) Metric(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IGateway) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGateway) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IGateway) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGateway) MetricSystemErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSystemErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSystemErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGateway) MetricTargetExecutionTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTargetExecutionTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTargetExecutionTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGateway) MetricTargetType(targetType *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTargetTypeParameters(targetType, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTargetType",
		[]interface{}{targetType, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGateway) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGateway) MetricUserErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricUserErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricUserErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGateway) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IGateway) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IGateway) AuthorizerConfiguration() IGatewayAuthorizerConfig {
	var returns IGatewayAuthorizerConfig
	_jsii_.Get(
		j,
		"authorizerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) ExceptionLevel() GatewayExceptionLevel {
	var returns GatewayExceptionLevel
	_jsii_.Get(
		j,
		"exceptionLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) GatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) GatewayUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) ProtocolConfiguration() IGatewayProtocolConfig {
	var returns IGatewayProtocolConfig
	_jsii_.Get(
		j,
		"protocolConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) StatusReason() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) GatewayRef() *interfacesawsbedrockagentcore.GatewayReference {
	var returns *interfacesawsbedrockagentcore.GatewayReference
	_jsii_.Get(
		j,
		"gatewayRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

