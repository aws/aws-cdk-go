package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsssm"
	"github.com/aws/constructs-go/constructs/v10"
)

// An SSM Parameter reference.
type IParameter interface {
	interfacesawsssm.IParameterRef
	awscdk.IResource
	// Grants read (DescribeParameter, GetParameters, GetParameter, GetParameterHistory) permissions on the SSM Parameter.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grants write (PutParameter) permissions on the SSM Parameter.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the SSM Parameter resource.
	ParameterArn() *string
	// The name of the SSM Parameter resource.
	ParameterName() *string
	// The type of the SSM Parameter resource.
	ParameterType() *string
}

// The jsii proxy for IParameter
type jsiiProxy_IParameter struct {
	internal.Type__interfacesawsssmIParameterRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IParameter) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IParameter) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IParameter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IParameter) ParameterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameter) ParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameter) ParameterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameter) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameter) ParameterRef() *interfacesawsssm.ParameterReference {
	var returns *interfacesawsssm.ParameterReference
	_jsii_.Get(
		j,
		"parameterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

