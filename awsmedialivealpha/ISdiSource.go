package awsmedialivealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
	"github.com/aws/aws-cdk-go/awsmedialivealpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a MediaLive SDI Source.
// Experimental.
type ISdiSource interface {
	awscdk.IResource
	interfacesawsmedialive.ISdiSourceRef
	// The SDI Source ARN.
	// Experimental.
	SdiSourceArn() *string
	// The SDI Source ID.
	// Experimental.
	SdiSourceId() *string
	// The list of inputs currently using this SDI source.
	// Experimental.
	SdiSourceInputs() *[]*string
	// The current state of the SDI source.
	// Experimental.
	SdiSourceState() *string
}

// The jsii proxy for ISdiSource
type jsiiProxy_ISdiSource struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsmedialiveISdiSourceRef
}

func (i *jsiiProxy_ISdiSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ISdiSource) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISdiSource) SdiSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdiSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISdiSource) SdiSourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdiSourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISdiSource) SdiSourceInputs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sdiSourceInputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISdiSource) SdiSourceState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdiSourceState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISdiSource) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISdiSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISdiSource) SdiSourceRef() *interfacesawsmedialive.SdiSourceReference {
	var returns *interfacesawsmedialive.SdiSourceReference
	_jsii_.Get(
		j,
		"sdiSourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISdiSource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

