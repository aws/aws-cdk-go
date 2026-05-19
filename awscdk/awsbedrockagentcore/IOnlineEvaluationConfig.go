package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrockagentcore/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for OnlineEvaluationConfig resources.
type IOnlineEvaluationConfig interface {
	awsiam.IGrantable
	interfacesawsbedrockagentcore.IOnlineEvaluationConfigRef
	awscdk.IResource
	// Grant the given principal identity permissions to perform actions on this configuration.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// The timestamp when the configuration was created.
	CreatedAt() *string
	// The IAM execution role for the evaluation.
	ExecutionRole() awsiam.IRole
	// The execution status of the evaluation (ENABLED or DISABLED).
	ExecutionStatus() *string
	// The ARN of the online evaluation configuration.
	OnlineEvaluationConfigArn() *string
	// The unique identifier of the online evaluation configuration.
	OnlineEvaluationConfigId() *string
	// The name of the online evaluation configuration.
	OnlineEvaluationConfigName() *string
	// The lifecycle status of the configuration (CREATING, ACTIVE, FAILED, DELETING).
	Status() *string
	// The timestamp when the configuration was last updated.
	UpdatedAt() *string
}

// The jsii proxy for IOnlineEvaluationConfig
type jsiiProxy_IOnlineEvaluationConfig struct {
	internal.Type__awsiamIGrantable
	internal.Type__interfacesawsbedrockagentcoreIOnlineEvaluationConfigRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IOnlineEvaluationConfig) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IOnlineEvaluationConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IOnlineEvaluationConfig) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IOnlineEvaluationConfig) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOnlineEvaluationConfig) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOnlineEvaluationConfig) ExecutionStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOnlineEvaluationConfig) OnlineEvaluationConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onlineEvaluationConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOnlineEvaluationConfig) OnlineEvaluationConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onlineEvaluationConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOnlineEvaluationConfig) OnlineEvaluationConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onlineEvaluationConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOnlineEvaluationConfig) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOnlineEvaluationConfig) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOnlineEvaluationConfig) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOnlineEvaluationConfig) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOnlineEvaluationConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOnlineEvaluationConfig) OnlineEvaluationConfigRef() *interfacesawsbedrockagentcore.OnlineEvaluationConfigReference {
	var returns *interfacesawsbedrockagentcore.OnlineEvaluationConfigReference
	_jsii_.Get(
		j,
		"onlineEvaluationConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOnlineEvaluationConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

