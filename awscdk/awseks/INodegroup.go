package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawseks"
	"github.com/aws/constructs-go/constructs/v10"
)

// NodeGroup interface.
type INodegroup interface {
	interfacesawseks.INodegroupRef
	awscdk.IResource
	// Name of the nodegroup.
	NodegroupName() *string
}

// The jsii proxy for INodegroup
type jsiiProxy_INodegroup struct {
	internal.Type__interfacesawseksINodegroupRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_INodegroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_INodegroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_INodegroup) NodegroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodegroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INodegroup) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INodegroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INodegroup) NodegroupRef() *interfacesawseks.NodegroupReference {
	var returns *interfacesawseks.NodegroupReference
	_jsii_.Get(
		j,
		"nodegroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INodegroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

