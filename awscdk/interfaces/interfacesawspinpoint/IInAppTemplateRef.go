package interfacesawspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InAppTemplate.
// Experimental.
type IInAppTemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a InAppTemplate resource.
	// Experimental.
	InAppTemplateRef() *InAppTemplateReference
}

// The jsii proxy for IInAppTemplateRef
type jsiiProxy_IInAppTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IInAppTemplateRef) InAppTemplateRef() *InAppTemplateReference {
	var returns *InAppTemplateReference
	_jsii_.Get(
		j,
		"inAppTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInAppTemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInAppTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

