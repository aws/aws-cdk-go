package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappsync"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for AppSync Functions.
type IAppsyncFunction interface {
	interfacesawsappsync.IFunctionConfigurationRef
	awscdk.IResource
	// the ARN of the AppSync function.
	FunctionArn() *string
	// the name of this AppSync Function.
	FunctionId() *string
}

// The jsii proxy for IAppsyncFunction
type jsiiProxy_IAppsyncFunction struct {
	internal.Type__interfacesawsappsyncIFunctionConfigurationRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAppsyncFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IAppsyncFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppsyncFunction) FunctionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppsyncFunction) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppsyncFunction) FunctionConfigurationRef() *interfacesawsappsync.FunctionConfigurationReference {
	var returns *interfacesawsappsync.FunctionConfigurationReference
	_jsii_.Get(
		j,
		"functionConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppsyncFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppsyncFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

