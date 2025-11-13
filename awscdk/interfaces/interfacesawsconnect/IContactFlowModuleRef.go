package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContactFlowModule.
// Experimental.
type IContactFlowModuleRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ContactFlowModule resource.
	// Experimental.
	ContactFlowModuleRef() *ContactFlowModuleReference
}

// The jsii proxy for IContactFlowModuleRef
type jsiiProxy_IContactFlowModuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IContactFlowModuleRef) ContactFlowModuleRef() *ContactFlowModuleReference {
	var returns *ContactFlowModuleReference
	_jsii_.Get(
		j,
		"contactFlowModuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContactFlowModuleRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContactFlowModuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

