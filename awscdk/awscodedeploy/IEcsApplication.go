package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodedeploy"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a reference to a CodeDeploy Application deploying to Amazon ECS.
//
// If you're managing the Application alongside the rest of your CDK resources,
// use the `EcsApplication` class.
//
// If you want to reference an already existing Application,
// or one defined in a different CDK Stack,
// use the `EcsApplication#fromEcsApplicationName` method.
type IEcsApplication interface {
	interfacesawscodedeploy.IApplicationRef
	awscdk.IResource
	ApplicationArn() *string
	ApplicationName() *string
}

// The jsii proxy for IEcsApplication
type jsiiProxy_IEcsApplication struct {
	internal.Type__interfacesawscodedeployIApplicationRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IEcsApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IEcsApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsApplication) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsApplication) ApplicationRef() *interfacesawscodedeploy.ApplicationReference {
	var returns *interfacesawscodedeploy.ApplicationReference
	_jsii_.Get(
		j,
		"applicationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsApplication) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsApplication) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

