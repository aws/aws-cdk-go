package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

type IFunction interface {
	awsec2.IConnectable
	awsiam.IGrantable
	awscdk.IResource
	// Adds an event source to this function.
	//
	// Event sources are implemented in the aws-cdk-lib/aws-lambda-event-sources module.
	//
	// The following example adds an SQS Queue as an event source:
	// ```
	// import { SqsEventSource } from 'aws-cdk-lib/aws-lambda-event-sources';
	// myFunction.addEventSource(new SqsEventSource(myQueue));
	// ```.
	AddEventSource(source IEventSource)
	// Adds an event source that maps to this AWS Lambda function.
	AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping
	// Adds a url to this lambda function.
	AddFunctionUrl(options *FunctionUrlOptions) FunctionUrl
	// Adds a permission to the Lambda resource policy.
	// See: Permission for details.
	//
	AddPermission(id *string, permission *Permission)
	// Adds a statement to the IAM role assumed by the instance.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Configures options for asynchronous invocation.
	ConfigureAsyncInvoke(options *EventInvokeConfigOptions)
	// Grant the given identity permissions to invoke this Lambda.
	GrantInvoke(identity awsiam.IGrantable) awsiam.Grant
	// Grant multiple principals the ability to invoke this Lambda via CompositePrincipal.
	GrantInvokeCompositePrincipal(compositePrincipal awsiam.CompositePrincipal) *[]awsiam.Grant
	// Grant the given identity permissions to invoke this Lambda Function URL.
	GrantInvokeUrl(identity awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Lambda Return the given named metric for this Function.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the Duration of this Lambda How long execution of this Lambda takes.
	//
	// Average over 5 minutes.
	// Default: average over 5 minutes.
	//
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How many invocations of this Lambda fail.
	//
	// Sum over 5 minutes.
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of invocations of this Lambda How often this Lambda is invoked.
	//
	// Sum over 5 minutes.
	// Default: sum over 5 minutes.
	//
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of throttled invocations of this Lambda How often this Lambda is throttled.
	//
	// Sum over 5 minutes.
	// Default: sum over 5 minutes.
	//
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The system architectures compatible with this lambda function.
	Architecture() Architecture
	// The ARN of the function.
	FunctionArn() *string
	// The name of the function.
	FunctionName() *string
	// Whether or not this Lambda function was bound to a VPC.
	//
	// If this is is `false`, trying to access the `connections` object will fail.
	IsBoundToVpc() *bool
	// The `$LATEST` version of this function.
	//
	// Note that this is reference to a non-specific AWS Lambda version, which
	// means the function this version refers to can return different results in
	// different invocations.
	//
	// To obtain a reference to an explicit version which references the current
	// function configuration, use `lambdaFunction.currentVersion` instead.
	LatestVersion() IVersion
	// The construct node where permissions are attached.
	PermissionsNode() constructs.Node
	// The ARN(s) to put into the resource field of the generated IAM policy for grantInvoke().
	//
	// This property is for cdk modules to consume only. You should not need to use this property.
	// Instead, use grantInvoke() directly.
	ResourceArnsForGrantInvoke() *[]*string
	// The IAM role associated with this function.
	Role() awsiam.IRole
}

// The jsii proxy for IFunction
type jsiiProxy_IFunction struct {
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IFunction) AddEventSource(source IEventSource) {
	if err := i.validateAddEventSourceParameters(source); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addEventSource",
		[]interface{}{source},
	)
}

func (i *jsiiProxy_IFunction) AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping {
	if err := i.validateAddEventSourceMappingParameters(id, options); err != nil {
		panic(err)
	}
	var returns EventSourceMapping

	_jsii_.Invoke(
		i,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFunction) AddFunctionUrl(options *FunctionUrlOptions) FunctionUrl {
	if err := i.validateAddFunctionUrlParameters(options); err != nil {
		panic(err)
	}
	var returns FunctionUrl

	_jsii_.Invoke(
		i,
		"addFunctionUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFunction) AddPermission(id *string, permission *Permission) {
	if err := i.validateAddPermissionParameters(id, permission); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addPermission",
		[]interface{}{id, permission},
	)
}

func (i *jsiiProxy_IFunction) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := i.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (i *jsiiProxy_IFunction) ConfigureAsyncInvoke(options *EventInvokeConfigOptions) {
	if err := i.validateConfigureAsyncInvokeParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_IFunction) GrantInvoke(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantInvokeParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvoke",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFunction) GrantInvokeCompositePrincipal(compositePrincipal awsiam.CompositePrincipal) *[]awsiam.Grant {
	if err := i.validateGrantInvokeCompositePrincipalParameters(compositePrincipal); err != nil {
		panic(err)
	}
	var returns *[]awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvokeCompositePrincipal",
		[]interface{}{compositePrincipal},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFunction) GrantInvokeUrl(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantInvokeUrlParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvokeUrl",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFunction) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IFunction) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IFunction) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IFunction) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IFunction) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IFunction) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) LatestVersion() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) ResourceArnsForGrantInvoke() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnsForGrantInvoke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

