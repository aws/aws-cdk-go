package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodedeploy"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a reference to a CodeDeploy Application deploying to EC2/on-premise instances.
//
// If you're managing the Application alongside the rest of your CDK resources,
// use the `ServerApplication` class.
//
// If you want to reference an already existing Application,
// or one defined in a different CDK Stack,
// use the `#fromServerApplicationName` method.
type IServerApplication interface {
	interfacesawscodedeploy.IApplicationRef
	awscdk.IResource
	ApplicationArn() *string
	ApplicationName() *string
}

// The jsii proxy for IServerApplication
type jsiiProxy_IServerApplication struct {
	internal.Type__interfacesawscodedeployIApplicationRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IServerApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IServerApplication) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IServerApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerApplication) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerApplication) ApplicationRef() *interfacesawscodedeploy.ApplicationReference {
	var returns *interfacesawscodedeploy.ApplicationReference
	_jsii_.Get(
		j,
		"applicationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerApplication) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerApplication) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

