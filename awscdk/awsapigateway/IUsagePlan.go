package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A UsagePlan, either managed by this CDK app, or imported.
type IUsagePlan interface {
	awscdk.IResource
	IUsagePlanRef
	// Adds an ApiKey.
	AddApiKey(apiKey IApiKey, options *AddApiKeyOptions)
	// Id of the usage plan.
	UsagePlanId() *string
}

// The jsii proxy for IUsagePlan
type jsiiProxy_IUsagePlan struct {
	internal.Type__awscdkIResource
	jsiiProxy_IUsagePlanRef
}

func (i *jsiiProxy_IUsagePlan) AddApiKey(apiKey IApiKey, options *AddApiKeyOptions) {
	if err := i.validateAddApiKeyParameters(apiKey, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addApiKey",
		[]interface{}{apiKey, options},
	)
}

func (i *jsiiProxy_IUsagePlan) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IUsagePlan) UsagePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUsagePlan) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUsagePlan) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUsagePlan) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

