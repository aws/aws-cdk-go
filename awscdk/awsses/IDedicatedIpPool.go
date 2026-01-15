package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses"
	"github.com/aws/constructs-go/constructs/v10"
)

// A dedicated IP pool.
type IDedicatedIpPool interface {
	interfacesawsses.IDedicatedIpPoolRef
	awscdk.IResource
	// The name of the dedicated IP pool.
	DedicatedIpPoolName() *string
}

// The jsii proxy for IDedicatedIpPool
type jsiiProxy_IDedicatedIpPool struct {
	internal.Type__interfacesawssesIDedicatedIpPoolRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IDedicatedIpPool) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IDedicatedIpPool) DedicatedIpPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedIpPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDedicatedIpPool) DedicatedIpPoolRef() *interfacesawsses.DedicatedIpPoolReference {
	var returns *interfacesawsses.DedicatedIpPoolReference
	_jsii_.Get(
		j,
		"dedicatedIpPoolRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDedicatedIpPool) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDedicatedIpPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDedicatedIpPool) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

