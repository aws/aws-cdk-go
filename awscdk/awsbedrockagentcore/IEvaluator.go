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

// Interface for Evaluator resources.
type IEvaluator interface {
	interfacesawsbedrockagentcore.IEvaluatorRef
	awscdk.IResource
	// Grant the given principal identity permissions to perform actions on this evaluator.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// The timestamp when the evaluator was created.
	CreatedAt() *string
	// The ARN of the evaluator.
	EvaluatorArn() *string
	// The unique identifier of the evaluator.
	EvaluatorId() *string
	// The name of the evaluator.
	EvaluatorName() *string
	// The lifecycle status of the evaluator (CREATING, ACTIVE, FAILED, DELETING).
	Status() *string
	// The timestamp when the evaluator was last updated.
	UpdatedAt() *string
}

// The jsii proxy for IEvaluator
type jsiiProxy_IEvaluator struct {
	internal.Type__interfacesawsbedrockagentcoreIEvaluatorRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IEvaluator) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IEvaluator) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IEvaluator) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEvaluator) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEvaluator) EvaluatorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluatorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEvaluator) EvaluatorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEvaluator) EvaluatorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluatorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEvaluator) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEvaluator) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEvaluator) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEvaluator) EvaluatorRef() *interfacesawsbedrockagentcore.EvaluatorReference {
	var returns *interfacesawsbedrockagentcore.EvaluatorReference
	_jsii_.Get(
		j,
		"evaluatorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEvaluator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEvaluator) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

