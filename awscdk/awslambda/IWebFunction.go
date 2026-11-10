package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a WebFunction resource, either created in this stack or imported.
type IWebFunction interface {
	awsiam.IGrantable
	awscdk.IResource
	awsiam.IResourceWithPolicyV2
	interfacesawslambda.IWebFunctionRef
	// Adds a statement to the WebFunction's execution role.
	//
	// For an imported WebFunction the role is unknown, so this adds a warning asking for the
	// permission to be granted manually instead.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// The execution role of the WebFunction.
	//
	// Undefined for an imported WebFunction.
	Role() interfacesawsiam.IRoleRef
	// The ARN of the WebFunction.
	WebFunctionArn() *string
	// The name of the WebFunction.
	WebFunctionName() *string
}

// The jsii proxy for IWebFunction
type jsiiProxy_IWebFunction struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
	internal.Type__awsiamIResourceWithPolicyV2
	internal.Type__interfacesawslambdaIWebFunctionRef
}

func (i *jsiiProxy_IWebFunction) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := i.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (i *jsiiProxy_IWebFunction) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := i.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		i,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IWebFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IWebFunction) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IWebFunction) Role() interfacesawsiam.IRoleRef {
	var returns interfacesawsiam.IRoleRef
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunction) WebFunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webFunctionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunction) WebFunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webFunctionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunction) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunction) WebFunctionRef() *interfacesawslambda.WebFunctionReference {
	var returns *interfacesawslambda.WebFunctionReference
	_jsii_.Get(
		j,
		"webFunctionRef",
		&returns,
	)
	return returns
}

