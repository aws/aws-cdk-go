package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContactFlowModuleVersion.
// Experimental.
type IContactFlowModuleVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ContactFlowModuleVersion resource.
	// Experimental.
	ContactFlowModuleVersionRef() *ContactFlowModuleVersionReference
}

// The jsii proxy for IContactFlowModuleVersionRef
type jsiiProxy_IContactFlowModuleVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IContactFlowModuleVersionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IContactFlowModuleVersionRef) ContactFlowModuleVersionRef() *ContactFlowModuleVersionReference {
	var returns *ContactFlowModuleVersionReference
	_jsii_.Get(
		j,
		"contactFlowModuleVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContactFlowModuleVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContactFlowModuleVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

