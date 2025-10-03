package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

type IVersion interface {
	IFunction
	IVersionRef
	// Defines an alias for this version.
	// Deprecated: Calling `addAlias` on a `Version` object will cause the Alias to be replaced on every function update. Call `function.addAlias()` or `new Alias()` instead.
	AddAlias(aliasName *string, options *AliasOptions) Alias
	// The ARN of the version for Lambda@Edge.
	EdgeArn() *string
	// The underlying AWS Lambda function.
	Lambda() IFunction
	// The most recently deployed version of this function.
	Version() *string
}

// The jsii proxy for IVersion
type jsiiProxy_IVersion struct {
	jsiiProxy_IFunction
	jsiiProxy_IVersionRef
}

func (i *jsiiProxy_IVersion) AddAlias(aliasName *string, options *AliasOptions) Alias {
	if err := i.validateAddAliasParameters(aliasName, options); err != nil {
		panic(err)
	}
	var returns Alias

	_jsii_.Invoke(
		i,
		"addAlias",
		[]interface{}{aliasName, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVersion) AddEventSource(source IEventSource) {
	if err := i.validateAddEventSourceParameters(source); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addEventSource",
		[]interface{}{source},
	)
}

func (i *jsiiProxy_IVersion) AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping {
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

func (i *jsiiProxy_IVersion) AddFunctionUrl(options *FunctionUrlOptions) FunctionUrl {
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

func (i *jsiiProxy_IVersion) AddPermission(id *string, permission *Permission) {
	if err := i.validateAddPermissionParameters(id, permission); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addPermission",
		[]interface{}{id, permission},
	)
}

func (i *jsiiProxy_IVersion) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := i.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (i *jsiiProxy_IVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IVersion) ConfigureAsyncInvoke(options *EventInvokeConfigOptions) {
	if err := i.validateConfigureAsyncInvokeParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_IVersion) GrantInvoke(identity awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IVersion) GrantInvokeCompositePrincipal(compositePrincipal awsiam.CompositePrincipal) *[]awsiam.Grant {
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

func (i *jsiiProxy_IVersion) GrantInvokeLatestVersion(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantInvokeLatestVersionParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvokeLatestVersion",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVersion) GrantInvokeUrl(identity awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IVersion) GrantInvokeVersion(identity awsiam.IGrantable, version IVersion) awsiam.Grant {
	if err := i.validateGrantInvokeVersionParameters(identity, version); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvokeVersion",
		[]interface{}{identity, version},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVersion) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IVersion) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IVersion) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IVersion) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IVersion) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (j *jsiiProxy_IVersion) EdgeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) Lambda() IFunction {
	var returns IFunction
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) LatestVersion() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) ResourceArnsForGrantInvoke() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnsForGrantInvoke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

